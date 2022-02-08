FROM golang:latest as builder
LABEL maintainer="Carlos Gamboa cegamboav@gmial.com"
WORKDIR     /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV PORT 8081
RUN go build -o ./prog_exercise
CMD ["./prog_exercise"]
