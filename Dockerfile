FROM golang:1.18.2 as builder
ARG buildDir=/Go/src/github.com/skrbox/ioseek
ARG goProxy
WORKDIR ${buildDir}
ENV GOPATH=/Go
ENV GOOS=linux
ENV GOPROXY=${goProxy}
COPY . .
RUN make pkg

FROM alpine:3.15.4 as runner
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
