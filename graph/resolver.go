package graph

import "github.com/ChihHaoChen/graphql-booking-server/graph/model"

//go:generate go get github.com/99designs/gqlgen/cmd@v0.14.0
//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	travels []*model.Travel
	users []*model.User
}
