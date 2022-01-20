package putio_test

import (
	"context"
	"fmt"
	"log"

	"github.com/putdotio/go-putio/putio"
	"golang.org/x/oauth2"
)

const token = "<YOUR-TOKEN-HERE>"

func Example() {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	oauthClient := oauth2.NewClient(context.TODO(), tokenSource)
	client := putio.NewClient(oauthClient)

	// get root directory
	root, err := client.Files.Get(context.TODO(), 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name of root folder is: %s\n", root.Name)
}
