package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"os/user"
	"runtime"
	"bytes"
		"strings"
	"os"
	"errors"
	"os/exec"
)

func GetSha256(source string) string {
	h := sha256.New()
	h.Write([]byte(source))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home(isDesktop bool) (string, error) {
	path := ""
	if isDesktop {
		path = "/Desktop"
	}
	user, err := user.Current()
	if nil == err {
		return user.HomeDir + path, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	if isDesktop {
		path = "~/Desktop"
	} else {
		path = "~$USER"
	}
	// Unix-like system, so just assume Unix
	return homeUnix("~/Desktop")
}

func homeUnix(p string) (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	//cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd := exec.Command("sh", "-c", "eval echo " + p)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
