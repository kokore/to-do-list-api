FROM golang:1.18


RUN apt update

ENV ROOT_PATH=/var/www/to-do-list

RUN mkdir -p $ROOT_PATH

WORKDIR $ROOT_PATH

COPY . .

RUN go mod tidy

RUN go build main.go

EXPOSE 3000

CMD go run ./main.go
