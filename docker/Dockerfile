FROM golang:latest

WORKDIR /go/src

ADD ./ /go/src

RUN go build

EXPOSE 8080

ENTRYPOINT ["./GoPush"]
