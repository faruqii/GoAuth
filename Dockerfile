FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN ls

RUN go build -o /app

EXPOSE 3000

CMD ["/app"]

