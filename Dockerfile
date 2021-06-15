############################
# STAGE 1
############################
FROM golang AS builder

# Set our working directory
WORKDIR /go/src/github.com/Nashluffy/mtb-trails-api-golang

# Copy all files over
COPY . . 

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/mtb-trails-api-golang .

############################
# STAGE 2
############################
FROM scratch

WORKDIR /go/bin

# Copies the built binary over from the previous stage
COPY --from=builder /go/bin/mtb-trails-api-golang /go/bin/mtb-trails-api-golang

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/go/bin/mtb-trails-api-golang"]

