package utils

import (
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"unsafe"
)

var dll syscall.Handle
var initFun uintptr
var ocrFunc uintptr

//InitDll 初始化orc
func InitDll() error {
	//设置google日志级别
	os.Setenv("GLOG_minloglevel", "4")
	var err error
	dll, err = syscall.LoadLibrary("ocr.dll")
	if err != nil {
		return err
	}
	initFun, err = syscall.GetProcAddress(dll, "init")
	if err != nil {
		return err
	}
	syscall.Syscall(initFun, 0, 0, 0, 0)
	ocrFunc, err = syscall.GetProcAddress(dll, "ocr")
	if err != nil {
		return err
	}
	return nil

}

//GetCode 传入图片字节返回识别后的文本
func GetCode(img *[]byte) string {
	p := *((*int32)(unsafe.Pointer(img)))
	ret, _, _ := syscall.Syscall(ocrFunc, 2, uintptr(p), uintptr(len(*img)), 0)
	return prt2Str(ret)
}

//prt2Str 转为字符
func prt2Str(vCode uintptr) string {
	if vCode <= 0 {
		return ""
	}
	var vBytes []byte
	for {
		sByte := *((*byte)(unsafe.Pointer(vCode)))
		if sByte == 0 {
			break
		}
		vBytes = append(vBytes, sByte)
		vCode += 1
	}
	return string(vBytes)
}

//获取图片数据
func readImgData(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Println("打开验证码图片文件异常", err)
	}
	defer file.Close()
	pix, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("读取验证码图片异常", err)
	}

	tmpCode := GetCode(&pix)
	return tmpCode
}

func Free() {
	syscall.FreeLibrary(dll)
}
