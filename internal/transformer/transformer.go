package transformer

import (
	"github.com/ibllex/go-fractal"
)

func defaultManager() *fractal.Manager {

	m := fractal.NewManager(nil)
	m.SetSerializer(&fractal.ArraySerializer{})

	return m
}

func Item(item fractal.Any, transformer fractal.Transformer) fractal.M {

	m := defaultManager()

	resource := fractal.NewItem(
		fractal.WithData(item),
		fractal.WithTransformer(transformer),
	)

	d,err := m.CreateData(resource, nil).ToMap()

	if err != nil {
		return fractal.M{}
	}

	return d
}



func Collection(items []fractal.Any, transformer fractal.Transformer) fractal.M {

	m := defaultManager()

	resource := fractal.NewCollection(
		fractal.WithData(items),
		fractal.WithTransformer(transformer),
	)


	d,err := m.CreateData(resource, nil).ToMap()

	if err != nil {
		return fractal.M{}
	}

	return d
}
