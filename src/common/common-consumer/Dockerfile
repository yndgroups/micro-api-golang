# 构建：使用golang:1.19版本
FROM golang:1.19 as build

# 容器环境变量添加
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置当前工作区
WORKDIR /go/work

# 把全部文件添加到/go/work目录
ADD . .

# 编译: 把main.go编译为可执行的二进制文件, 并命名为app
RUN ls
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go mod tidy
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app testmain.go

# 运行: 使用scratch作为基础镜像
FROM scratch as prod

# 在build阶段, 复制时区配置到镜像的/etc/localtime
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 在build阶段, 复制./app目录下的可执行二进制文件到当前目录
COPY --from=build /go/work/app /
# 在build阶段, 复制yaml配置文件到当前目录, 此处需要注意调用该配置文件时使用的相对路径, testmain.go在当前目录下执行
COPY --from=build /go/work/conf/*.yaml /conf

# 启动服务
CMD ["/app"]

#FROM golang:latest
#MAINTAINER common-axuexi-epidemic-monitoring-provider
#RUN mkdir -p /worker/commonProvider
#WORKDIR /worker/commonProvider
#COPY . /worker/commonProvider
#ENV GO111MODULE=on
#ENV GOPROXY="https://goproxy.cn"
#RUN go mod download
#RUN go build testmain.go
#EXPOSE 8081
#ENTRYPOINT  ["./testmain"]