FROM go:1.16 AS build
WORKDIR /go/src/to-do
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main . 

FROM alpine:latest
COPY --from=build /go/src/to-do/main .
COPY --from=build /go/src/to-do/.env .
ENV TZ=Asia/Taipei
ENTRYPOINT ["./main"]