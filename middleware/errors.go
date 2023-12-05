package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GlobalErrorData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type GlobalErrorDatas []*GlobalErrorData

func GlobalErrorHandle(ctx *gin.Context) {
	ctx.Next()
	ctx.Abort()
	errs := make(GlobalErrorDatas, 0)
	for _, v := range ctx.Errors {
		errs = append(errs, &GlobalErrorData{
			Code: strconv.Itoa(int(v.Type)),
			Msg:  v.Error(),
			Data: v.Meta,
		})
	}

	if len(errs) != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"errs": errs,
		})
	}
}
