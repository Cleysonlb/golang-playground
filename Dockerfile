FROM golang:alpine
ADD ./src /go/src
WORKDIR /go/src
ENV PORT=3001
CMD ["go", "run", "main.go"]