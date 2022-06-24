FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
# IMport static files!
COPY static/* static/ 

RUN go mod download

COPY *.go ./

RUN go build -o /go-webserver

CMD [ "/go-webserver" ]