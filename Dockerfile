FROM alpine AS base
RUN apk add --no-cache curl wget

FROM golang:1.13 AS go-builder
WORKDIR /go/app
COPY . /go/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/app/forum-api /go/app/src/forum-api.go

FROM base
COPY --from=go-builder /go/app/forum-api /forum-api
COPY --from=go-builder /go/app/forum-settings.json /forum-settings.json

CMD ["/forum-api"]

