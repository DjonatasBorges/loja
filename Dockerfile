FROM golang:1.20.6

WORKDIR /app

COPY . .

RUN go get
RUN go build -o bin

ENTRYPOINT [ "/app/bin" ]