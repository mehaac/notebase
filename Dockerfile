FROM golang:1.24-alpine AS backend
WORKDIR /app/
COPY . .
RUN go build -o bin/notebase .
RUN GOOS=linux GOARCH=arm64 go build -o bin/notebase-arm64 .

FROM node:22-alpine AS frontend
WORKDIR /app
RUN corepack enable
COPY ./web/package.json ./web/pnpm-lock.yaml ./
RUN pnpm install
COPY ./web ./
RUN pnpm run generate-prod

FROM alpine:3.19
WORKDIR /app/

ENV NOTES_ROOT=/tmp/example/notes
ENV SUPERUSER_EMAIL=
ENV SUPERUSER_PASSWORD=

COPY --from=backend /app/bin/notebase .
COPY --from=backend /app/bin/notebase-arm64 .
COPY --from=frontend /app/.output/public/ ./pb_public
COPY ./example/ /tmp/example/
COPY entrypoint.sh .

RUN chmod +x /app/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]

CMD ["serve", "--http=0.0.0.0:8080", "--dev"]
