package main

import (
	"example/events"

	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/MinimixMC/AuroraAPI/API/aurora/event"
)

// Define the plugin
var Plugin aurora.Plugin = aurora.Plugin{
	Name: "Aurora Example Plugin",
	Author: []string{
		"I'm Carsen",
	},
	Version: aurora.PluginVersion{
		Version: "v0.0.1",
		VersionSupport: map[int]bool{
			1: true,
		},
	},
	MainFunc: MainFunc,
}

// The starting point of the plugin
func MainFunc(p *aurora.Plugin) {
	// Register an event
	aurora.RegisterEvent(event.PingEvent, events.HandlePingEvent)
}
