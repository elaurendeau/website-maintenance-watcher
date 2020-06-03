FROM golang:1.14
WORKDIR /go
COPY /src .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application .

FROM alpine:latest

ENV USER_KEY=""
ENV TOKEN_KEY=""

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/application .
CMD ["sh", "-c", "./application ${TOKEN_KEY} ${USER_KEY}"]