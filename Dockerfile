FROM arm64v8/golang:alpine

ENV TOKEN 1234

# Intall Flume and dependancies
RUN go get github.com/dimensi0n/flume && \
    cd /go/src/github.com/dimensi0n/flume && \
    go build

EXPOSE 8080
CMD flume -t $TOKEN
