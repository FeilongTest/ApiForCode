package services

import (
	"ApiForCode/app/common/request"
	"ApiForCode/app/model"
	"ApiForCode/utils"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
)

func GetCode(params request.Code) (codeResult model.CodeResult, err error) {
	//处理base64
	if strings.Index(params.Img, "data:image/png;base64,") != -1 {
		params.Img = strings.Replace(params.Img, "data:image/png;base64,", "", 1)
	}
	//base64转图片
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(params.Img))
	imgByte, err := io.ReadAll(dec)
	if err != nil {
		log.Println("转换图片出错", err)
		return codeResult, fmt.Errorf("%s", "验证码图片转换出错")
	}
	codeResult.Result = utils.GetCode(&imgByte)
	return codeResult, err
}

func GetBase64(img []byte) (codeResult model.CodeResult, err error) {
	codeResult.Result = base64.StdEncoding.EncodeToString(img)
	return codeResult, err
}
