FROM registry.gitlab.com/qio/standard/golang:1.17

WORKDIR /api

COPY . .
RUN go mod tidy

RUN chmod +x /api/runner.sh
ENTRYPOINT /api/runner.sh
