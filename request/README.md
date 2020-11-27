## http请求

* 创建一个请求
> http.NewRequest

* 设置请求参数
> url.Values

* 设置请求头
> request.Header.Add("user-agent","chrome")

* 绑定请求和参数
> request.Url.RawQuery = params.Encode()

* 发起请求
> http.DefaultClient.Do()

* user-agent github 有项目定期更新

## http响应

* 响应需要关注
> body、status、header（响应头）、encoding

* 获取响应的编码信息
    * 可以通过Content-Type获取
    * 或者html head meta
    * 通过网页头部猜测编码信息
    > golang.org/x/net/html --> charset.DetermineEncoding()
    * 预取网页的前1024个字节
    > bufio.NewBuffer() --> reader.Peek(1024) 这里读取不会移动读取位置

* 如果不是utf-8编码，则需要转码
> golang.org/x/text/transform  
```go
transform.NewReader()
```

## 下载文件

* io.Copy 把下载的数据拷贝到文件中
// todo

## 发送复杂的post请求

* 设置content-type
    * application/json  （通过结构体 或者 map 构建json数据）
    * application/x-www-urlencoded( 设置参数 url.Values )
 
* 上传文件
```go
// 定义一个body
body := &Reader{}
// 创建一个multpart writer
writer :=  multpart.NewWriter()
// 写入表单内容
writer.WriteField("name","b")
// 创建表单文件
uploadWriter := writer.CreateFormFile("表单名","文件名")
// 写入文件数据
io.Copy(uploadWriter, file)
// 关闭writer
writer.Close()
// 从writer中获取Content-Type
writer.FormDataContentType()
```

## 重定向

* 状态码 3xx

* 限制重定向的次数
> 需要自己新建一个client并且自定义CheckRedirect函数，判断via的长度

## cookie

* cookieJar

## transport 设置超时细粒度的超时和代理等

* http代理

* shadowsocks代理（socks5）

# tinyproxy
* DisableViaHeader
> 隐藏代理请求的 "Via" 字段

* ViaProxyName
> ViaProxyName 字段定义的名称会出现在请求头的 "Via" 字段中，需要自定义

* Allow 127.0.0.1
> 代表只允许对本地请求进行代理。

* XTinyproxy Yes
> Yes：表示代理服务器请求真实服务器的时候会带上客户端真实的IP，所以需要设置为 "no" 或者 注释掉。


