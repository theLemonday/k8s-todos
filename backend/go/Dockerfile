# build
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /k8s-todos-backend .

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# run
FROM alpine:latest AS build-release-stage

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /k8s-todos-backend /k8s-todos-backend

# only needed if run the docker container alone, not when using k8s
# EXPOSE 8080

USER nonroot:nonroot

CMD [ "/k8s-todos-backend" ]