FROM golang:1.13
WORKDIR /go/src/github.com/idsulik/go-template-renderer/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/idsulik/go-template-renderer/app .
COPY --from=0 /go/src/github.com/idsulik/go-template-renderer/server/template.html ./server/template.html
CMD ["./app"]