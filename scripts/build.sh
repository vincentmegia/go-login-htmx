pack build \
  --builder paketobuildpacks/builder-jammy-buildpackless-static \
  --buildpack paketo-buildpacks/go \
  --env "CGO_ENABLED=0" \
  --env "BP_GO_BUILD_FLAGS=-buildmode=default" \
  --env "BP_GO_BUILD_LDFLAGS=-X main.version=0.0.2" \
  go-login-htmx_0.0.3
