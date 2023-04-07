//go:build windows

package main

// this file just writes embedded files to disk and adds them to the powershell profile
// This is useful as it allows us to do single curl and get all the files at once
//
// to add new files into the binary `go generate`

func main() {
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
	f, err := os.OpenFile("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\profile.ps1", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
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
