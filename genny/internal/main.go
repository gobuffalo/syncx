package main

import (
	"context"
	"log"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/syncx/genny/maps"
)

func main() {
	opts := &maps.Options{
		Maps: []maps.Map{
			{
				Name:   "Int",
				GoType: "int",
				Zero:   "0",
				A:      "0",
				B:      "1",
				BB:     "-1",
				C:      "2",
			},
			{
				Name:   "String",
				GoType: "string",
				Zero:   `""`,
				A:      `"A"`,
				B:      `"B"`,
				BB:     `"BB"`,
				C:      `"C"`,
			},
		},
	}

	run := genny.WetRunner(context.Background())

	gg, err := maps.New(opts)
	if err != nil {
		log.Fatal(err)
	}
	run.WithGroup(gg)

	err = run.WithNew(gotools.GoFmt(run.Root))
	if err != nil {
		log.Fatal(err)
	}

	err = run.Run()
	if err != nil {
		log.Fatal(err)
	}
}
