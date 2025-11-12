FROM golang:1.24.4
ADD . /code
WORKDIR /code
RUN go mod download
RUN go build -o blue-website main.go
CMD ["./blue-website"]