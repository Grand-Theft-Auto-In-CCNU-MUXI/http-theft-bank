package checkpoint3

import (
	"http-theft-bank/log"
	"http-theft-bank/pkg/constvar"
	"http-theft-bank/pkg/text"
	"http-theft-bank/util"
	"strings"

	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetMethod(c *gin.Context) {
	log.Info("Message GetMethod function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")
}

func PostMethod(c *gin.Context) {
	log.Info("Message PostMethod function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")

}

// PutMethod
// @Summary user use error code to hack this gate
// @Description 站点3，用 put 方法传输 error code
// @Tags bank
// @Accept  application/json
// @Produce  application/json
// @Param object body handler.Request true "request"
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /bank/gate [put]
func PutMethod(c *gin.Context) {
	log.Info("Message PutMethod function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var data handler.Request
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if data.Content == "" {
		handler.SendBadRequest(c, errno.ErrNoCode, nil, "where's your error code?")
		return
	}

	errorCode, err := encrypt.AESDecodeAfterBase64([]byte(data.Content), []byte(viper.GetString("sercet_key")))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrDecode, nil, err.Error())
		return
	}

	if strings.Compare(string(errorCode), viper.GetString("error_code")) != 0 {
		handler.SendBadRequest(c, errno.ErrMatch, nil, "输入病毒无效")
		return
	}

	handler.SetResponseHeader(c, constvar.FragmentField, constvar.Fragment3)
	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text: text.Text3Success,
	})
}

func DelMethod(c *gin.Context) {
	log.Info("Message DelMethod function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")
}

func PatchMethod(c *gin.Context) {
	log.Info("Message PatchMethod function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")
}
