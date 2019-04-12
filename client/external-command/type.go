package external_command

type Command struct {
	Cmd    byte
	Params []string
}

var Commands = map[byte]func(args []string){
	0: SetClip,
	1: Success,
	2: Failure,
} //Pointers because some commands can elevate privileges.
//For now, clients will have just one command available to the server: 0 updateClip

/*
Command structure:
[CMD,
<4 bytes, pushed to 32-bit integer leftwards-- 1, 0, 0, 0 == 1, describes parameter length. Do not include or leave as 0 for no parameters.>,
<parameter itself-- if smaller than allocated buffer, pad with null>
<repeat>
*/

/*
Server commands:
0: becomeAdmin
1: rcon
2: getClip
3: setClip
4: login

Client commands:
0 setClip
1 success
2 failure
*/
func ParseCmd(buffer []byte) Command {
	cmd := Command{Cmd: buffer[0], Params: make([]string, 0)}

	i := 1
	for {
		if len(buffer) < i+4 {
			break
		}

		arglen := bytesToUint32(buffer[i], buffer[i+1], buffer[i+2], buffer[i+3])

		if len(buffer) < i+4+int(arglen) {
			break
		}

		cmd.Params = append(cmd.Params, string(buffer[i+4:i+4+int(arglen)]))

		i += int(arglen) + 4
	}

	return cmd
}

func bytesToUint32(bytes ...byte) uint32 {
	var out uint32 = 0

	for k, v := range bytes {
		out |= uint32(v) << uint(8*k)
	}

	return out
}
