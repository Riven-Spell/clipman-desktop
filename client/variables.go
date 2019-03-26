package client

import (
	"github.com/virepri/clipman-desktop/client/external-command"
	"github.com/virepri/clipman-desktop/client/internal-command"
)

var Messages chan internal_command.Command
var ExternalMessages chan external_command.Command
