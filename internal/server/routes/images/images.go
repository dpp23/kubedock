package images

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET "/images/json"
func (im *Images) ImageList(c *gin.Context) {
	c.JSON(http.StatusOK, []string{})
}

// GET "/images/:image/json"
func (im *Images) ImageJson(c *gin.Context) {
	id := c.Param("image")
	log.Printf("image: %s", id)
	c.JSON(http.StatusOK, gin.H{
		"Id":      id,
		"Created": "2018-12-18T01:20:53.669016181Z",
		"Size":    0,
		"ContainerConfig": gin.H{
			"Image": id,
		},
	})
}

// POST "/images/create"
func (im *Images) ImageCreate(c *gin.Context) {
	// from := c.Query("fromImage")
	c.JSON(http.StatusOK, gin.H{
		"status": "Download complete",
		// TODO: add progressdetail...
	})
}