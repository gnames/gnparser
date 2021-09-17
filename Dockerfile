FROM alpine:3.14

LABEL maintainer="Dmitry Mozzherin"

ENV LAST_FULL_REBUILD 2021-04-07

WORKDIR /bin

COPY ./gnparser/gnparser /bin

ENTRYPOINT [ "gnparser" ]

CMD ["-p", "8778"]
