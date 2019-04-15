package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/moreal/boj-vs-code-api-server/problem"
	"google.golang.org/api/option"
	"log"
)

var connection Connection

type Connection struct {
	client *firestore.Client
	ctx context.Context
}

func (c *Connection) Initialize() {
	c.ctx = context.Background()

	option := option.WithCredentialsFile("./boj-vs-code-2b2e56865522.json")
	client, err := firestore.NewClient(c.ctx, "boj-vs-code", option)
	if err != nil {
		log.Panic("Couldn't make firestore client")
	}
	c.client = client
}

func (c *Connection) Add(collection string, id int, data interface{}) {
	_, err := c.client.Collection(collection).Doc(fmt.Sprintf("%d", id)).Set(c.ctx, data)
	if err != nil {
		log.Panic("Connection#Add Panic")
	}
}

func (c *Connection) Fetch(collection string, id int) *problem.ProblemModel {
	doc := c.client.Doc(fmt.Sprintf("%s/%d", collection, id))
	log.Print(doc)
	if doc == nil {
		return nil
	}

	docsnap, err := doc.Get(c.ctx)
	if err != nil {
		return nil
	}

	var problem problem.ProblemModel
	err = docsnap.DataTo(&problem)
	if err != nil {
		log.Panic("Failed to convert document ref to ProblemModel")
	}

	return &problem
}