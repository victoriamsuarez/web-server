package main

import "github.com/gin-gonic/gin"

func main() {

	server := gin.Default()

	server.GET("/ping", GetPong)
	server.POST("/saludo", PostGreeting)

	// Correr servidor
	server.Run(":8081")

}

func GetPong(c *gin.Context) {
	c.JSON(200, "pong")
}

type NameCompleted struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func PostGreeting(c *gin.Context) {
	var nameComp NameCompleted
	c.ShouldBindJSON(&nameComp)
	c.JSON(201, "Hola "+nameComp.Name+" "+nameComp.LastName)

}
