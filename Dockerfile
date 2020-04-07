FROM alpine:latest as certs
RUN apk --update add ca-certificates

## Grab the certs
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=certs /tmp /tmp

ADD build/_output/bin/mal /usr/local/bin/mal
ENTRYPOINT ["mal"]