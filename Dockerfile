
FROM golang:1.22.5-alpine3.19 AS build

# Buildx build-in ARGs
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""

# Additional build ARGs passed from --build-args
ARG APPLICATION_NAME="example"
ARG VERSION
ARG SHA

# Environment variables used at compile time by Golang
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=${TARGETOS} \
  GOARCH=${TARGETARCH} \
  GOARM=${TARGETVARIANT}

WORKDIR /go/src/github.com/xunholy/secure-software-supply-chain-demo/

# Disabled due to no modules
# COPY go.mod go.sum ./

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo \
  -ldflags="-w -extldflags '-static' -X 'main.ApplicationName=${APPLICATION_NAME}' -X 'main.Version=${VERSION}' -X 'main.SHA=${SHA}'" \
  -o example .

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=build --chown=nonroot /go/src/github.com/xunholy/secure-software-supply-chain-demo/example .

USER nonroot:nonroot

ENTRYPOINT ["/example"]
