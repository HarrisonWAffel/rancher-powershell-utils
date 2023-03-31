//go:build windows

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// this file just writes embedded files to disk and adds them to the path
// it exists so we can do a single curl and get all the files at once
// it also adds all the commands to your profile if they aren't there already

type cmd struct {
	filename string
	name     string
	content  string
}

func main() {
	var cmds = []cmd{
		{
			filename: "rancher-wins-logs.ps1",
			name:     "rancher-wins-logs",
			content:  winsLogs,
		},
		{
			filename: "watch.ps1",
			name:     "watch",
			content:  watch,
		},
		{
			filename: "rke2-logs.ps1",
			name:     "rke2-logs",
			content:  rke2Logs,
		},
		{
			filename: "install-core-gui.ps1",
			name:     "install-gui",
			content:  installCoreGUI,
		},
		{
			filename: "formatted-winevent.ps1",
			name:     "format-winevent",
			content:  formattedWinEvent,
		},
		{
			filename: "get-service-start-times.ps1",
			name:     "service-start-times",
			content:  serviceStartTimes,
		},
		{
			filename: "get-redirected-disks.ps1",
			name:     "redirected-disks",
			content:  getRedirectedDisks,
		},
		{
			filename: "get-cpu-load.ps1",
			name:     "cpu-load",
			content:  cpuLoadPercentage,
		},
		{
			filename: "scm-logs.ps1",
			name:     "scm-logs",
			content:  scmLogs,
		},
	}

	// all of these files are going to
	// be written in the same directory
	// as the executable
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	err = writeFiles(cmds, exPath)
	if err != nil {
		panic(err)
	}

	if err = setScriptAlias(cmds, exPath); err != nil {
		panic(err)
	}
}

func setScriptAlias(cmds []cmd, filesPath string) error {
	// creates or appends to a profile.ps1 file and adds
	// aliases to the newly written commands
	found := true
	f, err := os.Open("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\profile.ps1")
	if errors.Is(err, os.ErrNotExist) {
		found = false
	} else if err != nil {
		return err
	}

	if !found {
		// create new file
		f, err = os.Create("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\profile.ps1")
		if err != nil {
			return err
		}
	}

	currentFile, err := os.ReadFile("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\profile.ps1")
	if err != nil {
		return err
	}

	for _, c := range cmds {
		if strings.Contains(string(currentFile), fmt.Sprintf("Set-Alias %s %s\n", c.name, filepath.Join(filesPath, c.filename))) {
			continue
		}
		_, err = f.WriteString(fmt.Sprintf("Set-Alias %s %s\n", c.name, filepath.Join(filesPath, c.filename)))
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func writeFiles(cmd []cmd, wd string) error {
	var res []string
	for _, c := range cmd {
		// write files to disk
		f, err := os.Create(c.filename)
		if err != nil {
			return err
		}

		_, err = f.WriteString(c.content)
		if err != nil {
			return err
		}

		st, _ := f.Stat()
		res = append(res, filepath.Join(wd, st.Name()))
		if err = f.Close(); err != nil {
			return err
		}
	}

	return nil
}
