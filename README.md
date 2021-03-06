# ApiForCode
Golang 基于Gin框架的 Ocr验证码本地识别Api

1. 实现Golang调用Ocr.dll（仅Windows 32位使用）通过设置日志级别可忽略dll调用的报错
   
```
	//设置google日志级别
	os.Setenv("GLOG_minloglevel", "4")
```
 
2. 64位的Golang环境须使用32位运行 在Goland中可在运行配置-环境 添加如下：
    
```
    GOARCH=386
```
    
3. 实现图片验证码识别 提交图片的Base64编码即可
    
```
POST /api/ocr HTTP/1.1
Host: 127.0.0.1:9797
Content-Type: text/plain
Content-Length: 7410

{
    "img":"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"
}
```

4. 实现图片文件上传转Base64
        
```
POST /api/img2base64 HTTP/1.1
Host: 127.0.0.1:9797
Content-Length: 193
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="img"; filename="/C:/Users/Test/Desktop/img.png"
Content-Type: image/png

(data)
----WebKitFormBoundary7MA4YWxkTrZu0gW

```

5. 64位编译程序

```
C:\Users\Test\Desktop>cd /d D:

D:\>cd "Source Demo\go\ApiForCode

D:\Source Demo\go\ApiForCode>go env GOARCH
amd64

D:\Source Demo\go\ApiForCode>set GOARCH=386

D:\Source Demo\go\ApiForCode>go env GOARCH
386

D:\Source Demo\go\ApiForCode>go build

```

