# FROM scratch
FROM alpine
ADD /go-logging //
ADD /config.json //
EXPOSE 8080
ENTRYPOINT ["/go-logging"] 