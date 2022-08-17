FROM golang:1.16 as builder
WORKDIR /apps/
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build app.go

FROM scratch
COPY --from=builder /apps/app /

EXPOSE 8080
CMD ["/app"]