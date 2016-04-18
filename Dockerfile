FROM golang:1.5

WORKDIR /go/src/github.com/o0khoiclub0o/piflab-store-api-go

ADD . ./
RUN go get github.com/tools/godep
RUN godep restore

ENV PORT 80
EXPOSE 80

RUN go install
CMD piflab-store-api-go

# For development
RUN go get bitbucket.org/zealotnt/goose/cmd/goose
RUN go get github.com/codegangsta/gin
RUN go install github.com/onsi/ginkgo/ginkgo
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/mattn/goveralls
RUN chmod +x ./testcoverage.sh

