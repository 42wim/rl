podman run --rm -ti -v $PWD:/build -e GITHUB_TOKEN -w /build goreleaser/goreleaser release --rm-dist
