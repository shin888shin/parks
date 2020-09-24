FROM golang:latest as builder
WORKDIR /go/src/github.com/shin888shin/parks
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/shin888shin/parks/main .
EXPOSE 8082
CMD ["./main"]