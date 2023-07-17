FROM golang:1.20.6

WORKDIR /app

ENV POSTGRES_DNS=""

ENV APP_LEVEL="DEBUG"

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /time-coster

CMD ["/time-coster"]
