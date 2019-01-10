FROM golang AS builder

WORKDIR /home/app
COPY go.mod go.sum ./
RUN echo "download mod" \
    && go mod download

COPY . .
RUN echo "build" \
    && make build


FROM golang
WORKDIR /home

# migration
COPY --from=builder /go/bin/greenwall /usr/local/bin/
COPY --from=builder /home/app/config.yaml ./config.yaml
COPY --from=builder /home/app/frontend ./frontend

CMD [ "greenwall" ]