# rancher-powershell-utils
Utility scripts for working with rancher on Windows.

This repository packages several Powershell utility scripts into a single binary, 

allows you to execute a single installation script and automatically install several utility
Powershell files that are useful when debugging Windows applications. Several of these files are specific to Rancher, making log collection and issue diagnosis much easier. 

### Installation

There are two methods of installing these scripts onto a host 

1. Download and execute the `install.ps1` script on the desired host. This will download the required binary from github and automatically update the Path temporarily. If desired, the default Powershell profile can also be updated to permanently set the Path (when executed with the `-useProfile $true` argument). 

2. Directly downloading and executing the binary on the desired host. This will write files to disk, but will not update the Path or any Powershell profiles. 

#### Flags

```
 install.ps1
    - addToProfile [Bool]
        - specifies that the default Powershell profile (AllUsersCurrentHost) will be updated to automatically include the scripts in the PATH. 
    - utilFilesPath [String]
        - specifies where the utility scripts should be written to on the disk. By default this is '$env:USERPROFILE\AppData\Local\Temp\Rancher
    - utilsVesrion [String]
        - the release version of the binary which writes scripts to disk. This should equal the release name exactly (e.g. v1.1). 

```

```
 rancher-powershell-utilities.exe
    - script-path [String]
        - specifies where the utility scripts should be written to on disk. By default this is '$env:USERPROFILE\AppData\Local\Temp\Rancher'
    - add-development-scripts [Bool]
        - specifies if scripts focused on development should be written to disk. This is disabled by default. For a list of all development scripts, refer to 'powershell-scripts/development' 
```