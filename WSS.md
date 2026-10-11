# Headscale WSS Android build

Based on upstream release-branch/1.104. The Go core is pinned through the go.mod
replace directive to the same WSS fork used on Windows/Linux/macOS. Core patches
are not duplicated here. Android VPN socket protection remains enabled.

The application ID is `com.mikewong.headscale.wss` and the visible app name is
Headscale WSS. It installs alongside the official app; only one VPN can be active.
It is an independent build, not an official Tailscale release.

Build in the upstream Docker environment, supplying endpoint configuration:

```sh
make docker-build-image
make docker-android-dir
docker run --rm \
  -v "$PWD:/build/tailscale-android" \
  -v "$PWD/.android-docker:/root/.android" \
  -v "$PWD/.gradle-docker:/build/.gradle" \
  -v "$PWD/.cache-docker:/build/.cache" \
  --env GOPATH=/build/.cache/go --env GOMODCACHE=/build/.cache/go/pkg/mod \
  tailscale-android-build-amd64-20260912-1 make tailscale-debug \
  'WSS_LDFLAGS=-X github.com/tailscale/tailscale-android/libtailscale.wssControlURL=https://original.example.com -X github.com/tailscale/tailscale-android/libtailscale.wssEndpoint=https://wss.example.com'
```

Use the original URL in the existing custom-control-server login screen. The
transport maps it to the WSS endpoint without changing the profile URL. A build
without both linker settings uses the normal transport.

The separately signed app cannot read the official app's encrypted state, or
upgrade its APK in place. Plan a new registration unless a supported identity
export/import mechanism is established. Preserve the official app and its data
for rollback. Keep the private signing key outside Git; debug APKs are for
testing. Subsequent releases must use the same signing key to preserve this
fork's own device state.
