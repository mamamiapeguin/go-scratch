package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scratch/controllers" // relative path to $HOME/go/src
)

func main() {
	r := gin.Default()
	r.GET("/user/index", controllers.Index)

	r.Run(":8080")
}
