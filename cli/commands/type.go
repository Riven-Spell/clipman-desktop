package commands

type Command func(args []string)

var Aliases = map[string]Command{
	"help":help,
	"auth":auth,
	"clip":clip,
}

var HelpList = map[string]string{
	"help": `Displays this list.`,
	"auth": `auth user:
Attempts to log in with current user credentials.

auth admin:
Attempts to log in with current admin credentials.

auth pass [admin/user] [password]:
Updates the details on file for the admin or user.`,
	"config": `config server [ip:port]:
Sets the server IP and port.

config buffer [size]:
Sets the buffer size.`,
	"clip": `clip force [push/refresh]:
Forcefully pushes or refreshes the clipboard to/from the server.

clip recheck:
Forcefully checks the clipboard for an update.

clip empty:
Empties the clipboard.`,
}