// Package metadata provides common metadata types used across structured documents.
// These types ensure consistent representation of authors, versions, and status
// across PRD, MRD, TRD, V2MOM, OKR, Changelog, and Roadmap documents.
package metadata

import (
	"time"
)

// Author represents a document author or contributor.
type Author struct {
	// Name is the author's full name (required).
	Name string `json:"name"`

	// Email is the author's email address (optional).
	Email string `json:"email,omitempty"`

	// Role is the author's role (e.g., "Product Manager", "Tech Lead").
	Role string `json:"role,omitempty"`

	// URL is a link to the author's profile (optional).
	URL string `json:"url,omitempty"`
}

// Metadata contains common metadata fields for structured documents.
type Metadata struct {
	// ID is the unique document identifier.
	ID string `json:"id,omitempty"`

	// Title is the document title.
	Title string `json:"title,omitempty"`

	// Name is an alternative to Title (used by some document types).
	Name string `json:"name,omitempty"`

	// Version is the document version (e.g., "1.0.0", "2024.Q1").
	Version string `json:"version,omitempty"`

	// Status is the document status.
	Status Status `json:"status,omitempty"`

	// Authors is the list of document authors.
	Authors []Author `json:"authors,omitempty"`

	// Owner is the document owner (alternative to Authors for single-owner docs).
	Owner string `json:"owner,omitempty"`

	// CreatedAt is when the document was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// UpdatedAt is when the document was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Description is a brief description of the document.
	Description string `json:"description,omitempty"`

	// Tags are labels for categorization.
	Tags []string `json:"tags,omitempty"`
}

// GetTitle returns the title, falling back to name if title is empty.
func (m Metadata) GetTitle() string {
	if m.Title != "" {
		return m.Title
	}
	return m.Name
}

// GetOwner returns the owner, or the first author's name if owner is empty.
func (m Metadata) GetOwner() string {
	if m.Owner != "" {
		return m.Owner
	}
	if len(m.Authors) > 0 {
		return m.Authors[0].Name
	}
	return ""
}

// Touch updates the UpdatedAt timestamp to the current time.
func (m *Metadata) Touch() {
	m.UpdatedAt = time.Now()
}

// SetCreated sets the CreatedAt timestamp if not already set.
func (m *Metadata) SetCreated() {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
}
