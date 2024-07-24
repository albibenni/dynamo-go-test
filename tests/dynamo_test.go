package main

import (
	"testing"

	"github.com/albibenni/dynamo-go-test/database"
)

const (
	TABLE_NAME = "test-db"
)

func TestPort(t *testing.T) {
	_, port := database.CreateLocalClient(8000)
	if port != 8000 {
		t.Error("Error, wrong port")
	}
}

func TestExists(t *testing.T) {
	client, _ := database.CreateLocalClient(8000)
	exist := database.TableExists(client, "test")
	if exist {
		t.Error("Error, it doesn't")
	}
}

func TestCreateTable(t *testing.T) {
	client, _ := database.CreateLocalClient(8000)
	created := database.CreateTableIfNotExists(client, TABLE_NAME)
	if !created {
		t.Error("Error, not created")
	}

}
