DESKTOP_TAGS := embedded exclude_graphdriver_devicemapper exclude_graphdriver_btrfs containers_image_openpgp

.PHONY: help build run run-desktop desktop-build

help:
	@printf '%s\n' \
		'make build          Build Angular UI and Go browser binary' \
		'make run            Build and start browser interface' \
		'make run-desktop    Build UI and run Wails desktop app' \
		'make desktop-build  Build UI and Wails desktop binary'

build:
	NG_CLI_ANALYTICS=false go run build.go ci-quick

run: build
	./build/octant

run-desktop:
	NG_CLI_ANALYTICS=false go run build.go web-build
	cd desktop && WEBKIT_DISABLE_SANDBOX_THIS_IS_DANGEROUS=1 go run -tags '$(DESKTOP_TAGS)' .

desktop-build:
	NG_CLI_ANALYTICS=false go run build.go web-build
	cd desktop && go build -tags '$(DESKTOP_TAGS)' -o ../build/octantr-desktop .
