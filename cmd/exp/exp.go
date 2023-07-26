// experimenting with context

package main

import (
	"context"
	"fmt"
)

type key string

const (
	favoriteColorKey key = "favorite-color"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "dark-red")

	ctx = context.WithValue(ctx, "favorite-color", 123)

	value := ctx.Value(favoriteColorKey)
	value2 := ctx.Value("favorite-color")
	fmt.Println(value, value2)
}
