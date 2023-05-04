## useful-tools-golang
**一些工作生活中使用的小工具**

打包命令:
```shell
# mac os build .exe file
CGO_ENABLED=0  GOOS=windows  GOARCH=amd64 go build edit_filename.go
```
> 参数说明：
> - CGO_ENABLED : CGO 表示golang中的工具，CGO_ENABLED 表示CGO禁用，交叉编译中不能使用CGO的
> - GOOS : 目标平台
>   - mac 对应 darwin
>   - linux 对应 linux
>   - windows 对应 windows
> - GOARCH ：目标平台的体系架构【386，amd64,arm】, 目前市面上的个人电脑一般都是amd64架构的
>   - 386 也称 x86 对应 32位操作系统
>   - amd64 也称 x64 对应 64位操作系统
>   - arm 这种架构一般用于嵌入式开发。 比如 Android , IOS , Win mobile , TIZEN 等

#### 根据文件修改时间，批量更新指定目录下文件名称
- application/edit_filename.go 

#### 生成snowflake id(分布式全局唯一ID)
- application/snowflake.go 

#### chatGPT api集成
- application/chatpgt3.go 
> 条件： 
> - 注册openai账号[chat openai](https://chat.openai.com/chat)
> - 生成token，参考文章[openai docs](https://platform.openai.com/docs/quickstart/build-your-application)
> - [github go-openai](https://github.com/sashabaranov/go-openai)