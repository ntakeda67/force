package command

import (
	"flag"
	"fmt"
	"strings"
)

var flagEnv string
var flagProcfile string

var Commands = []*Command{
	cmdApiVersion,
	cmdUseDXAuth,
	cmdLogin,
	cmdLogout,
	cmdLogins,
	cmdCreate,
	cmdActive,
	cmdWhoami,
	cmdDescribe,
	cmdSobject,
	cmdBigObject,
	cmdField,
	cmdRecord,
	cmdBulk,
	cmdFetch,
	cmdImport,
	cmdExport,
	cmdQuery,
	cmdApex,
	cmdGet,
	cmdPost,
	cmdSession,
	cmdTrace,
	cmdLog,
	cmdEventLogFile,
	cmdOauth,
	cmdTest,
	cmdSecurity,
	cmdVersion,
	cmdUpdate,
	cmdPush,
	cmdAura,
	cmdPassword,
	cmdNotifySet,
	cmdLimits,
	cmdHelp,
	cmdDataPipe,
}

type Command struct {
	// args does not include the command name
	Run  func(cmd *Command, args []string)
	Flag flag.FlagSet

	Usage string // first word is the command name
	Short string // `forego help` output
	Long  string // `forego help cmd` output
}

func (c *Command) PrintUsage() {
	if c.Runnable() {
		fmt.Printf("Usage: force %s\n\n", c.Usage)
	}
	fmt.Println(strings.Trim(c.Long, "\n"))
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

func (c *Command) List() bool {
	return c.Short != ""
}