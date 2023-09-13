// experimenting with context

package main

import (
	"fmt"

	"github.com/AlexTLDR/WebDev/models"
)

func main() {
	// testing that globbing of the image file works - replace 2 with any gallery ID
	gs := models.GalleryService{}
	fmt.Println(gs.Images(2))
}
