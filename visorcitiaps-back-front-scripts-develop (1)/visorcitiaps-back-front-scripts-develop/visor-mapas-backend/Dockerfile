FROM golang:1.20.5 as builder

WORKDIR /go/src/github.com/citiaps/visorcitiaps-back-front-scripts
COPY . .
RUN go mod init
RUN go get -d -v

#RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix app.go -o visor-backend


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /go/src/gitub.com/citiaps/visorcitiaps-back-front-scripts/visor-mapas-backend

COPY --from=builder /go/src/github.com/citiaps/visorcitiaps-back-front-scripts/visor-mapas-backend .
COPY --from=builder /go/src/github.com/citiaps/visorcitiaps-back-front-scripts/config/config.yml ./config/config.yml

COPY --from=builder /go/src/github.com/citiaps/visorcitiaps-back-front-scripts/visor-mapas-backend .
EXPOSE 2000

ENTRYPOINT ["./visor-mapas-backend"]
