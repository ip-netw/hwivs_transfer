// @Author : DAIPENGYUAN
// @File : main
// @Time : 2020/11/3 10:02 
// @Description : 

package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"hwivs_transfer/utils/httputil"
	"hwivs_transfer/utils/log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	flagSet  = flag.NewFlagSet("hwivs", flag.ExitOnError)
	port     = flagSet.Int("port", 8080, "service port")
	logLevel = flagSet.String("level", "info", "log level")
	logPath  = flagSet.String("log", "", "log file path")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery(), log.GinMiddleware())
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	log.ResetLogger(*logLevel, *logPath)
	engine.Handle("POST", "/hwivs", ReqIvs)
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: engine,
	}
	log.Logger.Info("Server Started at Port:" + strconv.Itoa(*port))
	if err := server.ListenAndServe(); err != nil {
		log.Logger.Info(err.Error())
	}
}

func ReqIvs(c *gin.Context) {
	var (
		httpClient          = new(httputil.Client)
		body, cttType, resp []byte
		err                 error
	)
	body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	httpClient.URI = string(body)
	httpClient.Method = "POST"
	resp, cttType, err = httpClient.RAW()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	ginResponse(resp, cttType, c)
}

// 返回文件时的快捷方法
func ginResponse(ctt []byte, cttType []byte, c *gin.Context) {
	hd := c.Writer.Header()
	if cttType != nil {
		hd.Set("Content-Type", string(cttType))
	}
	hd.Set("Pragma", "No-cache")
	hd.Set("Cache-Control", "no-cache")
	hd.Set("Expires", "0")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(ctt)
	c.Writer.Flush()
}
