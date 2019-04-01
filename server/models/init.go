package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"log"
)

var connection Connection

type Connection struct {
	client *firestore.Client
	ctx context.Context
}

func (c *Connection) Initialize() {
	c.ctx = context.Background()
	client, err := firestore.NewClient(c.ctx, "boj-vs-code")
	if err != nil {
		log.Panic("Couldn't make firestore client")
	}
	c.client = client
}

func (c *Connection) Add(collection string, data interface{}) {
	_, _, err := c.client.Collection(collection).Add(c.ctx, data)
	if err != nil {
		log.Panic("Connection#Add Panic")
	}
}

func (c *Connection) Fetch(collection string, id int) *ProblemModel {
	doc := c.client.Doc(fmt.Sprintf("%s/%d", collection, id))
	if doc == nil {
		return nil
	} else {
		docsnap, err := doc.Get(c.ctx)
		if err != nil {
			log.Panic("Failed to get document ref")
		}

		var problem ProblemModel
		err = docsnap.DataTo(&problem)
		if err != nil {
			log.Panic("Failed to convert document ref to ProblemModel")
		}

		return &problem
	}
}