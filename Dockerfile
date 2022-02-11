FROM --platform=linux/amd64 golang:1.17.7 as builder
ARG buildDir=/Go/src/github.com/pmeta/ioseek
WORKDIR ${buildDir}
ENV GOPATH=/Go
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN make pkg

FROM --platform=linux/amd64 alpine:3.15.0 as runner
ARG buildDir=/Go/src/github.com/pmeta/ioseek
ARG pkgDir
WORKDIR /app
COPY --from=builder ${buildDir}/_output/${pkgDir}/ .
EXPOSE 80
CMD ["/app/ioseek", "--meta.config-file=/app/ioseek.yml"]
