FROM golang:1.21 as builder

WORKDIR /app

#Copy source code for golang app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o chunker cmd/nfx-file-chunker/main.go
FROM debian:latest
RUN apt-get -y update && apt-get -y upgrade && apt-get install -y ffmpeg

COPY --from=builder /app/chunker /app/chunker
COPY --from=builder /app/config.yaml /app/config.yaml

#Copy aws configuration
RUN mkdir /root/.aws
COPY .aws/config /root/.aws
COPY .aws/credentials /root/.aws
EXPOSE 8080

WORKDIR /app
RUN mkdir data

CMD ["./chunker"]



