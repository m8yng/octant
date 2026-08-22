# OctantR desktop

Experimental desktop shell using Wails v3 beta. It reuses OctantR's existing loopback HTTP/WebSocket server; browser mode remains unchanged.

## Requirements

- Go 1.25 or newer
- pnpm 11
- Wails v3 native dependencies

Ubuntu 24.04:

```sh
sudo apt install build-essential pkg-config libgtk-4-dev libwebkitgtk-6.0-dev
```

## Run

From the repository root:

```sh
make run-desktop
```

Build a desktop binary:

```sh
make desktop-build
```

The Wails dependency is pinned in `desktop/go.mod` because v3 is currently beta.

## Ubuntu sandbox troubleshooting

Ubuntu may block WebKitGTK's Bubblewrap user namespace and report `bwrap: setting up uid map: Permission denied`. The `make run-desktop` development target disables the WebKit sandbox to work around this locally.

Configure an AppArmor policy that permits user namespaces before distributing or running the desktop build in production. `make desktop-build` does not embed or configure the unsafe environment variable.
