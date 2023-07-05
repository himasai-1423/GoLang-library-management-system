FROM golang:1.21rc2-bookworm

COPY controllers /app/
COPY models /app/
COPY views /app/
COPY go.mod /app/
COPY go.sum /app/
COPY main.go /app/

WORKDIR /app/

RUN go get 
RUN go build -v

CMD [ "go run main.go" ]