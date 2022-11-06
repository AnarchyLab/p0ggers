package malfun

import (
	coldfire "github.com/redcode-labs/Coldfire"
)

func EXEC(command string) string {
	output, err := coldfire.CmdOut(command)
	if err != nil {
		return err.Error()
	}
	return output
}
