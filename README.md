# rancher-powershell-utils
utility commands for working with rancher on Windows. This repository allows you to execute a single binary and automatically install several powerscript utility functions that can be used to debug rancher-wins or other windows processes.

To add new scripts simply place new powershell files into the `powershell-scripts` directory and then run `go genereate` and `go build`. The resulting binary will have all the scripts embedded into it and will be able to write them to disk on execution.