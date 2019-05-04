FROM golang:1.12-alpine
ADD . /go/src/github.com/azbyluthfan/go-hr/
WORKDIR /go/src/github.com/azbyluthfan/go-hr
RUN go mod download
RUN go build -o "/go/src/github.com/azbyluthfan/go-hr/go-hr"
EXPOSE 9000
CMD ["/go/src/github.com/azbyluthfan/go-hr/go-hr"]