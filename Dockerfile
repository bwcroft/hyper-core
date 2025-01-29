FROM alpine:3.21.2 AS goose
RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash && \
    rm -rf /var/cache/apk/*
ADD https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose && /bin/goose -version
WORKDIR /app
ENTRYPOINT ["/bin/goose"]
