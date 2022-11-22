# 多阶段构建
#构建一个 builder 镜像，目的是在其中编译出可执行文件mck
#构建时需要将此文件放到代码根目录下
FROM golang:alpine  as builder
#将上层整个文件夹拷贝到/build
ADD . /build/src
#去掉了调试信息 -ldflags="-s -w" 以减小镜像尺寸
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go go build -ldflags="-s -w"  -o alert-manager /build/src/cmd/alert-manager/main.go

FROM alpine
#更新软件源  使用挂载卷挂载 /etc/localtime:/etc/localtime做宿主机时间映射
#RUN apk update --no-cache && apk add --no-cache tzdata \
# && apk add --no-cache docker-cli
##设置本地时区，这样我们在日志里看到的是北京时间了
#ENV TZ Asia/Shanghai

WORKDIR /home

COPY --from=builder  /build/src/config/config.yml /home/config/config.yml
COPY --from=builder /build/src/alert-manager /home/alert-manager

CMD ["./alert-manager"]
EXPOSE 9680

#docker run  -d --name publish -p 9680:9680 -v /etc/localtime:/etc/localtime publish
