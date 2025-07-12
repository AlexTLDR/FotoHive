FROM node:latest AS tailwind-builder
WORKDIR /tailwind
COPY ./tailwind/package.json ./tailwind/package-lock.json ./
RUN npm install
COPY ./templates /templates
COPY ./tailwind/tailwind.config.js ./tailwind.config.js
COPY ./tailwind/styles.css ./styles.css
RUN npx tailwindcss -c ./tailwind.config.js -i ./styles.css -o /styles.css --minify

FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./server ./cmd/server/

FROM alpine
WORKDIR /app
COPY ./assets ./assets

COPY --from=builder /app/server ./server
COPY --from=tailwind-builder /styles.css /app/assets/styles.css
CMD ["./server"]
