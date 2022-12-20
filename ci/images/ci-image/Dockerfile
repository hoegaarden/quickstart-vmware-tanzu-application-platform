ARG BASE=debian:stable-slim
FROM "${BASE}"

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -qy update \
        && apt-get -qy install -o APT::Install-Suggests=0 -o APT::Install-Recommends=0 \
            curl \
            jq \
            perl `# we need this for ytt's installer, which needs shasum` \
            python3-pip \
        && rm -rf /var/lib/apt/lists/*

RUN pip3 install --no-cache-dir \
        taskcat \
        yq

RUN dl="$(mktemp)" \
        && curl -Ls https://carvel.dev/install.sh > "$dl" \
        && bash "$dl" \
        && rm -f -- "$dl"
