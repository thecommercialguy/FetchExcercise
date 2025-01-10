FROM debian:stable-slim
# FROM golang1.23.1-alpine AS builder
# # COPY source destination
COPY FetchExcercise.git /bin/FetchExcercise.git

# RUN go test ./...

ENV PORT=8080

CMD ["/bin/FetchExcercise.git"]