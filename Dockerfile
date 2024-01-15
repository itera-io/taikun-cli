FROM golang:1.20-alpine AS build

WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /taikun .

FROM scratch AS bin
COPY --from=build /taikun /