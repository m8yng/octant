// Copyright (c) 2026 m8yng and OctantR contributors.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"

	internalLog "github.com/vmware-tanzu/octant/internal/log"
	"github.com/vmware-tanzu/octant/pkg/config"
	"github.com/vmware-tanzu/octant/pkg/dash"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	logger, err := internalLog.Init(0)
	if err != nil {
		cancel()
		return fmt.Errorf("initialize logger: %w", err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		cancel()
		logger.Close()
		return fmt.Errorf("listen: %w", err)
	}

	options := []dash.RunnerOption{
		dash.WithKubeConfig(kubeConfigPath()),
		dash.WithListener(listener),
		dash.WithoutOpeningBrowser(),
		dash.WithClientQPS(200),
		dash.WithClientBurst(400),
		dash.WithClientUserAgent("octantr-desktop"),
		dash.WithBuildInfo(config.BuildInfo{Version: "desktop-dev"}),
	}
	runner, err := dash.NewRunner(ctx, logger, options...)
	if err != nil {
		cancel()
		listener.Close()
		logger.Close()
		return fmt.Errorf("create dashboard: %w", err)
	}

	started := make(chan bool, 1)
	stopped := make(chan bool, 1)
	go func() {
		if err := runner.Start(started, stopped, options...); err != nil {
			logger.Errorf("dashboard stopped: %v", err)
		}
	}()
	select {
	case <-started:
	case <-time.After(30 * time.Second):
		cancel()
		listener.Close()
		logger.Close()
		return fmt.Errorf("dashboard did not start within 30 seconds")
	}

	var shutdownOnce sync.Once
	shutdown := func() {
		shutdownOnce.Do(func() {
			cancel()
			<-stopped
			logger.Close()
		})
	}

	// ponytail: reuse the proven loopback server; embed the handler in Wails only if a second transport becomes necessary.
	app := application.New(application.Options{
		Name:        "OctantR",
		Description: "Kubernetes dashboard",
		Assets:      application.AlphaAssets,
		OnShutdown:  shutdown,
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:     "OctantR",
		URL:       "http://" + listener.Addr().String(),
		Width:     1440,
		Height:    900,
		MinWidth:  900,
		MinHeight: 600,
	})

	err = app.Run()
	shutdown()
	if err != nil {
		return fmt.Errorf("run desktop app: %w", err)
	}
	return nil
}

func kubeConfigPath() string {
	if path := os.Getenv("KUBECONFIG"); path != "" {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".kube", "config")
}
