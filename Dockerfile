FROM  golang:1.18.1-alpine
WORKDIR /app
COPY . .
ENV TZ "Asia/Shanghai"
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build  ./cmd/main.go \
    && chmod +x main \
	&& apk -U upgrade \
	&& apk add tzdata \
	&& cp /usr/share/zoneinfo/${TZ} /etc/localtime && echo ${TZ} > /etc/timezone

EXPOSE 80
ENTRYPOINT ["./main"]