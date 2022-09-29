package v1

import "encoding/json"

// Collection provides a schema for a manifest that extend the OCI manifest when marshaled to JSON.
type Collection struct {
	// MediaType is the media type of the object this schema refers to.
	MediaType string `json:"mediaType"`

	// ArtifactType is the IANA media type of the artifact this schema refers to.
	ArtifactType string `json:"artifactType"`

	// Blobs is a collection of blobs referenced by this manifest.
	Blobs []Descriptor `json:"blobs,omitempty"`

	// Refers is an optional link to any existing manifest within the repository.
	Refers *Descriptor `json:"refers,omitempty"`

	// Attributes contains typed attributes for this artifact.
	Attributes map[string]json.RawMessage `json:"attributes,omitempty"`

	// Schema is an optional link to an existing manifest that contains a schema for
	// attributes validation.
	Schema *Descriptor `json:"schema,omitempty"`

	// Links defines are reference to any manifest(s) for Collection that this Collection
	// is dependent on or relates to.
	Links []*Collection `json:"links,omitempty"`

	// Annotations contains arbitrary metadata for the artifact manifest.
	Annotations map[string]string `json:"annotations,omitempty"`
}
