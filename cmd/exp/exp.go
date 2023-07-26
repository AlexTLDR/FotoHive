package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "favorite-color", "dark-red")
	value := ctx.Value("favorite-color")
	fmt.Println(value)
}
