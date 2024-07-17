# Setup a node to install a custom rke2 binary on Windows

param(
    [Parameter(Mandatory=$true)]
    [String]$ServerIP="",

    [Parameter(Mandatory=$true)]
    [String]$Token="",

    [Parameter(Mandatory=$true)]
    [String]$CustomBinaryLocation=""
)

if ($ServerIP -eq "") {
    Write-Host "Must specify a valid Server IP address"
    exit 1
}

if ($Token -eq "") {
    Write-Host "Must specify a valid server token"
    exit 1
}

if ($CustomBinaryLocation -eq "") {
    Write-Host "Must specify the location of the custom RKE2 binary"
    exit 1
}

# Download and Run the rke2 install command, as it configures a few things on the host. We will replace
# the rke2 binary with our custom one later on
Invoke-WebRequest -Uri https://raw.githubusercontent.com/rancher/rke2/master/install.ps1 -Outfile install.ps1
./install.ps1

# Create the configuration file which points to the server using a specified token
New-Item -Type Directory c:/etc/rancher/rke2 -Force
Set-Content -Path c:/etc/rancher/rke2/config.yaml -Value @"
server: https://$ServerIP`:9345
token: $Token
"@

# Configure the path
$env:PATH+=";c:\var\lib\rancher\rke2\bin;c:\usr\local\bin"
[Environment]::SetEnvironmentVariable(
    "Path",
    [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";c:\var\lib\rancher\rke2\bin;c:\usr\local\bin",
    [EnvironmentVariableTarget]::Machine)

# Delete the existing binary and nove the custom one into place
rm /usr/local/bin/rke2.exe
mv $CustomBinaryLocation /usr/local/bin

# Verify
cat /etc/rancher/rke2/config.yaml
ls /usr/local/bin
rke2 --version