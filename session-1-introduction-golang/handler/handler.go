package handler

import "github.com/gin-gonic/gin"

//RootHandler untuk get request
func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

//RootHandler untuk post request
func PostHandler(c *gin.Context) {
	var json struct {
		Message  string `json:"message"`
		Location string `json:"location"`
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(200, gin.H{
			"message":  json.Message,
			"location": json.Location,
		})
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}
