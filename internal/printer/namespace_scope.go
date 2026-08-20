// Copyright (c) 2026 m8yng and OctantR contributors.
// SPDX-License-Identifier: MIT

package printer

import "context"

type allNamespacesKey struct{}

func WithAllNamespaces(ctx context.Context) context.Context {
	return context.WithValue(ctx, allNamespacesKey{}, true)
}

func IsAllNamespaces(ctx context.Context) bool {
	all, _ := ctx.Value(allNamespacesKey{}).(bool)
	return all
}
