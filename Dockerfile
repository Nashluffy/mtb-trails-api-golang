############################
# STAGE 1
############################
FROM golang:alpine AS builder

# Downloads all certificates
RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

# Set our working directory
RUN mkdir -p /build
WORKDIR /build

# Copy all files over
COPY . . 

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o mtb-trails-api-golang
############################
# STAGE 2
############################
FROM scratch

# Copy all certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copies the built binary over from the previous stage
COPY --from=builder /build/mtb-trails-api-golang /go/bin/mtb-trails-api-golang
# Document that the service listens on port 8080.
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/go/bin/mtb-trails-api-golang"]

