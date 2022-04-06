FROM golang:1.7.4-alpine

COPY . /home/ubuntu/test2

RUN export GOBIN=$GOPATH/bin && cd /home/ubuntu/test2 && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT test2
