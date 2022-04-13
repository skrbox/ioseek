FROM --platform=linux/amd64 golang:1.17.7-alpine3.15 as builder
ARG buildDir=/Go/src/github.com/skrbox/ioseek
ARG goProxy=""
WORKDIR ${buildDir}
ENV GOPATH=/Go
ENV GOOS=linux
ENV GOPROXY=${goProxy}
COPY . .
RUN apk add gcc g++ make cmake gfortran libffi-dev openssl-dev libtool && make pkg

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
