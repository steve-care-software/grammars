package everythings

import grammars "github.com/steve-care-software/grammars/domain"

type everything struct {
	everythingBuilder grammars.EverythingBuilder
}

func createEverything(
	everythingBuilder grammars.EverythingBuilder,
) Everything {
	out := everything{
		everythingBuilder: everythingBuilder,
	}

	return &out
}

// WithoutEscape returns an everything without escape
func (app *everything) WithoutEscape(exception grammars.Token) grammars.Everything {
	return app.Everything(exception, nil)
}

// Everything returns an everything instance
func (app *everything) Everything(exception grammars.Token, escape grammars.Token) grammars.Everything {
	builder := app.everythingBuilder.Create().
		WithException(exception)

	if escape != nil {
		builder.WithEscape(escape)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
