FROM golang:alpine as builder
WORKDIR /src/app
COPY . ./
RUN go build -o shorturl

FROM alpine
WORKDIR /root/
COPY --from=builder /src/app ./app
CMD ["./app/shorturl"]