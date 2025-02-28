FROM golang:1.23-alpine AS build
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v ./cmd/shorts

FROM alpine AS production

COPY assets assets
COPY --from=build /usr/src/app/shorts .

EXPOSE 8080

CMD [ "./shorts" ]
