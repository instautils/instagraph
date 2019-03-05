FROM golang:1.10

RUN mkdir -p $GOPATH/src/github.com/ahmdrz/instagraph
COPY . $GOPATH/src/github.com/ahmdrz/instagraph
WORKDIR $GOPATH/src/github.com/ahmdrz/instagraph
RUN go build -i -o instagraph

CMD [ "./instagraph" ]