FROM golang:1.24-alpine AS backend
WORKDIR /app/
COPY . .
RUN go build -o bin/notebase .

FROM alpine:3.19
WORKDIR /app/
COPY --from=backend /app/bin/notebase .
COPY ./pb_public/ .
ENTRYPOINT ["./notebase"]
