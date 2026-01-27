//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageJSON struct {
	Version string `json:"version"`
}

func main() {
	data, err := os.ReadFile("package.json")
	if err != nil {
		panic(err)
	}

	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		panic(err)
	}

	if pkg.Version == "" {
		panic("package.json version is empty")
	}

	out := fmt.Sprintf(`// Code auto-generated - DO NOT EDIT!

package fingerprint

const Version = %q
`, pkg.Version)
	if err := os.WriteFile("version.go", []byte(out), 0644); err != nil {
		panic(err)
	}
}
