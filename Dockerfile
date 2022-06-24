FROM golang:1.16-alpine

WORKDIR /app

ADD go ./

#COPY go.mod ./
#COPY go.sum ./
RUN go mod download


# COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 80

CMD [ "/docker-gs-ping" ]