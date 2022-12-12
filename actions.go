package router_actions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type ActionMap map[string]map[string]func(c *gin.Context)

var ActionNotfound = errors.New("action not found")

func Actions(c *gin.Context, actionMap ActionMap) error {
	action, _ := c.Params.Get("action")
	mapping, ok := actionMap[action]
	if !ok {
		return ActionNotfound
	}
	f, ok := mapping[c.Request.Method]
	if !ok {
		return ActionNotfound
	}
	f(c)
	return nil
}

func ActionsInList(c *gin.Context, actionMap ActionMap) error {
	var action string
	separate := strings.Split(c.Params[0].Value, ":")
	switch len(separate) {
	case 1:
		separate = strings.Split(c.Params[0].Value, "/")
		if len(separate) == 3 {
			action = separate[2]
		}
	case 2:
		action = ":" + separate[1]
	default:
		action = ""
	}
	mapping, ok := actionMap[action]
	if !ok {
		return ActionNotfound
	}
	f, ok := mapping[c.Request.Method]
	if !ok {
		return ActionNotfound
	}
	f(c)
	return nil
}

func ActionsWithObjId(c *gin.Context, actionMap ActionMap) error {
	action := c.Params[len(c.Params)-1].Value
	uris := strings.Split(action, ":")
	if uris[0] == "batch" {
		return Actions(c, actionMap)
	}
	if len(uris) != 2 || !primitive.IsValidObjectID(uris[0]) {
		return ActionNotfound
	}
	c.Params[len(c.Params)-1].Value = uris[0]
	action = ":" + uris[1]
	mapping, ok := actionMap[action]
	if !ok {
		return ActionNotfound
	}
	f, ok := mapping[c.Request.Method]
	if !ok {
		return ActionNotfound
	}
	c.Set("_id", uris[0])
	c.Next()
	f(c)
	return nil
}

func ActionsWithString(c *gin.Context, actionMap ActionMap, key string) error {
	action := c.Params[len(c.Params)-1].Value
	uris := strings.Split(action, ":")
	if uris[0] == "batch" {
		return Actions(c, actionMap)
	}
	if len(uris) != 2 {
		return ActionNotfound
	}
	c.Params[len(c.Params)-1].Value = uris[0]
	action = ":" + uris[1]
	mapping, ok := actionMap[action]
	if !ok {
		return ActionNotfound
	}
	f, ok := mapping[c.Request.Method]
	if !ok {
		return ActionNotfound
	}
	c.Set(key, uris[0])
	c.Next()
	f(c)
	return nil
}
