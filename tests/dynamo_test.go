package main

import (
	"github.com/albibenni/dynamo-go-test/database"
	"testing"
)

func Test(t *testing.T) {
	client, err := database.NewDynamoDBClient()
	if err != nil {
		t.Error("Error creating database")
	}
	_, err2 := client.GetUser("Nope")
	if err2 == nil {
		t.Error("Error, no user yet imported")
	}
}
