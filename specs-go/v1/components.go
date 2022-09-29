package v1

import "encoding/json"

// Component schema defines information to create a component list.
type Component struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	Version            string          `json:"version"`
	Type               string          `json:"type"`
	FoundBy            string          `json:"foundBy"`
	Locations          []string        `json:"locations"`
	Licenses           []string        `json:"licenses"`
	Language           string          `json:"language"`
	CPEs               []string        `json:"cpes"`
	PURL               string          `json:"purl"`
	AdditionalMetadata json.RawMessage `json:",inline"`
}
