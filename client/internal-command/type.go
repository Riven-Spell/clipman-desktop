package internal_command

type Command struct {
	Cmd    CommandID
	Params []string
}

type CommandID uint8

const (
	CONNECT CommandID = iota //resets the connection.
	DISCONNECT

	AUTH_USER //Requires an active connection.
	AUTH_ADMIN

	PUSH_CLIP //Requires authentication.
	REFRESH_CLIP

	RCON_COMMAND
)

var Commands = map[CommandID]func([]string){
	CONNECT:      Connect,
	DISCONNECT:   Disconnect,
	AUTH_USER:    AuthUser,
	AUTH_ADMIN:   AuthAdmin,
	PUSH_CLIP:    PushClip,
	REFRESH_CLIP: RefreshClip,
	RCON_COMMAND: Rcon,
}

func CmdToBytes(cmd Command) []byte {
	out := []byte{byte(cmd.Cmd)}

	for _, v := range cmd.Params {
		//len []byte(v) in case go treats multibyte UTF-8 characters as one byte
		out = append(out, uint32ToBytes(uint32(len([]byte(v))))...)
		out = append(out, []byte(v)...)
	}

	return out
}

func uint32ToBytes(in uint32) []byte {
	out := []byte{
		uint8(((uint32(255) << uint(8*3)) & in) >> uint(8*3)),
		uint8(((uint32(255) << uint(8*2)) & in) >> uint(8*2)),
		uint8(((uint32(255) << uint(8*1)) & in) >> uint(8*1)),
		uint8(uint32(255) & in)}

	return out
}
