package handlers

import (
	"mgmt/database"
	"mgmt/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Testfunc2(c *gin.Context) {
	// var structData FormData
	name := c.PostForm("fname")
	last := c.PostForm("lname")

	formOutput := new(views.FormOutput)
	formOutput.FirstName = name
	formOutput.LastName = last

	// json, _ := json.Marshal(formOutput)

	// jsonForm := string(json)

	cfg, db := database.AWSConnection()

	database.ListTables(cfg, db)
	database.PutItems(cfg, db, formOutput)

	// Response to client
	c.JSON(http.StatusOK, gin.H{
		"body": "testing dynamo",
	})
}