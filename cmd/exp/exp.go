// experimenting with context

package main

import (
	"context"
	"fmt"
	"strings"
)

type key string

const (
	favoriteColorKey key = "favorite-color"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "dark-red")

	value := ctx.Value(favoriteColorKey)
	strValue := value.(string)
	fmt.Println(strValue)
	fmt.Println(strings.HasPrefix(strValue, "dark"))
}
