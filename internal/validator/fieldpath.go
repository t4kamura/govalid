// Package validator implements rules for validating fields.
package validator

import "strings"

// FieldPath represents a field path with dot-separated components.
type FieldPath string

// NewFieldPath creates a new FieldPath from the given components.
func NewFieldPath(components ...string) FieldPath {
	return FieldPath(strings.Join(components, "."))
}

// WithoutDots returns the field path with all dots removed for use in variable names and keys.
func (fp FieldPath) WithoutDots() string {
	return strings.ReplaceAll(string(fp), ".", "")
}

// String returns the string representation of the field path.
func (fp FieldPath) String() string {
	return string(fp)
}
