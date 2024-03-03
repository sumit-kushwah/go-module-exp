FROM golang:latest

WORKDIR /app

ARG mode

COPY go.mod .
COPY main.go .

RUN if [ "$mode" = "dev" ] ; then \
        go get github.com/amit-rajawat/my-library@master ; \
    else \
        go get github.com/amit-rajawat/my-library@prod ; \
    fi

RUN go mod tidy
RUN go mod vendor
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /main main.go

CMD [ "/main" ]