FROM golang:alpine
EXPOSE 80

RUN apk add -U --no-cache curl && \
 curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/abihf/entitle/
ADD . .

RUN dep ensure
RUN go install .

CMD entitle
