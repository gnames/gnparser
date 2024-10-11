FROM alpine:3.14

LABEL maintainer="Dmitry Mozzherin"

ENV LAST_FULL_REBUILD=2024-10-11

WORKDIR /bin

COPY ./gnparser/gnparser /bin

ENTRYPOINT [ "gnparser" ]

CMD ["-p", "8778"]
