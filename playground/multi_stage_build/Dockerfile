FROM golang:1.22
WORKDIR /work/app
COPY main.go go.mod ./
RUN go mod tidy
RUN go build .
CMD ["./main"]

# FROM golang:1.22 AS builder
# WORKDIR /work/app
# COPY main.go go.mod ./
# RUN go mod tidy
# RUN go build .

# FROM gcr.io/distroless/static:nonroot
# COPY --from=builder /work/app ./
# CMD ["./main"]
