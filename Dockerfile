FROM golang:alpine as build
RUN apk add git make
RUN mkdir pg-ping
WORKDIR /opt/pg-ping
ADD . .
RUN make compile

FROM alpine
COPY --from=build /opt/pg-ping/pg-ping /bin/pg-ping
ENTRYPOINT [ "pg-ping" ]