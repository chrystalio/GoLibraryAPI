FROM golang:alpine

WORKDIR /GoLibraryAPI
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/GoLibraryAPI/bin/api"]
EXPOSE 3000