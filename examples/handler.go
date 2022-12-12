package examples

import (
	router_actions "gin-utils-router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerWithoutId(c *gin.Context) {
	if err := router_actions.Actions(c, router_actions.ActionMap{
		":ping": {"POST": func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"action":  "ping",
				"message": "pong",
			})
		}},
	}); err != nil {
		c.JSON(http.StatusNotFound, err)
	}
}

func HandlerWithId(c *gin.Context) {
	if err := router_actions.ActionsWithObjId(c, router_actions.ActionMap{
		":ping": {"POST": func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"action":  "ping",
				"message": "pong",
				"_id":     c.GetString("_id"),
			})
		}},
	}); err != nil {
		c.JSON(http.StatusNotFound, err)
	}
}

func HandlerWithString(c *gin.Context) {
	key := "name"
	if err := router_actions.ActionsWithString(c, router_actions.ActionMap{
		":ping": {"POST": func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"action":  "ping",
				"message": "pong",
				key:       c.GetString(key),
			})
		}},
	}, key); err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}
}
