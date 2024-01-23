param([Parameter(Mandatory=$false)][Int32]$PortNumber=3389)

# Enable Remote Desktop connection
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Terminal Server' -Name 'fDenyTSConnections' -Value 0

# Set Remote Desktop service starting automatically
Set-Service TermService -StartupType Automatic -Status Running -PassThru

# Create firewall rules
New-NetFirewallRule -DisplayName 'Remote Desktop - Custom Mode (TCP-In)' -Name 'RemoteDesktop-CustomMode-In-TCP' -Direction Inbound -Protocol TCP -LocalPort $PortNumber -Action Allow
New-NetFirewallRule -DisplayName 'Remote Desktop - Custom Mode (UDP-In)' -Name 'RemoteDesktop-CustomMode-In-UDP' -Direction Inbound -Protocol UDP -LocalPort $PortNumber -Action Allow