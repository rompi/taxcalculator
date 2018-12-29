FROM golang:1.10-alpine as builder

RUN apk add --no-cache git
RUN mkdir -p /code/src/github.com/rompi/tax-calc

ENV GOPATH /code
ENV PATH /code/bin:/usr/local/go/bin:$PATH

RUN go get -u github.com/golang/dep/cmd/dep

COPY app /code/src/github.com/rompi/tax-calc/app
COPY cmd /code/src/github.com/rompi/tax-calc/cmd
COPY config /code/src/github.com/rompi/tax-calc/config
COPY constant /code/src/github.com/rompi/tax-calc/constant

COPY Gopkg.toml /code/src/github.com/rompi/tax-calc
COPY Gopkg.lock /code/src/github.com/rompi/tax-calc

RUN cd /code/src/github.com/rompi/tax-calc && dep ensure -v

RUN cd /code/src/github.com/rompi/tax-calc && go test ./... -cover

RUN cd /code/src/github.com/rompi/tax-calc/cmd && env GOARCH=amd64  GOOS=linux go build -o taxcalc && mv taxcalc /code/bin

FROM alpine:latest

ENV TZ=Asia/Jakarta
RUN apk --no-cache add ca-certificates bash bash-completion tzdata

RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p /opt/bin && mkdir -p /opt/config && mkdir -p /var/log/tax-calc && touch /var/log/tax-calc/taxcalc.log
COPY config /opt/config
WORKDIR /opt/bin

COPY --from=builder /code/bin/taxcalc .

EXPOSE 3000

CMD ["/opt/bin/taxcalc"]

