FROM golang:1.15.6
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /go/src/revel-prac
COPY . .
RUN go get -u github.com/revel/cmd/revel && \
  mkdir /build && \
  revel package --application-path= . --target-path=/build/app.tar.gz


FROM alpine
COPY --from=0 /build/app.tar.gz .
RUN tar xzvf app.tar.gz
ENTRYPOINT ["sh", "run.sh"]
