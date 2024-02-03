# This is a microservice that is intended to be used in a sidecar model
FROM golang:1.21.6-bullseye

RUN apt-get update -y && apt-get upgrade -y

WORKDIR /opt/certlint/

# Copy source code over
COPY certlint/go.mod /opt/certlint/
COPY certlint/go.sum /opt/certlint/
COPY certlint/main.go /opt/certlint/

ENV GIN_MODE=release

# Build 
RUN go build

CMD [ "/opt/certlint/certlint" ]
