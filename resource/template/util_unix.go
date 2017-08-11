// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package template

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// fileStat return a fileInfo describing the named file.
func fileStat(name string) (fi fileInfo, err error) {
	if isFileExist(name) {
		f, err := os.Open(name)
		if err != nil {
			return fi, err
		}
		defer f.Close()
		stats, _ := f.Stat()
		fi.Uid = stats.Sys().(*syscall.Stat_t).Uid
		fi.Gid = stats.Sys().(*syscall.Stat_t).Gid
		fi.Mode = stats.Mode()
		h := md5.New()
		io.Copy(h, f)
		fi.Md5 = fmt.Sprintf("%x", h.Sum(nil))
		return fi, nil
	} else {
		return fi, errors.New("File not found")
	}
}

func command(cmd string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", cmd)
}
