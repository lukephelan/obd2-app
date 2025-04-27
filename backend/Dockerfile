FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
RUN go build -o obd2-app ./cmd

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=build /app/obd2-app .

EXPOSE 8080

CMD ["./obd2-app"]
