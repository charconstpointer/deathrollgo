FROM golang:1.14

COPY . /app
WORKDIR /app

ENV GO111MODULE=on

RUN apt-get update && apt-get install build-essential -y
RUN cd cmd/dr && GOOS=linux go build -o deathroll
EXPOSE 8080
ENTRYPOINT [ "./cmd/dr/deathroll" ]  