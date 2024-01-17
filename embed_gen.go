// Code generated by generator, DO NOT EDIT.
package main

import _ "embed"

//go:embed powershell-scripts/cpu-load-percentage.ps1
var cpu_load_percentage string

//go:embed powershell-scripts/enable-rdp.ps1
var enable_rdp string

//go:embed powershell-scripts/get-redirected-disks.ps1
var get_redirected_disks string

//go:embed powershell-scripts/get-service-start-times.ps1
var get_service_start_times string

//go:embed powershell-scripts/install-core-gui.ps1
var install_core_gui string

//go:embed powershell-scripts/rancherwins-logs.ps1
var rancherwins_logs string

//go:embed powershell-scripts/rke2-logs.ps1
var rke2_logs string

//go:embed powershell-scripts/scm-logs.ps1
var scm_logs string

//go:embed powershell-scripts/set-rdp-port.ps1
var set_rdp_port string

//go:embed powershell-scripts/show-hotfixes.ps1
var show_hotfixes string

//go:embed powershell-scripts/show-rdp-port.ps1
var show_rdp_port string

//go:embed powershell-scripts/watch-rke2-service.ps1
var watch_rke2_service string

//go:embed powershell-scripts/watch.ps1
var watch string

type cmd struct {
	filename string
	name     string
	content  string
}

var cmds = []cmd{{
	content:  cpu_load_percentage,
	filename: "cpu-load-percentage.ps1",
	name:     "cpu-load-percentage",
}, {
	content:  enable_rdp,
	filename: "enable-rdp.ps1",
	name:     "enable-rdp",
}, {
	content:  get_redirected_disks,
	filename: "get-redirected-disks.ps1",
	name:     "get-redirected-disks",
}, {
	content:  get_service_start_times,
	filename: "get-service-start-times.ps1",
	name:     "get-service-start-times",
}, {
	content:  install_core_gui,
	filename: "install-core-gui.ps1",
	name:     "install-core-gui",
}, {
	content:  rancherwins_logs,
	filename: "rancherwins-logs.ps1",
	name:     "rancherwins-logs",
}, {
	content:  rke2_logs,
	filename: "rke2-logs.ps1",
	name:     "rke2-logs",
}, {
	content:  scm_logs,
	filename: "scm-logs.ps1",
	name:     "scm-logs",
}, {
	content:  set_rdp_port,
	filename: "set-rdp-port.ps1",
	name:     "set-rdp-port",
}, {
	content:  show_hotfixes,
	filename: "show-hotfixes.ps1",
	name:     "show-hotfixes",
}, {
	content:  show_rdp_port,
	filename: "show-rdp-port.ps1",
	name:     "show-rdp-port",
}, {
	content:  watch_rke2_service,
	filename: "watch-rke2-service.ps1",
	name:     "watch-rke2-service",
}, {
	content:  watch,
	filename: "watch.ps1",
	name:     "watch",
}}