FROM golang:1.20.3-alpine

WORKDIR /app/
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./bin/ ./...

CMD ["/app/bin/doctor-assistant-app"]