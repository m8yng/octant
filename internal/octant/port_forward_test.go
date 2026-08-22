// Copyright (c) 2026 m8yng and OctantR contributors.
// SPDX-License-Identifier: MIT

package octant

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware-tanzu/octant/pkg/action"
)

func TestPortForwardRequestKinds(t *testing.T) {
	payload := action.Payload{
		"apiVersion": "v1",
		"kind":       "Service",
		"name":       "dns",
		"namespace":  "kube-system",
		"port":       float64(53),
	}

	_, err := portForwardRequestFromPayload(payload)
	require.NoError(t, err)

	payload["kind"] = "Deployment"
	_, err = portForwardRequestFromPayload(payload)
	require.EqualError(t, err, "only supports forwards for v1 Pods and Services")
}
