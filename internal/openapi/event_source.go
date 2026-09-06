package openapi

import (
	"encoding/json"
	"fmt"
)

// maybeHydrateEmptyEventSource inserts `"source":"device"` when Event JSON
// omits source, or source is null/empty. Unknown non-empty values fail.
// No-op for other oneOf models (EventRuleAction).
func maybeHydrateEmptyEventSource(className string, data []byte) ([]byte, error) {
	if className != "Event" {
		return data, nil
	}
	return hydrateEmptyEventSource(data)
}

func hydrateEmptyEventSource(data []byte) ([]byte, error) {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, fmt.Errorf("event JSON must be an object: %w", err)
	}
	if obj == nil {
		return nil, fmt.Errorf("event JSON must be an object")
	}

	raw, ok := obj["source"]
	if !ok {
		obj["source"] = json.RawMessage(`"device"`)
		return json.Marshal(obj)
	}

	var source *string
	if err := json.Unmarshal(raw, &source); err != nil {
		return nil, fmt.Errorf("invalid Event source: %w", err)
	}
	if source == nil || *source == "" {
		obj["source"] = json.RawMessage(`"device"`)
		return json.Marshal(obj)
	}
	if *source != "device" && *source != "edge" {
		return nil, fmt.Errorf("unknown Event source %q", *source)
	}
	return data, nil
}
