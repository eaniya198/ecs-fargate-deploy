FROM public.ecr.aws/docker/library/golang:latest as builder
WORKDIR /apps/
COPY src/ .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go 

FROM scratch
COPY --from=builder /apps/main /
EXPOSE 5000
CMD ["/main"]