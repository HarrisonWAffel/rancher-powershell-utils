package main

import _ "embed"

//go:embed powershell-scripts/cpu-load-percentage.ps1
var cpuLoadPercentage string

//go:embed powershell-scripts/get-redirected-disks.ps1
var getRedirectedDisks string

//go:embed powershell-scripts/get-service-start-times.ps1
var serviceStartTimes string

//go:embed powershell-scripts/get-winevent-formatted-message.ps1
var formattedWinEvent string

//go:embed powershell-scripts/install-core-gui.ps1
var installCoreGUI string

//go:embed powershell-scripts/rke2-logs.ps1
var rke2Logs string

//go:embed powershell-scripts/watch.ps1
var watch string

//go:embed powershell-scripts/rancherwins-logs.ps1
var winsLogs string

//go:embed powershell-scripts/scm-logs.ps1
var scmLogs string
