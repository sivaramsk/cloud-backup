FROM alpine:3.3

RUN set -x \
	&& apk add --no-cache --virtual .build-deps bash go bzr git mercurial subversion openssh-client ca-certificates \
        && mkdir -p /opt/build \
	&& cd /opt/build \
	&& git clone https://github.com/restic/restic.git \
	&& cd /opt/build/restic \
	&& go run build.go \
	&& cp restic /usr/local/bin/ \
        && rm -rf /opt/build/restic \
        && runDeps="$( \
                scanelf --needed --nobanner --recursive /usr/local \
                        | awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' \
                        | sort -u \
                        | xargs -r apk info --installed \
                        | sort -u \
        )" \
        && apk add --virtual .haproxy-rundeps $runDeps \
        && apk del .build-deps

