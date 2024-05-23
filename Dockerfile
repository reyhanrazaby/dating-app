FROM golang:1.21-alpine3.18 as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Set the working directory to /app
WORKDIR /app

# Copy the project directory contents into WORKDIR
COPY . .

RUN cd /app
RUN go mod download
RUN go mod verify

# go build command with the -ldflags="-w -s" option to produce a smaller binary file by stripping debug information and symbol tables.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o bin main.go

FROM alpine:3.18

RUN apk update
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy the executable.
COPY --from=builder /app/bin /app

ENTRYPOINT ["/app"]
