FROM golang:alpine AS build

WORKDIR /app
COPY . .

ENV GO111MODULE=on

RUN go mod download

RUN go build -o server ./main.go

FROM golang:alpine AS server

WORKDIR /app

COPY --from=build /app .
# COPY --from=build /src/app/server .
EXPOSE 8080 443

CMD ./server cl