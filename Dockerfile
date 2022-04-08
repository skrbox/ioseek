FROM --platform=linux/amd64 golang:1.17.7 as builder
ARG buildDir=/Go/src/github.com/skrbox/ioseek
WORKDIR ${buildDir}
ENV GOPATH=/Go
ENV GOOS=linux
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN make pkg

FROM --platform=linux/amd64 alpine:3.15.0 as runner
ARG commitId
LABEL author=Jeyrce.Lu<jeyrce@gmail.com> \
      poweredBy=https://github.com/skrbox/ioseek \
      commitId=${commitId}
ARG buildDir=/Go/src/github.com/skrbox/ioseek
ARG pkgDir
WORKDIR /app
COPY --from=builder --chown=bin ${buildDir}/_output/${pkgDir}/ .
EXPOSE 80
CMD ["/app/ioseek"]
