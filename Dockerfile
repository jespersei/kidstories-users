FROM golang:alpine

ADD . /go/src/github.com/gin-gonic/gin
ADD . /go/src/github.com/itsjamie/gin-cors
ADD . /go/src/golang.org/x/crypto/bcrypt
ADD . /go/src/gopkg.in/appleboy/gin-jwt.v2
ADD . /go/src/gopkg.in/mgo.v2
ADD . /go/src/gopkg.in/mgo.v2/bson

ADD . /go/src/github.com/kidstories/users
RUN go install github.com/kidstories/users
CMD ["/go/bin/gin-mongo-api"]
EXPOSE 7324
