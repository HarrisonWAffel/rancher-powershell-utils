# This script is the powershell equivilent to the linux 'watch' command

# n - number of seconds between refreshes. Default is 2
param([Parameter(Mandatory=$false)][Int32]$n=2)

while (1) {
    clear
    # setup the header
    Write-Host "watch -n $n $args : $(Get-Date) `n---------------"

    # the command to be executed
    Invoke-Expression "$args"

    Start-Sleep -s $n
}