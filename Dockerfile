FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch

# Statically linked go binary requires CA certs for
# SSL HTTP connections, import from latest alpine
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# requires statically linked go binary
COPY deamon /deamon

ENTRYPOINT ["/deamon"]