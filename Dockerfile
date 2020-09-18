FROM golang

MAINTAINER 龙泉

WORKDIR /data/go/src/runGo

COPY . .

RUN go build main.go

EXPOSE 9990

CMD ["./main"]