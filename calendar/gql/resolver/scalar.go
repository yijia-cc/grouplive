package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go/decode"
)

var _ decode.Unmarshaler = (*Duration)(nil)

type Duration struct {
	time.Duration
}

func (d Duration) ImplementsGraphQLType(name string) bool {
	return name == "Duration"
}

func (d *Duration) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Duration:
		d.Duration = input
	case int:
		d.Duration = time.Duration(input)
	}
	return nil
}

var _ decode.Unmarshaler = (*Void)(nil)

type Void struct{}

func (v Void) ImplementsGraphQLType(name string) bool {
	return name == "Void"
}

func (v Void) UnmarshalGraphQL(input interface{}) error {
	return nil
}
