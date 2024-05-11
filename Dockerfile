FROM golang:alpine

WORKDIR /GoLibraryAPI
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api \
    && go build -o ./bin/migrate ./cmd/migrate

CMD ["/GoLibraryAPI/bin/api"]
EXPOSE 3000