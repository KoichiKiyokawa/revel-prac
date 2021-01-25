FROM golang:1.15.6
ENV GO111MODULE=on
WORKDIR /go/src/revel-prac
COPY . .
RUN go get -u github.com/revel/cmd/revel && revel build . /dist


FROM alpine
COPY --from=0 /dist/run.sh .
ENTRYPOINT ["sh", "run.sh"]
