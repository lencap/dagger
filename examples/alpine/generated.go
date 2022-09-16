// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"

	"go.dagger.io/dagger/sdk/go/dagger"
)

func (r *query) alpine(ctx context.Context) (*Alpine, error) {

	return new(Alpine), nil

}

type alpine struct{}
type query struct{}

func main() {
	dagger.Serve(context.Background(), map[string]func(context.Context, dagger.ArgsInput) (interface{}, error){
		"Alpine.build": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			var pkgs []string

			bytes, err = json.Marshal(fc.Args["pkgs"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &pkgs); err != nil {
				return nil, err
			}

			return (&alpine{}).build(ctx,

				pkgs,
			)
		},
		"Query.alpine": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			return (&query{}).alpine(ctx)
		},
	})
}
