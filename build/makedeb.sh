#!/bin/bash
set -ex

if [[ $1 ]]; then
  TAG=$1 # tag provided in command line overrides environment variable
fi

if [[ -z $TAG ]]; then
  echo "Debian code name isn't specified! Expected jessie, stretch, buster or bullseye."
  exit 1
fi

# upstream-version
UPSTREAM=$(git describe --tags --always --dirty  --match="v*")
UPSTREAM=${UPSTREAM#v} # v1.0.0 -> 1.0.0
BUILD_TIME="$(date +%FT%T%z)"
if [[ -z $BUILD_VERSION ]]; then
  if [[ -z $BUILD_VERSION_SUFFIX ]]; then
    # epoch:upstream-version-debian.revision
    BUILD_VERSION="1:$UPSTREAM-$TAG"
  else
    BUILD_VERSION="1:$UPSTREAM-$BUILD_VERSION_SUFFIX"
  fi
fi

if [[ -z $GID ]]; then
  GID=$(id -g)
fi

# build tlgen
if [[ -z $GOLANG_IMAGE ]]; then
  GOLANG_IMAGE="golang:1.20-$TAG" # e.g. golang:1.19-bullseye
fi
GOCACHE=build/go-cache
mkdir -p "$PWD/$GOCACHE"
docker run --rm -u "$UID:$GID" -v "$PWD:/src" -w /src \
  -e BUILD_MACHINE="$(uname -n -m -r -s)" -e BUILD_TIME="$BUILD_TIME" -e BUILD_VERSION="$UPSTREAM" \
  -e BUILD_COMMIT="$(git log --format="%H" -n 1)" -e BUILD_COMMIT_TS="$(git log --format="%ct" -n 1)" \
  -e GOCACHE="/src/$GOCACHE" \
  "$GOLANG_IMAGE" make build

# build debian package
(cd build
rm -f debian/changelog
DEB_IMAGE="tlgen-build-deb"
docker build -t "$DEB_IMAGE" - < debuild.Dockerfile
docker run --rm -v "$PWD:/src" -w /src -u "$UID:$GID" "$DEB_IMAGE" dch \
  --create --distribution stable --package tlgen \
  --newversion "$BUILD_VERSION" "up to version $BUILD_VERSION"
docker run --rm -v "$PWD/..:/src" -w /src/build -u "$UID:$GID" "$DEB_IMAGE" debuild --no-lintian -us -uc -b)

# drop to target directory
mkdir -p target/deb
for f in *"${BUILD_VERSION##[0-9]*\:}"*; do mv -u "$f" "target/deb/$f"; done
