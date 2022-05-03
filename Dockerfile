FROM golang:1.16-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src

COPY . .
RUN go mod download
RUN go build -o app ./cmd/app

FROM busybox

COPY --from=build /go/src/app /go/src/app
COPY --from=build /go/src/config /go/src/config
COPY --from=build /go/src/.keys /go/src/.keys

EXPOSE 8080

CMD ["/go/src/app"]
