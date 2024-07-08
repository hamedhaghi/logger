FROM golang:1.22.4-bullseye

LABEL author="Hamed Haghi" email="hamedhaghi07@gmail.com"

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl && \
    go install github.com/pilu/fresh@latest

WORKDIR /app

COPY ./notifier /app

RUN go mod tidy

EXPOSE 3000

CMD ["fresh"]
