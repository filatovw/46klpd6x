FROM golang:1.15.7 as builder

COPY . /usr/app/
WORKDIR /usr/app/
RUN make api/build

FROM golang:1.15.7

ARG REVISION
ARG BUILDTIME
ARG VERSION
ARG APP
LABEL REVISION=${REVISION}
LABEL BUILDTIME=${BUILDTIME}
LABEL VERSION=${VERSION}

RUN groupadd executor \
    && useradd -u 1000 -g executor -d /usr/app executor

ENV APP_USER=executor APP_GROUP=executor
ENV ROOT_USER=root ROOT_GROUP=root
ENV APP=${APP}

WORKDIR /usr/bin/
COPY --from=builder /usr/app/bin/${APP} .
COPY ./deployment/${APP}/entrypoint.sh /entrypoint.sh
USER ${ROOT_USER}
RUN chmod +x /entrypoint.sh
USER ${APP_USER}
ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
