package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", HelloWorld)
    r.Run()
}

func HelloWorld(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Hello World"})
    return
}
