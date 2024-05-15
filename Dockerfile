FROM golang:1.22.3-alpine
WORKDIR /app
COPY /app/* ./
RUN apk update && apk add git
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
RUN go install github.com/cosmtrek/air@latest
RUN go mod download
CMD ["air","-c", "air.toml"]