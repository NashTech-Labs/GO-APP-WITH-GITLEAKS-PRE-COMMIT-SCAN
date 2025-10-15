package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var app_secret_key string = os.Getenv("app_secret_key")

type user struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   int    `json:"phone"`
}

type erruser struct {
	Msg string `json:"msg"`
}

var users = []user{
	{ID: "1", Name: "Shahenvaz", Address: "Kanpur", Phone: 9823493332},
	{ID: "2", Name: "Abhishek", Address: "Prayagraj", Phone: 9823493333},
	{ID: "3", Name: "TestUser3", Address: "Kanpur", Phone: 9823493331},
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user
	var payload map[string]interface{}

	if payloaderr := c.BindJSON(&payload); payloaderr != nil {
		fmt.Print(payloaderr)
		return
	}

	newUser.ID = payload["ID"].(string)
	newUser.Name = payload["Name"].(string)
	newUser.Address = payload["Address"].(string)
	if phoneFloat, ok := payload["Phone"].(float64); ok {
		newUser.Phone = int(phoneFloat)
	}

	if payload["app_secret_key"] != app_secret_key {
		fmt.Print(payload["app_secret_key"])
		var erra = erruser{Msg: "Secret key mismatched unable to create user!!!"}
		c.IndentedJSON(http.StatusUnauthorized, erra)
		return
	}
	if len(fmt.Sprintf("%d", newUser.Phone)) < 10 {
		var erra = erruser{Msg: "phone number must be exactly 10 digits"}
		c.IndentedJSON(http.StatusBadRequest, erra)
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	fmt.Print("heloo")
	fmt.Print(app_secret_key)
	router := gin.Default()
	router.GET("/get_users", getUsers)
	router.POST("/add_user", createUser)

	router.Run("localhost:8000")
}
