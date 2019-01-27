FROM golang:latest

COPY . /go/src/github.com/reijo1337/online-library-rsoi2/auth-service/

RUN go get github.com/dgrijalva/jwt-go && \
go get github.com/lib/pq && \
go get -u github.com/gin-gonic/gin

WORKDIR /go/src/github.com/reijo1337/online-library-rsoi2/auth-service

RUN go install github.com/reijo1337/online-library-rsoi2/auth-service

CMD [ "auth-service" ]