FROM ubuntu:18.04

# install golang
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install wget -y
RUN wget https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz
RUN tar -xvf go1.11.2.linux-amd64.tar.gz
RUN mv go /usr/local

# set path for golang
ENV GOPATH=$HOME/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

RUN mkdir $HOME/go

# package needed to use exif
RUN apt-get install pkg-config -y
RUN apt-get install git -y

RUN apt-get install libexif-dev -y
RUN go get github.com/xiam/exif

RUN go get github.com/sendgrid/sendgrid-go

# utils library
RUN go get github.com/cvhariharan/Utils

RUN go get github.com/fatih/set
RUN go get github.com/go-ego/riot

WORKDIR $HOME/go/src/app
COPY . .

RUN mkdir temp-images
RUN go build -o main

EXPOSE 7500
RUN ls
CMD ["./main"]