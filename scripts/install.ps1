$VERSION = "v1.1"

# download the binary
wget "https://github.com/harrisonwaffel/rancher-powershell-utils/releases/$VERSION/rancher-windows-utilities.exe" -outfile rancher-utilities.exe

# run the binary and install the embedded pwsh scripts
# pipe the output into powershell so the env var is set properly
./rancher-utilities.exe | powershell

