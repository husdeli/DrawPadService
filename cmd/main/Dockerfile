FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make && \
    rm -rf /var/cache/apk/*

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN make build

##
## Deploy
##
FROM alpine:latest
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/main .
COPY --from=build /app/cmd/main/entrypoint.sh .
COPY --from=build /app/.env .
RUN ls -la
ENTRYPOINT ["./entrypoint.sh"]
