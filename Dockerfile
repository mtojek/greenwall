FROM golang AS builder

LABEL maintainer="llitfkitfk@gmail.com"

WORKDIR /home/app
COPY go.mod go.sum ./
RUN echo "download mod" \
    && go mod download
COPY . .
RUN echo "build app" \
    && make build


FROM golang
WORKDIR /home

COPY --from=builder /go/bin/greenwall /usr/local/bin/
COPY --from=builder /home/app/config.yaml ./config.yaml
COPY --from=builder /home/app/frontend ./frontend

CMD [ "greenwall" ]