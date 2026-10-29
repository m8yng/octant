/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package component

import "github.com/vmware-tanzu/octant/internal/util/json"

var defaultDurations = []Since{
	{Label: "5 minutes", Seconds: 300},
	{Label: "10 minutes", Seconds: 600},
	{Label: "30 minutes", Seconds: 1800},
	{Label: "1 hour", Seconds: 3600},
	{Label: "3 hours", Seconds: 10800},
	{Label: "5 hours", Seconds: 18000},
	{Label: "Creation", Seconds: -1},
}

type Since struct {
	Label   string `json:"label,omitempty"`
	Seconds int64  `json:"seconds,omitempty"`
}

type LogsConfig struct {
	Namespace  string   `json:"namespace,omitempty"`
	Name       string   `json:"name,omitempty"`
	Containers []string `json:"containers,omitempty"`
	Durations  []Since  `json:"durations,omitempty"`
}

// Logs is a logs component.
//
// +octant:component
type Logs struct {
	Base
	Config LogsConfig `json:"config,omitempty"`
}

func NewLogs(namespace, name string, containers ...string) *Logs {
	return &Logs{
		Config: LogsConfig{
			Namespace:  namespace,
			Name:       name,
			Containers: containers,
			Durations:  defaultDurations,
		},
		Base: newBase(TypeLogs, TitleFromString("Logs")),
	}
}

// GetMetadata accesses the components metadata. Implements Component.
func (l *Logs) GetMetadata() Metadata {
	return l.Metadata
}

type logsMarshal Logs

func (l *Logs) MarshalJSON() ([]byte, error) {
	m := logsMarshal(*l)
	m.Metadata.Type = TypeLogs

	return json.Marshal(&m)
}
