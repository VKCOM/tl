#!/usr/bin/env bash
set -euo pipefail

TAG="${1:-}"

# 1) взять самую актуальную версию master
git fetch origin master
git checkout master
git pull --ff-only origin master

# 2) если тег не передан — сгенерировать следующий и добавить -rc
if [[ -z "${TAG}" ]]; then
  # Берём последний тег (по semver). Если тегов нет — стартуем с v0.0.0
  LAST_TAG="$(git tag -l --sort=-v:refname | head -n 1 || true)"
  if [[ -z "${LAST_TAG}" ]]; then
    LAST_TAG="v0.0.0"
  fi

  # Парсим vMAJOR.MINOR.PATCH + любой суффикс (например: v1.2.3, v1.2.3-rc, v1.2.3-rc1, v1.2.3-anything)
  if [[ ! "${LAST_TAG}" =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)([-+].*)?$ ]]; then
    echo "Не могу распарсить последний тег: '${LAST_TAG}'. Ожидаю формат vMAJOR.MINOR.PATCH{suffix}"
    echo "Примеры: v1.2.3, v1.2.3-rc, v1.2.3-rc1"
    exit 1
  fi

  MAJOR="${BASH_REMATCH[1]}"
  MINOR="${BASH_REMATCH[2]}"
  PATCH="${BASH_REMATCH[3]}"

  PATCH=$((PATCH + 1))
  TAG="v${MAJOR}.${MINOR}.${PATCH}"
fi

go mod tidy

git tag "${TAG}"
git push origin "${TAG}"

echo "Created and pushed tag: ${TAG}"