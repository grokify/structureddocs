// Package schema provides utilities for JSON Schema embedding and validation
// in structured document projects.
package schema

import (
	"encoding/json"
	"fmt"
)

// Registry holds embedded JSON schemas for a document type.
type Registry struct {
	// ID is the schema identifier URL (e.g., "https://github.com/grokify/structured-goals/schema/v2mom.schema.json").
	ID string

	// Version is the schema version (e.g., "1.0.0").
	Version string

	// Data is the raw JSON schema bytes.
	Data []byte
}

// NewRegistry creates a new schema registry.
func NewRegistry(id, version string, data []byte) *Registry {
	return &Registry{
		ID:      id,
		Version: version,
		Data:    data,
	}
}

// JSON returns the schema as bytes.
func (r *Registry) JSON() []byte {
	return r.Data
}

// JSONString returns the schema as a string.
func (r *Registry) JSONString() string {
	return string(r.Data)
}

// Unmarshal parses the schema into the provided struct.
func (r *Registry) Unmarshal(v any) error {
	return json.Unmarshal(r.Data, v)
}

// SchemaRef represents a $schema reference in a document.
type SchemaRef struct {
	Schema string `json:"$schema,omitempty"`
}

// ExtractSchemaRef extracts the $schema field from JSON data.
func ExtractSchemaRef(data []byte) (string, error) {
	var ref SchemaRef
	if err := json.Unmarshal(data, &ref); err != nil {
		return "", fmt.Errorf("extracting $schema: %w", err)
	}
	return ref.Schema, nil
}

// ValidateSchemaRef checks if the document's $schema matches the expected schema ID.
func ValidateSchemaRef(data []byte, expectedID string) error {
	ref, err := ExtractSchemaRef(data)
	if err != nil {
		return err
	}
	if ref == "" {
		return nil // No schema reference, skip validation
	}
	if ref != expectedID {
		return fmt.Errorf("schema mismatch: expected %s, got %s", expectedID, ref)
	}
	return nil
}
