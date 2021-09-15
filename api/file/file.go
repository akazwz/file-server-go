package file

import (
	"github.com/akazwz/file-server/model/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
	"time"
)

func CreateFile(c *gin.Context) {
	fullPath := c.FullPath()
	log.Println(fullPath)

	file, err := c.FormFile("file")
	if err != nil {
		response.CommonFailed("get file error", 4001, c)
		return
	}

	dirDate := time.Now().Format("2006-01-02")

	dir := "public/file/" + dirDate + "/"

	_, err = os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			response.CommonFailed("create dir error", 4002, c)
			return
		}
	}

	fileNamePrefix := uuid.NewV4().String()
	fileName := fileNamePrefix + "-" + file.Filename

}
