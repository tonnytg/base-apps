FROM ubuntu:24.04 as core
WORKDIR /app
COPY ./app/ .
RUN apt-get update && apt-get install -y curl wget telnet nginx certbot

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN mkdir -p /go /go/src /go/bin /go/pkg
ENV GOPATH=/go
ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:$GOBIN
ENV GOROOT=/usr/local/go

RUN go build -o main .

# COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 4001:4001
EXPOSE 80:4080
EXPOSE 443:4443

ENV PORT=4001

CMD ["./main"]