FROM golang:alpine as build
LABEL authors="taosu"
ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    GOARCH=amd64
WORKDIR /go/src/memorandum
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o memorandum .
FROM scratch as deploy
COPY --from=build /go/src/memorandum/memorandum .
COPY --from=build /go/src/memorandum/config.yaml .
CMD ["./memorandum"]
