package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var l log

func main() {
	path := flag.String("script-path", fmt.Sprintf("%s\\AppData\\Local\\Temp\\Rancher", os.Getenv("USERPROFILE")), "Where the embedded scripts will be written to")
	addDevScripts := flag.Bool("add-development-scripts", false, "Adds scripts focusing on development to the host")
	executedFromInstallScript := flag.Bool("installed-via-script", false, "Disables specific logs when run via the install script. When enabled, the only output will be the command needed to add these scripts to the path")
	l.executedFromInstallScript = *executedFromInstallScript

	flag.Parse()
	fullPath := *path
	devScripts := *addDevScripts

	err := createTempDirectory(fullPath)
	if err != nil {
		panic(err)
	}

	err = writeFiles(cmds, fullPath, devScripts)
	if err != nil {
		panic(err)
	}

	l.print("Install complete! Run the following command to add the scripts to your PATH")
	fmt.Println(getUpdatePathCommand(fullPath))
}

type log struct {
	executedFromInstallScript bool
}

func (l log) print(s string) {
	if l.executedFromInstallScript {
		return
	}
	fmt.Println(s)
}

func getUpdatePathCommand(fullPath string) string {
	return fmt.Sprintf(`$env:PATH+=";%s"`, fullPath)
}

func createTempDirectory(fullPath string) error {
	err := os.Mkdir(fullPath, os.ModePerm)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return fmt.Errorf("failed to create temporary scripts directory (%s): %w", fullPath, err)
	}
	return nil
}

func writeFiles(cmd []cmd, wd string, addDev bool) error {
	for _, c := range cmd {
		if c.developmentScript && !addDev {
			continue
		}

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

		l.print(fmt.Sprintf("Added script %s to %s", c.name, wd))
	}

	return nil
}
