############################
# STEP 1 build executable binary
############################
FROM golang:1.17.3-alpine3.13 as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR $GOPATH/src/practice
COPY . .

RUN go get .
RUN go mod verify
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/practice

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable
COPY --from=builder /go/bin/practice /go/bin/practice

ENTRYPOINT ["/go/bin/practice"]