// Code generated by generator, DO NOT EDIT.
package main

import _ "embed"

//go:embed powershell-scripts/get-cpu-load.ps1
var get_cpu_load string

//go:embed powershell-scripts/get-redirected-disks.ps1
var get_redirected_disks string

//go:embed powershell-scripts/get-service-start-times.ps1
var get_service_start_times string

//go:embed powershell-scripts/getLogs-rancherwins.ps1
var getLogs_rancherwins string

//go:embed powershell-scripts/getLogs-rke2.ps1
var getLogs_rke2 string

//go:embed powershell-scripts/getLogs-scm.ps1
var getLogs_scm string

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

//go:embed powershell-scripts/development/enable-rdp.ps1
var enable_rdp string

//go:embed powershell-scripts/development/install-chocolatey.ps1
var install_chocolatey string

//go:embed powershell-scripts/development/install-core-gui.ps1
var install_core_gui string

//go:embed powershell-scripts/development/install-docker.ps1
var install_docker string

//go:embed powershell-scripts/development/install-git.ps1
var install_git string

//go:embed powershell-scripts/development/install-go.ps1
var install_go string

type cmd struct {
	filename          string
	name              string
	content           string
	developmentScript bool
}

var cmds = []cmd{{
	content:           get_cpu_load,
	developmentScript: false,
	filename:          "get-cpu-load.ps1",
	name:              "get-cpu-load",
}, {
	content:           get_redirected_disks,
	developmentScript: false,
	filename:          "get-redirected-disks.ps1",
	name:              "get-redirected-disks",
}, {
	content:           get_service_start_times,
	developmentScript: false,
	filename:          "get-service-start-times.ps1",
	name:              "get-service-start-times",
}, {
	content:           getLogs_rancherwins,
	developmentScript: false,
	filename:          "getLogs-rancherwins.ps1",
	name:              "getLogs-rancherwins",
}, {
	content:           getLogs_rke2,
	developmentScript: false,
	filename:          "getLogs-rke2.ps1",
	name:              "getLogs-rke2",
}, {
	content:           getLogs_scm,
	developmentScript: false,
	filename:          "getLogs-scm.ps1",
	name:              "getLogs-scm",
}, {
	content:           set_rdp_port,
	developmentScript: false,
	filename:          "set-rdp-port.ps1",
	name:              "set-rdp-port",
}, {
	content:           show_hotfixes,
	developmentScript: false,
	filename:          "show-hotfixes.ps1",
	name:              "show-hotfixes",
}, {
	content:           show_rdp_port,
	developmentScript: false,
	filename:          "show-rdp-port.ps1",
	name:              "show-rdp-port",
}, {
	content:           watch_rke2_service,
	developmentScript: false,
	filename:          "watch-rke2-service.ps1",
	name:              "watch-rke2-service",
}, {
	content:           watch,
	developmentScript: false,
	filename:          "watch.ps1",
	name:              "watch",
}, {
	content:           enable_rdp,
	developmentScript: true,
	filename:          "enable-rdp.ps1",
	name:              "enable-rdp",
}, {
	content:           install_chocolatey,
	developmentScript: true,
	filename:          "install-chocolatey.ps1",
	name:              "install-chocolatey",
}, {
	content:           install_core_gui,
	developmentScript: true,
	filename:          "install-core-gui.ps1",
	name:              "install-core-gui",
}, {
	content:           install_docker,
	developmentScript: true,
	filename:          "install-docker.ps1",
	name:              "install-docker",
}, {
	content:           install_git,
	developmentScript: true,
	filename:          "install-git.ps1",
	name:              "install-git",
}, {
	content:           install_go,
	developmentScript: true,
	filename:          "install-go.ps1",
	name:              "install-go",
}}
