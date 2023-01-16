FROM golang:1.19.5

LABEL MAINTAINER = "Ali Vatandoost"

RUN mkdir /home/todoapp
WORKDIR /home/todoapp

COPY . .
COPY ./templates ../dist/templates
COPY ./static ../dist/static

RUN go mod download
RUN go build -o ../dist

WORKDIR /home/dist
RUN rm -r /home/todoapp
ENTRYPOINT ["./todo"]