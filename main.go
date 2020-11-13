package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// 啟動telegram bot
	newBot()

	// 啟動gin
	r := gin.New()
	r.POST("api/test", test)
	r.POST("api/photo-upload", photoUpload)
	r.POST("api/file-upload", fileUpload)
	r.Run(":8866")
}

func fileUpload(c *gin.Context) {
	// request 參數
	file, _, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// 新建立一個空的file，放在output
	output, err := os.Create("report.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	// 將從request拿到的file，copy到output
	_, err = io.Copy(output, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, "upload successful \n")

	go sendFile("report.csv")
}

func photoUpload(c *gin.Context) {
	sendPhoto("rv.jpeg")
}
