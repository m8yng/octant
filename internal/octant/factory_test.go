// Copyright (c) 2026 m8yng and OctantR contributors.
// SPDX-License-Identifier: MIT

package octant

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware-tanzu/octant/pkg/navigation"
	"github.com/vmware-tanzu/octant/pkg/store"
)

func TestNavigationFactoryAllNamespaces(t *testing.T) {
	var queryNamespace string
	factory := NewNavigationFactory(AllNamespaces, "overview", nil, NavigationEntries{
		Lookup: map[string]string{"Pods": "workloads/pods"},
		EntriesFuncs: map[string]EntriesFunc{
			"Pods": func(_ context.Context, _ string, namespace string, _ store.Store, _ bool) ([]navigation.Navigation, bool, error) {
				queryNamespace = namespace
				return nil, false, nil
			},
		},
		Order: []string{"Pods"},
	})

	entries, err := factory.Generate(context.Background(), "", false)
	require.NoError(t, err)
	require.Equal(t, "overview/all-namespaces/", factory.Root())
	require.Equal(t, "", queryNamespace)
	require.Equal(t, "overview/all-namespaces/workloads/pods", entries[0].Path)
}
