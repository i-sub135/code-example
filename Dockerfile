FROM golang:latest
MAINTAINER iyan Subdiana '<subdiana.project@gmail.id>'

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
COPY . .

RUN go mod download
RUN make build
CMD ./main

EXPOSE 6100
