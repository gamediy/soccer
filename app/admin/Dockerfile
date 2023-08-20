FROM alpine:latest AS production
MAINTAINER whisper
WORKDIR /app
EXPOSE 3000
COPY . /app
RUN apk add -U tzdata
RUN apk add openssh-client
RUN apk add vim
RUN chmod +x ./admin
ENTRYPOINT ["./admin", "--gf.gcfg.file=config-dev.yaml"]