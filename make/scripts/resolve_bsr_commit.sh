#!/usr/bin/env bash

set -eo pipefail

# resolve_bsr_commit resolves a git ref (semver tag, branch, or commit
# hash) to a BSR module commit name.
#
# Usage: resolve_bsr_commit.sh <module> <ref>
#
# Examples:
#   resolve_bsr_commit.sh buf.build/bufbuild/protovalidate v1.1.0
#   resolve_bsr_commit.sh buf.build/bufbuild/protovalidate main
#   resolve_bsr_commit.sh buf.build/bufbuild/protovalidate 895eefca6d13

MODULE="${1}"
REF="${2}"

# If the ref is a valid semver, return it directly.
if echo "${REF}" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?$'; then
  echo "${REF}"
  exit 0
fi

# Try to resolve the ref directly as a BSR label.
if buf registry module commit resolve "${MODULE}:${REF}" --format json >/dev/null 2>&1; then
  echo "${REF}"
  exit 0
fi

# Fall back to searching commits by source_control_url suffix.
PAGE_TOKEN=""
while true; do
  PAGE_ARGS=()
  if [[ -n "${PAGE_TOKEN}" ]]; then
    PAGE_ARGS=(--page-token="${PAGE_TOKEN}")
  fi

  RESULT="$(buf registry module commit list "${MODULE}" --format json --page-size 100 "${PAGE_ARGS[@]+"${PAGE_ARGS[@]}"}")"

  FOUND="$(echo "${RESULT}" | jq -r --arg ref "${REF}" \
    '.commits[] | select(.source_control_url != null) | select(.source_control_url | endswith($ref)) | .commit')"

  if [[ -n "${FOUND}" ]]; then
    echo "${FOUND}"
    exit 0
  fi

  PAGE_TOKEN="$(echo "${RESULT}" | jq -r '.next_page // empty')"
  if [[ -z "${PAGE_TOKEN}" ]]; then
    echo "error: could not resolve ref '${REF}' for module '${MODULE}'" >&2
    exit 1
  fi
done
