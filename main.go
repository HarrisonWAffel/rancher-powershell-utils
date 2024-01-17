package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// this file just writes embedded files to disk and adds them to the powershell profile
// This is useful as it allows us to do single curl and get all the files at once
//
// to add new files into the binary create a new ps1 file in the 'powershell-scripts' directory and run `go generate`

func main() {
	fullPath := fmt.Sprintf("%s\\AppData\\Local\\Temp", os.Getenv("USERPROFILE"))
	// write all files to a temp directory on dis
	err := writeFiles(cmds, fullPath)
	if err != nil {
		panic(err)
	}

	if err = updatePATH(fullPath); err != nil {
		panic(err)
	}
}

func updatePATH(fullPath string) error {
	if strings.Contains(os.Getenv("PATH"), fullPath) {
		return nil
	}
	//	cmd := exec.Command("powershell", fmt.Sprintf(`
	//[Environment]::SetEnvironmentVariable(
	//    "Path",
	//    [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";%s",
	//    [EnvironmentVariableTarget]::Machine)
	//`, fullPath))
	//
	//	err := os.Setenv("PATH", fmt.Sprintf("$env:PATH;%s", fullPath))
	//	if err != nil {
	//		return err
	//	}

	//cmd := exec.Command("powershell", )
	//fmt.Println(cmd.String())
	//o, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Println("cmd output: ", string(o))
	//	return err
	//}
	fmt.Println(fmt.Sprintf(`$env:PATH+=";%s"`, fullPath))
	return nil
}

func writeFiles(cmd []cmd, wd string) error {
	for _, c := range cmd {
		// write files to disk
		f, err := os.Create(filepath.Join(wd, c.filename))
		if err != nil {
			return err
		}

		_, err = f.WriteString(c.content)
		if err != nil {
			return err
		}

		if err = f.Close(); err != nil {
			return err
		}
	}

	return nil
}
