package checkpoint5

import (
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/util"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UploadFile ... 上传全排列.go 文件，线上 ac
// @Summary 站点5 上传文件 ac
// @Description 站点5，上传文件
// @Tags end
// @Accept  application/json
// @Produce  application/json
// @Param file body string true "file,这个使用表单！！！"
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /muxi/backend/computer/examination [post]
func UploadFile(c *gin.Context) {
	log.Info("Message UploadFile function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// get file
	file, err := c.FormFile("file")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrFormFile, nil, err.Error())
		return
	}

	// 根据时间戳生成新的文件名
	fileName, fileNameOnly := newFileName(file.Filename)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "./file/"+fileName)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrSaveFile, nil, err.Error())
		return
	}

	// ac
	err = testProgramme(fileName, fileNameOnly)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrTestProgramme, nil, err.Error())
		return
	}

	// response
	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text: text.Text5Success,
	})
	return
}

func newFileName(fileName string) (string, string) {
	// 根据时间戳生成新的文件名
	extString := path.Ext(fileName)                           // 获取后缀
	fileNameTimeInt := time.Now().Unix()                      // 获取时间戳
	fileNameTimeStr := strconv.FormatInt(fileNameTimeInt, 10) // 时间戳格式化
	filenameOnly := strings.TrimSuffix(fileName, extString)   // 去掉原文件名后缀

	return filenameOnly + fileNameTimeStr + extString, filenameOnly + fileNameTimeStr
}
