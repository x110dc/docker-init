#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail

(
cat <<'EOF'
APP=appname
NAMESPACE=namespace

build:
	docker build --tag=${NAMESPACE}/${APP} .
#
#  You could use a completely different image to debug; maybe one with your
#  favorite tools already present:
#
debug:
	docker run --volumes-from=${APP} --interactive=true --tty=true ${NAMESPACE}/${APP} bash
#
#  The publish here will depend on your environment
#
run:
	docker run --name=${APP} --detach=true --publish=80:8000 ${NAMESPACE}/${APP}
clean:
	docker stop ${APP} && docker rm ${APP}

EOF
) > Makefile

(
cat <<'EOF'
#  decent chance you will not need .git as part of the context; for a big repo
#  this will speed things up:
.git
EOF
) > .dockerignore

(
cat <<'EOF'
FROM ubuntu:trusty
MAINTAINER foo@bar.org
EOF
) > Dockerfile
