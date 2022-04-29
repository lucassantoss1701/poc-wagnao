FROM golang:1.16 as builder 
WORKDIR /

RUN apt-get update && apt-get install -y wget
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . . 

RUN go build -o /publisher


FROM golang:1.16  

COPY --from=builder /publisher /publisher

RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz




