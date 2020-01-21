FROM golang:alpine
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -v
CMD ["/app/realtime-chat-go-react"]