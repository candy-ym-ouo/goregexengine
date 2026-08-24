# 官方 Go 镜像保留完整工具链，便于评测容器内继续编译和测试。
FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build ./...

CMD ["bash"]
