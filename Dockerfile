FROM golang

WORKDIR /backend_app

COPY . .

RUN go mod download

ENV DB_PORT localhost

EXPOSE 8000

CMD ["go", "run", "Main.go"]