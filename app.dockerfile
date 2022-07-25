FROM registry.gitlab.com/rubin-dev/api/base:db58fad AS base
ARG CI_COMMIT_SHORT_SHA=local
ENV CI_COMMIT_SHORT_SHA=${CI_COMMIT_SHORT_SHA}
ENV TZ=Europe/Moscow
ENV DEBIAN_FRONTEND=noninteractive
WORKDIR /app
COPY . .
RUN make build-app

FROM gitlab.com/rubin-dev/dependency_proxy/containers/debian:bullseye-slim
RUN apt-get update && apt-get install -yq ca-certificates
ARG CI_COMMIT_SHORT_SHA=local
ENV CI_COMMIT_SHORT_SHA=${CI_COMMIT_SHORT_SHA}
WORKDIR /app/
ADD ./root.crt /app/
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    "appuser"
USER appuser:appuser
COPY --from=base /app/bin/app .
CMD ["/app/app"]
