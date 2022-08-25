FROM golang:alpine AS builder
RUN mkdir /goBark
# 把当前目录下的所有内容拷到上面创建的goBark目录
ADD . /goBark/   
# 设定工作目录为goBark   
WORKDIR /goBark
# 编译cmd目录下的goBark.go 文件，输出为goBark
## CGO_ENABLED=0 go build 打包避免交叉编译
## -o output 指定编译输出的名称，代替默认的包名
RUN CGO_ENABLED=0 go build -o goBark cmd/goBark.go

FROM gcr.io/distroless/base
# 把上面build中生成的/goBark/goBark编译结果文件 拷到当前根目录下的存为 /goBark文件
COPY --from=builder /goBark/goBark /goBark
ENTRYPOINT ["/goBark"]
