# Octant-Revibe

OctantR is an independent community revival of the archived [Octant](https://github.com/vmware-archive/octant) Kubernetes dashboard.

This project is not affiliated with or endorsed by VMware or Broadcom. The upstream name is used only to identify the project's origin.

## Status

OctantR is under active revival. It is based on Octant 0.25.1, and there is not yet a supported OctantR release. Build from source and expect compatibility work with current Kubernetes versions.

## Features

- Visualize relationships between Kubernetes resources.
- Inspect resource configuration and status.
- Stream container logs.
- Forward local ports to pods.
- Filter resources by labels.
- Switch between namespaces and kubeconfig contexts.
- Extend views through plugins.

## Build and run

Requirements:

- Go 1.17 or newer (Go 1.25 for the desktop app)
- Node.js 20 or newer
- pnpm 11
- Access to a Kubernetes cluster through a kubeconfig

```sh
git clone https://github.com/m8yng/octant.git
cd octant
make run
```

Open <http://127.0.0.1:7777> if the browser does not open automatically.

Available targets:

```text
make help           Show available targets
make build          Build the browser binary
make run            Build and run in the browser
make run-desktop    Build and run the Wails desktop app
make desktop-build  Build the Wails desktop binary
```

The desktop app uses [Wails v3](https://github.com/wailsapp/wails) and reuses the same local HTTP/WebSocket backend as browser mode. Desktop prerequisites are documented in [`desktop/README.md`](desktop/README.md).

Verify cluster access before starting:

```sh
kubectl cluster-info
```

For frontend development with automatic rebuilding:

```sh
go run build.go serve
```

The frontend runs on <http://localhost:4200> and the backend on <http://localhost:7777>.

See [HACKING.md](HACKING.md) for additional development commands.

## Plugins

The sample plugin is in [`cmd/octant-sample-plugin`](cmd/octant-sample-plugin). Install it to the default plugin directory with:

```sh
go run build.go install-test-plugin
```

The plugin API is inherited from upstream Octant and may change as the revival progresses.

## Contributing

Use [GitHub issues](https://github.com/m8yng/octant/issues) for bugs and proposals. See [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.

## License and attribution

Original Octant code remains under the [Apache License 2.0](LICENSE-APACHE-2.0). Original OctantR contributions beginning with commit `8024ecda` are under the [MIT License](LICENSE). Both licenses apply to the combined codebase; see [NOTICE](NOTICE).
