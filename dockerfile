FROM golang:latest

WORKDIR /go
ADD . /go

RUN go get -u github.com/pilu/fresh \
              github.com/gin-gonic/gin \
              github.com/jinzhu/gorm \
              github.com/gin-contrib/cors \
              github.com/go-sql-driver/mysql \
              firebase.google.com/go \
              google.golang.org/api/option
              
ENV GOPATH="/go/main"

CMD ["go", "run", "main.go"]
