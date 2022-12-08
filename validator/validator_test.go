package validator_test

import (
	"testing"

	"github.com/beyondan/gqlparser/v2"
	"github.com/beyondan/gqlparser/v2/ast"
	"github.com/beyondan/gqlparser/v2/parser"
	"github.com/beyondan/gqlparser/v2/validator"
	"github.com/stretchr/testify/require"
)

func TestExtendingNonExistantTypes(t *testing.T) {
	s := gqlparser.MustLoadSchema(
		&ast.Source{Name: "graph/schema.graphqls", Input: `
extend type User {
    id: ID!
}

extend type Product {
    upc: String!
}

union _Entity = Product | User

extend type Query {
	entity: _Entity
}
`, BuiltIn: false},
	)

	q, err := parser.ParseQuery(&ast.Source{Name: "ff", Input: `{
		entity {
		  ... on User {
			id
		  }
		}
	}`})
	require.Nil(t, err)
	require.Nil(t, validator.Validate(s, q))
}
