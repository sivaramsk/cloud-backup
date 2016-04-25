# Pull base image
FROM debian:jessie

# Update packages
RUN apt-get update && \
  apt-get dist-upgrade -y && \
  apt-get install -y golang curl git && \
  mkdir -p /opt/build && \
  cd /opt/build && \
  git clone https://github.com/restic/restic.git && \
  cd /opt/build/restic && \
  go run build.go && \
  cp restic /usr/local/bin/ && \
  rm -rf /opt/build/restic && \
  apt-get clean -y && \
  apt-get autoclean -y && \
  apt-get autoremove -y && \
  rm -rf /usr/share/locale/* && \
  rm -rf /var/cache/debconf/*-old && \
  rm -rf /var/lib/apt/lists/* && \
  rm -rf /usr/share/doc/*
