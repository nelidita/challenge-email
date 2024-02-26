FROM golang:1.21.1-alpine

WORKDIR "/email-api"

COPY go.mod go.sum .

RUN go mod download

ARG PORT=8080
ARG ZS_BASE_API_URL=http://localhost:4080/
ARG ZS_INDEX=emails
ARG ZS_USER=admin
ARG ZS_PASSWORD=Complexpass#123

ENV PORT=${PORT} \
    ZS_BASE_API_URL=${ZS_BASE_API_URL} \
    ZS_INDEX=${ZS_INDEX} \
    ZS_USER=${ZS_USER} \
    ZS_PASSWORD=${ZS_PASSWORD}

COPY . .

RUN go build -o main .

CMD ["/email-api/main"]
