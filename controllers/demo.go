package controllers

import (
	"github.com/gin-gonic/gin"
	log "go-session-demo/helpers"
	"net/http"
)

type Demo struct {
}

func (d *Demo) DemoA(ctx *gin.Context) {
	a := &log.PrintLog{}
	a.Print("这是 demo-main")

	res := map[string]string{
		"A": "1",
	}

	ctx.JSON(http.StatusOK, res)
}
