# reference documentation: https://learn.microsoft.com/en-us/windows-server/get-started/server-core-app-compatibility-feature-on-demand
# used to install explorer.exe on a windows server core. It's much easier to manage files and redirected disks with this added UI.

Install-WindowsFeature -Name Failover-Clustering -IncludeManagementTools

Add-WindowsCapability -Online -Name ServerCore.AppCompatibility~~~~0.0.1.0

Write-Host "Restarting system in 10 seconds"

Start-Sleep -s 10

Restart-Computer