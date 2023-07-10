FROM golang:1.21rc2-bookworm

WORKDIR /app/

RUN mkdir -p /home/app

COPY . /app/

RUN go get 

CMD [ "./" ]