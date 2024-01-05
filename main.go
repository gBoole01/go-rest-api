package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "The Slim Shady LP", Artist: "Eminem", Price: 12.99},
	{ID: "2", Title: "The Marshall Mathers LP", Artist: "Eminem", Price: 14.99},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Price: 16.99},
	{ID: "4", Title: "Illmatic", Artist: "Nas", Price: 11.99},
	{ID: "5", Title: "To Pimp a Butterfly", Artist: "Kendrick Lamar", Price: 14.99},
	{ID: "6", Title: "My Beautiful Dark Twisted Fantasy", Artist: "Kanye West", Price: 15.99},
	{ID: "7", Title: "The Blueprint", Artist: "Jay Z", Price: 13.99},
	{ID: "8", Title: "It Takes a Nation of Millions to Hold Us Back", Artist: "Public Enemy", Price: 12.99},
	{ID: "9", Title: "Straight Outta Compton", Artist: "N.W.A", Price: 11.99},
	{ID: "10", Title: "Ready to Die", Artist: "The Notorious B.I.G.", Price: 16.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}