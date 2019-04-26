FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/go-upload/
COPY . /go/src/github.com/go-upload/
RUN go build -o ./dist/go-upload

FROM alpine:3.5
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

WORKDIR /app
COPY ./config/go-upload.yaml /var/app/
COPY ./config/go-upload.yaml /
COPY ./config/fine-iterator-231706-04d01216bb7a.json /var/app/
COPY ./form.html .
COPY --from=builder go/src/github.com/go-upload/dist/go-upload .
ENV GOOGLE_APPLICATION_CREDENTIALS /var/app/fine-iterator-231706-04d01216bb7a.json
EXPOSE 9090
ENTRYPOINT ["./go-upload"]
