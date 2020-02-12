FROM golang:alpine3.10

ARG USER=appuser
ARG GROUP=appgroup

RUN apk --no-cache add ca-certificates && \
  addgroup -S ${GROUP} && adduser -S ${USER} -G ${GROUP}

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe

WORKDIR /home/${USER}

COPY --chown=${USER}:${GROUP} app/ .

USER ${USER}

RUN go build -o server

CMD ["./server"]

