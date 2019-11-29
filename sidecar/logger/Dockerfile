############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder


WORKDIR /Users/eddie/Workspace/designing-distributed-systems/sidecar/logger/
COPY main.go .
# Fetch dependencies.
# Using go get.
# RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/logger
# RUN go build -o /go/bin/logger

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/logger /go/bin/logger
EXPOSE 8090
# Run the hello binary.
ENTRYPOINT ["/go/bin/logger"]