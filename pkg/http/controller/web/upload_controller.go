package web

import (
	"github.com/gin-gonic/gin"
	"leinao-os-adapter/pkg/global/consts"
	"leinao-os-adapter/pkg/global/variable"
	"leinao-os-adapter/pkg/service/upload_file"
	"leinao-os-adapter/pkg/utils/response"
)

type Upload struct {
}

//  文件上传是一个独立模块，给任何业务返回文件上传后的存储路径即可。
// 开始上传
func (u *Upload) StartUpload(context *gin.Context) {
	savePath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")
	if r, finnalSavePath := upload_file.Upload(context, savePath); r == true {
		response.Success(context, consts.CurdStatusOkMsg, finnalSavePath)
	} else {
		response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg, "")
	}
}
