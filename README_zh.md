<!-- omit in toc -->
# LowFrame
用于快速构建简易项目的 MVC 模板框架.
适合没有精力的非专业程序员, 解决内网零散紧急需求使用.

基于流行的, 活跃的三方库实现以下主要功能:
- 多路路由
- render
- js
- 模板
- ...

# 代码组织

```
├── main.go                 # 主入口
├── settings.yaml           # 主配置文件
├── cmd                     # 命令行入口
│   ├── config.go           # init 加载配置
│   ├── root.go
│   └── server.go           # 调用 server 启动服务
├── config                  # 配置文件定义
│   ├── cfg.go
│   └── server.go
├── server                  # 服务入口
│   ├── router.go           # APP 与路由关联
│   └── server.go           # http.Server 命令
├── app                     # 业务代码
│   ├── index               # index 入口
│   │   ├── api_router.go
│   │   ├── web_handler.go
│   │   └── web_router.go
│   ├── public              # 公共 handler
│   │   └── api_handler.go  # 有一个占位用的 handler
│   ├── subroute.go         # 子路由结构体定义
│   └── user                # 用户业务入口
│       ├── api_handler.go  # 用户 api 方法
│       ├── api_router.go   # 用户 api 路由
│       ├── model.go        # 用户结构体定义
│       ├── web_handler.go  # 用户 模板渲染 方法
│       └── web_router.go   # 用户 模板渲染 路由
└── templates               # 模板
    ├── template.go         # 模板 render 初始化
    ├── assets              # 样式文件, js 文件
    └── templates           # 模板文件


```


# 样例实现的功能：
- 静态文件服务
- 简单 js 访问
- 提交表单并输出表格
- todo接口占位
- 三方接口调用

# 特别感谢
- [warpmatrix](https://github.com/warpmatrix/go-web)
- gorilla/mux
- urfave/negroni
- unrolled/render
- go-resty/resty
- https://milligram.io/#getting-started
