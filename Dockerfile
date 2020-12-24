FROM golang:1.15.6-alpine3.12
RUN mkdir -p /zeki-demo
WORKDIR /zeki-demo
COPY . .
RUN  go build
EXPOSE 8080
RUN chmod +x zeki-demo
ENV RUN_MODE=docker
ENTRYPOINT ./zeki-demo
