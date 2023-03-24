FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN ["go", "build", "."]


# Second Stage

FROM alpine

COPY --from=builder /app/go-app .

CMD ["./go-app"]