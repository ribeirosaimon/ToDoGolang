FROM golang as builder

WORKDIR /build/api
COPY go.mod ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o api
# post build stage
FROM alpine
WORKDIR /root
COPY --from=builder /build/api/api .
EXPOSE 3000
CMD ["./api"]