## useful-tools-golang
一些工作生活中使用的小工具
打包说明
```shell
# mac os build .exe file
CGO_ENABLED=0  GOOS=windows  GOARCH=amd64 go build main.go
```

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