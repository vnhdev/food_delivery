package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-services-g04/common"
	"simple-services-g04/component"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
