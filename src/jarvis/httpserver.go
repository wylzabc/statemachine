package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func convertTo(c *gin.Context){
	name := c.Param("name")
	res := convertAndLog(sm, name)
	c.String(http.StatusOK, res)
}

func getCurrentState(c *gin.Context) {
	s := sm.GetCurrentState()

	var req string
	if s == nil {
		req = "none"
	}else {
		req = s.Name()
	}
	c.String(http.StatusOK, req)
}
