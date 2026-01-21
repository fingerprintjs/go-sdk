package sdk

import "context"

type Region string

const (
	RegionUS   Region = "us"
	RegionEU   Region = "eu"
	RegionAsia Region = "asia"
)

var regionBaseURLs = map[Region]string{
	RegionUS:   "https://api.fpjs.io/v4",
	RegionEU:   "https://eu.api.fpjs.io/v4",
	RegionAsia: "https://ap.api.fpjs.io/v4",
}

func WithRegionContext(ctx context.Context, cfg *Configuration, region Region) context.Context {
	baseURL, ok := regionBaseURLs[region]
	if !ok {
		// invalid region is passed
		return ctx
	}

	for i, server := range cfg.Servers {
		if server.URL == baseURL {
			return context.WithValue(ctx, ContextServerIndex, i)
		}
	}

	// region is not found in the OpenAPI schema
	return ctx
}
