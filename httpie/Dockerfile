FROM alpine:edge
LABEL maintainer="Ben Burke <glvnst@users.noreply.github.com>"

RUN apk add --no-cache \
    python3 \
    py3-pip \
  && pip3 --no-cache-dir install \
    httpie \
  && true
  #&& find / -name __pycache__ -type d -print0 | xargs -0 rm -r --

WORKDIR /work

ENTRYPOINT ["/usr/bin/http"]
CMD ["--help"]
