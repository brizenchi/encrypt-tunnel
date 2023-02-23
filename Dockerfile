FROM golang:1.18-alpine3.17 as base
# FROM okteto/golang:1.16 as base
ADD . /root/encrypt-tunnel
WORKDIR /root/encrypt-tunnel
ENV GOPROXY=https://goproxy.cn


RUN go mod tidy
RUN go build .
#
FROM alpine:3.17.2
#
#WORKDIR /root
ENV TZ=Asia/Shanghai
COPY --from=base /root/encrypt-tunnel/encrypt-tunnel /root/encrypt-tunnel

CMD ["/root/encrypt-tunnel"]