package maps

import "html/template"

// Options for a sync.Map implementation
type Options struct {
	Maps []Map
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}

// Map represents an implementation of the map
type Map struct {
	Name   template.HTML
	GoType template.HTML
	Zero   template.HTML
	A      template.HTML
	B      template.HTML
	BB     template.HTML
	C      template.HTML
}
