FROM golang:1.14

WORKDIR /usr/src/groupie-tracker

COPY ./ ./

RUN go build -o gt

EXPOSE 9000

ENTRYPOINT ["./gt"]