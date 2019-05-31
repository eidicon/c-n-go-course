FROM alpine:3.9

COPY ./bin/cloud-native-go /app/cloud-native-go
RUN chmod +x /app/cloud-native-go
ENV PORT 8080
EXPOSE ${PORT}

ENTRYPOINT ["/app/cloud-native-go"]