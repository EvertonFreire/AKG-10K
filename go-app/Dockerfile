FROM golang:1.14

WORKDIR /opt/go/application/
COPY ./ /opt/go/application/

RUN go get -u github.com/gorilla/mux

CMD ["go","run","main.go"]