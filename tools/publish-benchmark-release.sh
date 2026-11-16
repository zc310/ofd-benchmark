#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: publish-benchmark-release.sh <asset-directory>" >&2
  exit 2
fi
: "${GITHUB_REPOSITORY:?GITHUB_REPOSITORY is required}"
: "${RELEASE_TAG:?RELEASE_TAG is required}"
asset_dir=$(realpath "$1")
test -f "$asset_dir/RELEASE_NOTES.md"
test -f "$asset_dir/manifest.json"
(
  cd "$asset_dir"
  sha256sum --check SHA256SUMS
)
shopt -s nullglob
archives=("$asset_dir"/*.zip)
if [[ ${#archives[@]} -eq 0 ]]; then
  echo "No library ZIPs to publish." >&2
  exit 1
fi
if gh release view "$RELEASE_TAG" --repo "$GITHUB_REPOSITORY" >/dev/null 2>&1; then
  echo "Release $RELEASE_TAG already exists; refusing to overwrite it (including drafts)." >&2
  exit 1
fi

# Upload everything to a new draft first. On upload failure it stays a draft;
# remove that draft manually before retrying. Never modify an existing release.
assets=("$asset_dir"/*)
gh release create "$RELEASE_TAG" "${assets[@]}" \
  --repo "$GITHUB_REPOSITORY" --verify-tag --draft \
  --title "$RELEASE_TAG" --notes-file "$asset_dir/RELEASE_NOTES.md"
gh release edit "$RELEASE_TAG" --repo "$GITHUB_REPOSITORY" --draft=false
