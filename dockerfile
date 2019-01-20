FROM golang:latest

WORKDIR /hubb
ADD . /hubb

RUN go install github.com/pilu/fresh
RUN go build main/hubb.go

CMD ["./hubb"]
