FROM golang:1.21 as builder
#配置go代理
ARG GOPROXY="https://goproxy.cn,direct"
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

ENV GOPROXY=$GOPROXY

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux  GO111MODULE=on go build -a -o manager main.go


FROM alpine:3.11.2
WORKDIR /
COPY --from=builder /workspace/manager .
CMD ["/manager"]