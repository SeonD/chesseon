package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/SeonD/chesseon/db"
	"github.com/gin-gonic/gin"
)

// Connect to database
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// ErrorHandler function
func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}

// SetConstants sets constant values
func SetConstants(c *gin.Context) {
	jsonFile, err := os.Open("board_setup.json")

	if err != nil {
		c.HTML(http.StatusBadRequest, "500", gin.H{})
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var Board map[string]string
	json.Unmarshal([]byte(byteValue), &Board)

	c.Set("boardSetup", Board)
	c.Next()
}
