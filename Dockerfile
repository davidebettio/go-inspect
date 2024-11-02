FROM golang AS builder

COPY . /go-inspect
WORKDIR /go-inspect
ENV GO111MODULE=on
# do this in a separate layer to cache deps from build to build
RUN go get
RUN CGO_ENABLED=0 GOOOS=linux go build -o go-inspect

FROM alpine
WORKDIR /root/
COPY --from=builder /go-inspect .
CMD ["./go-inspect"]
