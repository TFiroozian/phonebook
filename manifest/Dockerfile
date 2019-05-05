FROM golang:1.10-alpine3.8 AS build    
    
RUN apk add --no-cache git    
RUN go get github.com/golang/dep/cmd/dep    

COPY Gopkg.lock Gopkg.toml $GOPATH/src/github.com/tfiroozian/phonebook/
WORKDIR $GOPATH/src/github.com/tfiroozian/phonebook   
RUN dep ensure -vendor-only    
    
COPY go/ $GOPATH/src/github.com/tfiroozian/phonebook/go
RUN go test -v ./... -tags=testing
RUN CGO_ENABLED=0 GOOS=linux go build -o $GOPATH/bin/sx1 -a -ldflags '-extldflags "-static"' ./go                                                                 
FROM scratch 
ENV GOPATH /go 

COPY --from=build $GOPATH/bin $GOPATH/bin
ENTRYPOINT ["/go/bin/phonebook"]
