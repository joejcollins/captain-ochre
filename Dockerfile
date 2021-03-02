FROM golang:1.14

WORKDIR /captain-ochre
COPY . .

RUN go get -d -v ./src
RUN go install -v ./src

CMD ["app"]