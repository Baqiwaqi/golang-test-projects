package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create struct for a user
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// create a slice of users
var users = []User{
	{ID: "1", Name: "John"},
	{ID: "2", Name: "Pete"},
	{ID: "3", Name: "Mary"},
}

// create function to get users
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// get a user by id
func getUser(c *gin.Context) {
	id := c.Param("id")
	b, err := getUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, b)

}

func getUserById(id string) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func main() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUser)
	r.Run()
}
