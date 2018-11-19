package maps

import (
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}
	box := packr.New("../maps/templates", "../maps/templates")

	if err := opts.Validate(); err != nil {
		return gg, errors.WithStack(err)
	}

	for _, m := range opts.Maps {
		g := genny.New()

		for _, n := range []string{"map", "map_test"} {
			n = n + ".go.plush"
			s, err := box.FindString(n)
			if err != nil {
				return gg, errors.WithStack(err)
			}
			g.File(genny.NewFileS(strings.ToLower(string(m.Name))+"_"+n, s))
		}

		ctx := plush.NewContext()
		ctx.Set("opts", opts)
		ctx.Set("m", m)
		g.Transformer(plushgen.Transformer(ctx))
		gg.Add(g)
	}
	return gg, nil
}
