FROM golang:1.8

RUN mkdir -p /go/src/github.com/ederavilaprado/golang-web-architecture-template
WORKDIR /go/src/github.com/ederavilaprado/golang-web-architecture-template
COPY . /go/src/github.com/ederavilaprado/golang-web-architecture-template

# RUN go get github.com/tools/godep
RUN make deps
RUN make build


# CMD ["./server"]
# EXPOSE 8080

# docker run -it gowebtemplate:v0.1.0 /bin/bash
