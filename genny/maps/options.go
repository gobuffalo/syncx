package maps

import "html/template"

type Options struct {
	Maps []Map
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}

type Map struct {
	Name   template.HTML
	GoType template.HTML
	A      template.HTML
	B      template.HTML
	BB     template.HTML
	C      template.HTML
}
