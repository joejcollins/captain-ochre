FROM golang:1.15-alpine

WORKDIR /captain-ochre
COPY . .

ENV GOPATH="/captain-ochre"
RUN go get -d -v ./src/captain-ochre
RUN go install -v ./src/captain-ochre

CMD ["app"]