FROM golang:1.12-alpine as builder
ENV SOURCES /go/src/cloud-native-go
COPY . ${SOURCES}
RUN go get -d -v ./...
RUN cd ${SOURCES} && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/cloud-native-go .

FROM alpine:3.9
#COPY ./bin/cloud-native-go /app/cloud-native-go
COPY --from=builder /go/bin/cloud-native-go app/cloud-native-go
RUN chmod +x /app/cloud-native-go
ENV PORT 8080
EXPOSE ${PORT}

ENTRYPOINT ["/app/cloud-native-go"]