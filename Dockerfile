FROM golang:1.16-alpine AS backend
#ENV GOPROXY="https://proxy.golang.org"
ENV GOPROXY="https://goproxy.cn,https://goproxy.io,direct"
COPY . /build
WORKDIR /build
# swaggo
#RUN GO111MODULE="on" GOPROXY=$GOPROXY go get -u github.com/swaggo/swag/cmd/swag
#RUN swag init --parseDependency --parseInternal
# build
RUN go mod tidy
RUN GO111MODULE="on" GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
RUN sed -i "s@https://dl-cdn.alpinelinux.org/@https://mirrors.huaweicloud.com/@g" /etc/apk/repositories
RUN apk update && apk add upx
RUN upx dkdk

FROM node:14-slim AS frontend
ENV NPM_REGISTRY https://mirrors.huaweicloud.com/repository/npm/
ENV NPM_REGISTRY https://registry.npm.taobao.org
COPY frontend /dkdk
WORKDIR /dkdk
RUN npm config set registry $NPM_REGISTRY && \
npm cache clean -f && \
npm install
RUN npm run-script build

FROM alpine
ENV WAIT_VERSION 2.7.3
# ENV WAIT_RELEASE https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait
ENV WAIT_RELEASE https://st0n3-dev.obs.cn-south-1.myhuaweicloud.com/docker-compose-wait/release/$WAIT_VERSION/wait
ADD $WAIT_RELEASE /wait
RUN chmod +x /wait

RUN mkdir -p /app
COPY --from=backend /build/dkdk /app/
COPY --from=frontend /dkdk/dist /app/dist
WORKDIR /app
CMD /wait && ./dkdk