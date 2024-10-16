FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o ./build ./cmd/app/

EXPOSE 8080

CMD [ "./build" ]