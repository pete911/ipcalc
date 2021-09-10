FROM golang:1.17.1-alpine AS build
RUN apk add --no-cache gcc libc-dev

WORKDIR /go/src/app
COPY . .
RUN go test ./...
ARG version=dev
RUN go build -ldflags "-X main.Version=$version" -o /bin/ipcalc

FROM alpine:3.14.2

COPY --from=build /bin/ipcalc /usr/local/bin/ipcalc
ENTRYPOINT ["ipcalc"]
