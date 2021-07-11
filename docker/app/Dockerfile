FROM golang:1.16-alpine as build

ENV CGO_ENABLED=0
# Create workdir
WORKDIR /app

# Copy modules files
# COPY ./go.sum .
# COPY ./go.mod .

COPY . .
#Installation modules
RUN go mod download

#Copy last part project
# COPY ./cmd .
# COPY ./pkg .
# COPY ./ui .

# Build project bin
RUN go build -o /test ./cmd/web/
CMD sleep 10;
CMD top
# Deploy artifact to small image
# FROM debian:buster
# ENV CGO_ENABLED=0
# WORKDIR /

# COPY --from=build /test /test



