ARG TARGETPLATFORM="linux/amd64"

FROM golang:1.26

COPY msg_consumer /usr/bin/msg_consumer

CMD ["msg_consumer"]
