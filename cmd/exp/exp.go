// experimenting with context

package main

import (
	standardCtx "context"
	"fmt"

	"github.com/AlexTLDR/WebDev/context"
	"github.com/AlexTLDR/WebDev/models"
)

type key string

const (
	favoriteColorKey key = "favorite-color"
)

func main() {
	ctx := standardCtx.Background()

	user := models.User{
		Email: "alex@alex.com",
	}
	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
