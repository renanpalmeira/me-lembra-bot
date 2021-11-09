FROM golang:alpine
COPY . /code
WORKDIR /code
RUN go build -ldflags="-s -w" -o ./app ./cmd/api
CMD ["/code/app"]
