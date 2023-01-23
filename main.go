package main

import (
	"fmt"
	confluence "github.com/virtomize/confluence-go-api"
	"log"
)

func main() {
	// Initialize a new Confluence client with your Confluence URL and API token
	client, err := confluence.NewAPI("https://jimmyktest.atlassian.net/wiki/rest/api", "castlevanity@gmail.com", "370Cda45fgt")
	if err != nil {
		log.Fatal("newApi: ", err)
	}

	// get content by content id
	//c, err := client.GetContentByID("425985", confluence.ContentQuery{
	//	SpaceKey: "TEST",
	//	Expand:   []string{"body.storage", "version"},
	//})
	//if err != nil {
	//	log.Fatal("Get: ", err)
	//}
	//fmt.Printf("%+v\n", c)

	/////////////////

	// Define the content of the new page
	content := &confluence.Content{
		Type:  "page",         // can also be blogpost
		Title: "SomeTestPage", // page title
		Ancestors: []confluence.Ancestor{
			{
				ID: "425985", // ancestor-id optional if you want to create sub-pages
			},
		},
		Body: confluence.Body{
			Storage: confluence.Storage{
				Value:          "apiTest!!!", // your page content here
				Representation: "storage",
			},
		},
		Version: &confluence.Version{
			Number: 1,
		},
		Space: &confluence.Space{
			Key: "TEST", // Space
		},
	}

	c, err := client.CreateContent(content)
	if err != nil {
		log.Fatal("CreateContent: ", err)
	}

	fmt.Printf("%+v\n", c)
}
