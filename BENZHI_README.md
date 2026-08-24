基于 Go 实现的正则表达式引擎 Web 项目，一款在线正则测试与可视化服务，支持表达式编译、文本匹配、AST 树和 NFA 图查看。

# GoRegexEngine 打包与运行说明

## 项目类型

Go 标准库 Web 服务，提供正则表达式编译、匹配和 AST/NFA 可视化接口，并附带内嵌的单页前端。

## 本地验证

```bash
go test ./...
go vet ./...
go build ./...
go run . -port 8080
```

启动后访问 `http://127.0.0.1:8080`，健康检查地址为 `/healthz`。

## Docker 打包

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh goregexengine linux/amd64
./build_benzhi_docker.sh goregexengine linux/arm64
docker run -it -p 8080:8080 goregexengine:latest
```

项目仅依赖 Go 标准库，不需要额外的 `go.sum` 或前端依赖安装。
