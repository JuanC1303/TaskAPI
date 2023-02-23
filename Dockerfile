FROM golang

WORKDIR /go/src/taskapi
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
RUN go build -o /docker-gs-ping

EXPOSE 1303

CMD ["/docker-gs-ping"]