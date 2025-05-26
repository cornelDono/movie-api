FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p tmp
RUN go install github.com/air-verse/air@latest

ENV PATH="/go/bin:${PATH}"

EXPOSE 4000

ENTRYPOINT ["air"]