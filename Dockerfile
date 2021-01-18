FROM alpine

LABEL maintainer="Dmitry Mozzherin"

ENV LAST_FULL_REBUILD 2019-01-16

WORKDIR /bin

COPY ./gnparser/gnparser /bin

ENTRYPOINT [ "gnparser" ]

CMD ["-p", "8778"]
