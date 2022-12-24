FROM golang:alpine AS build

WORKDIR /src/app
COPY . .  

ENV GO111MODULE=on
RUN go mod download

RUN go build -o server .

FROM golang:alpine AS server

ENV PORT=8080
WORKDIR /app
COPY --from=build /src/app .
# COPY --from=build /src/app/server .

EXPOSE 8080

CMD ./server