package main

import (
	"encoding/json"
	"log"
	"mgmt/internal/data"
	"mgmt/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (app *application) createUsers(c *gin.Context) {
	// var struct Users
	jsonPost := new(data.Users)

	// var unknownJson map[string]any

	// Create UUID for each requests
	uuidV4, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	// Add the UUID to the Users ID struct field
	jsonPost.ID = uuidV4.String()

	// Bind the JSON request into the Users Struct
	err = c.BindJSON(&jsonPost)
	if err != nil {
		log.Fatal("Function Forms has failed with error:", err)
	}

	// Nice output like jq
	//nolint:errchkjson
	niceOutput, _ := json.MarshalIndent(jsonPost, "", "\t")
	log.Printf("JSON provided\n %+v", string(niceOutput))

	database.ListTables(app.db)
	database.PutItems(app.db, jsonPost)

	// Response to client
	c.IndentedJSON(http.StatusOK, jsonPost)
}
