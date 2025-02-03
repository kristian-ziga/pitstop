FROM golang:1.23

WORKDIR /app

COPY ./api ./

RUN go mod download

RUN go build -o /bin/app

EXPOSE 8000
RUN chmod +x /bin/app

CMD ["/bin/app"]
