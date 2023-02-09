# 编译
FROM golang:1.19 as builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /workspace
COPY main.go main.go 
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o main main.go

# 最小运行环境
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/main .
USER 65532:65532
ENTRYPOINT ["/main"]
