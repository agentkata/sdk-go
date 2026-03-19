#!/bin/sh

set -eu

if [ "$#" -ne 1 ]; then
  echo "usage: $0 /path/to/solver.yaml" >&2
  exit 1
fi

SOURCE_SPEC=$1
ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
cp "$SOURCE_SPEC" "$ROOT_DIR/openapi/solver.yaml"

SOURCE_REPO=$(git -C "$(dirname -- "$SOURCE_SPEC")" rev-parse --show-toplevel 2>/dev/null || true)
SOURCE_COMMIT=null
SOURCE_REPO_NAME=unknown
if [ -n "$SOURCE_REPO" ]; then
  SOURCE_COMMIT=$(git -C "$SOURCE_REPO" rev-parse HEAD 2>/dev/null || echo null)
  SOURCE_REPO_NAME=$(basename "$SOURCE_REPO")
fi

cat > "$ROOT_DIR/openapi/SOURCE.json" <<JSON
{
  "source_repo": "$SOURCE_REPO_NAME",
  "source_commit": "$SOURCE_COMMIT",
  "source_spec": "$SOURCE_SPEC",
  "synced_at": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
}
JSON
