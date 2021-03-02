FROM golang:1.15-alpine

WORKDIR /workspace/captain-ochre
COPY . .

RUN go install ./...
EXPOSE 8080

CMD ["api"]