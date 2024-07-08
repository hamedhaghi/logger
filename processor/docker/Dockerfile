FROM golang:1.22.4-bullseye

LABEL author="Hamed Haghi" email="hamedhaghi07@gmail.com"

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl && \
    go install github.com/pilu/fresh@latest

WORKDIR /app

COPY ./processor /app

RUN go mod tidy

CMD ["fresh"]
