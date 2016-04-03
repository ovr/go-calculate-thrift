
FROM golang:1.6

MAINTAINER Patsura Dmitry <talk@dmtry.me>

# Display version information
RUN go version
RUN go env

EXPOSE 9090

CMD cd src/server && go build .

ENTRYPOINT src/server/server
