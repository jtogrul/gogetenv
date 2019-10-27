FROM golang:latest AS builder
ADD . /source
WORKDIR /source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /gogetenv .

FROM scratch
COPY --from=builder /gogetenv ./
EXPOSE 80
CMD ["/gogetenv"]