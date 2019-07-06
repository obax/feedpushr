package model

import "net/url"

// PluginType is the plugin type qualifier
type PluginType int

const (
	// OUTPUT_PLUGIN output plugin type
	OUTPUT_PLUGIN PluginType = iota
	// FILTER_PLUGIN filter plugin type
	FILTER_PLUGIN
)

// PluginInfo contains plugins informations
type PluginInfo struct {
	Name string
	Type PluginType
}

// OutputPlugin is the interface of an output plugin
type OutputPlugin interface {
	// Build an output plugin
	Build(props url.Values, tags []string) (OutputProvider, error)
}

// FilterPlugin is the interface of an filter plugin
type FilterPlugin interface {
	// Build a filter
	Build(props FilterProps, tags []string) (Filter, error)
}
