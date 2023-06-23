FROM golang:1.20.5 AS builder

ARG VERSION
ARG INPUT_FILE
ARG OUTPUT_FILE
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV GOSUMDB=off

WORKDIR /go/src/build
#RUN go env
COPY . .

RUN go mod download
RUN CGO_ENABLED=0  go build -ldflags "-X main.Version=${VERSION}" -o ${OUTPUT_FILE} ${INPUT_FILE}


FROM alpine:latest
ARG OUTPUT_FILE
WORKDIR /app
COPY --from=builder /go/src/build/${OUTPUT_FILE} ./
CMD ["./app", "-conf", "./configs"]

