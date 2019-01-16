FROM alpine

MAINTAINER Dmitry Mozzherin

ENV LAST_FULL_REBUILD 2019-01-16

WORKDIR /bin

COPY ./gnparser/gnparser /bin

CMD ["gnparser", "-g", "8778"]


