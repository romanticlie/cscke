package resource

import (
	"github.com/gin-gonic/gin"
	"cscke/internal/response"
)

func Upload(c *gin.Context) {

	form, err := c.MultipartForm()

	if err != nil {
		response.Msg(c, err.Error())
		return
	}

	files,ok := form.File["files"]

	if !ok {
		response.Msg(c, "files field is not exists")
		return
	}

	for _, file := range files {
		if err := c.SaveUploadedFile(file, "/www/"+file.Filename); err != nil {
			response.Msg(c,err.Error())
			return
		}
	}

	response.OK(c)
}
