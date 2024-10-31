FROM ghcr.io/skip-mev/connect-dev-base AS builder

WORKDIR /src/connect

# Set environment variables for repository, tag, and URL
ENV REPO_URL=https://github.com/skip-mev/connect
# Accept a build argument for the tag and set it as an environment variable
ARG TAG
ENV TAG=${TAG}

# Download the repository at the specific tag using curl && extract the downloaded tarball
RUN curl -L ${REPO_URL}/archive/refs/tags/v${TAG}.tar.gz -o /tmp/connect-v${TAG}.tar.gz \
    && tar -xzf /tmp/connect-v${TAG}.tar.gz -C /tmp \
    && mv -T /tmp/connect-${TAG} /src/connect \
    && rm -r /tmp/*

RUN go mod download

FROM builder AS connect-builder
RUN make build

FROM ubuntu:rolling AS connect
EXPOSE 8080 8002

COPY --from=connect-builder /src/connect/build/* /usr/local/bin/
COPY ./scripts/docker-entrypoint.sh /docker-entrypoint.sh

RUN apt-get update \
    && apt-get install jq -y \
    && apt-get install ca-certificates -y \
    && apt-get install curl -y

WORKDIR /usr/local/bin/
ENTRYPOINT ["/docker-entrypoint.sh"]

FROM builder AS blockchain

RUN make build-test-app

## Prepare the final clear binary
## This will expose the tendermint and cosmos ports alongside 
## starting up the sim app and the connect daemon
EXPOSE 26656 26657 1317 9090 7171 26655 8081 26660
RUN apt-get update && apt-get install jq -y && apt-get install ca-certificates -y
ENTRYPOINT ["make", "build-and-start-app"]

