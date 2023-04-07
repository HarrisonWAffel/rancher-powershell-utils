param([Parameter(Mandatory=$false)][Int32]$PortNumber=3389)

Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp' -Name 'PortNumber' -Value $PortNumber
