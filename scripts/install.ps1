param(
    [Boolean]$addToProfile=$false,
    [Boolean]$deleteFromProfile=$false,
    [String]$utilFilesPath="$env:USERPROFILE\AppData\Local\Temp\Rancher",
    [String]$UtilsVersion="v1.1"
)

$ErrorActionPreference = 'Stop'

if (!(Test-Path "rancher-powershell-utilities.exe")) {
    # download the binary
    wget "https://github.com/harrisonwaffel/rancher-powershell-utils/releases/$UtilsVersion/rancher-powershell-utilities.exe" -outfile rancher-utilities.exe
}

if ($useProfile) {
    # update the default powershell profile. Useful for
    # ensuring you have access to debugging scripts across sessions
    # and reboots
    $updatedPath = $(./rancher-powershell-utilities.exe --script-path $utilFilesPath)
    if ($updatedPath -contains "panic") {
        Write-Host "Error encountered staging scripts $updatedPath"
        exit 1
    }

    if (![System.IO.File]::Exists($PROFILE.AllUsersCurrentHost)) {
        New-Item -ItemType File -Force -Path $PROFILE.AllUsersCurrentHost
    }

    if ((Select-String -Path $PROFILE.AllUsersCurrentHost `$updatedPath) -eq $null) {
        Add-Content -Path $PROFILE.AllUsersCurrentHost -Value $updatedPath
    } else {
        Write-Host "Current profile already contains correct PATH extension"
    }

    Write-Host "Reloading profile..."

    . $PROFILE.AllUsersCurrentHost

    $files = Get-ChildItem $env:USERPROFILE\AppData\Local\Temp\Rancher
    Write-Host The following files have been successfully written to disk and permenantly added to the default PowerShell Profile
    Write-Host ------
    for ($i=0; $i -lt $files.Count; $i++) {
        Write-Host $files[$i]
    }
} else {
    # Update the Path temporarily. Useful for short term debugging
    # on non-development machines, or machines which do not need
    # these scripts forever
    $updatedPath = $(./rancher-powershell-utilities.exe --script-path $utilFilesPath)
    if ($updatedPath -contains "panic") {
        Write-Host "Error encountered staging scripts $updatedPath"
        exit 1
    }

    if ($updatedPath) {
        Write-Host $updatedPath
        Invoke-Expression $updatedPath
    }

    $files = Get-ChildItem $utilFilesPath
    Write-Host The following files have been successfully written to disk and temporarily added to the PATH
    Write-Host ------
    for ($i=0; $i -lt $files.Count; $i++) {
        Write-Host $files[$i]
    }
}
