package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := flag.String("script-path", fmt.Sprintf("%s\\AppData\\Local\\Temp\\Rancher", os.Getenv("USERPROFILE")), "Where the embedded scripts will be written to")
	addDevScripts := flag.Bool("add-development-scripts", false, "Adds scripts focusing on development to the host")

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

	fmt.Println(getUpdatePathCommand(fullPath))
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
	}

	return nil
}
