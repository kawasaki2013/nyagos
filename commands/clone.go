package commands

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/zetamatta/nyagos/dos"
	"github.com/zetamatta/nyagos/shell"
)

func getwd_() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dos.NetDriveToUNC(wd)
}

func clone_(action string, out io.Writer) (int, error) {
	wd := getwd_()
	var err error
	var me string
	me, err = os.Executable()
	if err != nil {
		return 1, err
	}
	err = dos.ShellExecute(action, me, "", wd)
	if err != nil {
		err = dos.ShellExecute(action, dos.TruePath(me), "", wd)
	}
	if err != nil {
		err2 := dos.ShellExecute(action, "CMD.EXE", "/c \""+me+"\"", wd)
		if err2 != nil {
			return 1, err // return original error
		}
	}
	return 0, nil
}

func cmd_sudo(ctx context.Context, cmd *shell.Cmd) (int, error) {
	if len(cmd.Args) < 2 {
		return clone_("runas", cmd.Stderr)
	}
	rawargs := cmd.RawArgs
	var args string
	if len(rawargs) >= 3 {
		args = strings.Join(rawargs[2:], " ")
	} else {
		args = ""
	}
	err := dos.ShellExecute("runas", dos.TruePath(cmd.Args[1]), args, getwd_())
	if err != nil {
		return 1, err
	} else {
		return 0, nil
	}
}

func cmd_clone(ctx context.Context, cmd *shell.Cmd) (int, error) {
	return clone_("open", cmd.Stderr)
}

func cmd_su(ctx context.Context, cmd *shell.Cmd) (int, error) {
	return clone_("runas", cmd.Stderr)
}
