package external_command

import (
	"reflect"
	"testing"
)

func TestParseCmdZero(t *testing.T) {
	cmdZero := []byte{0, 10, 65, 0}
	expected := Command{Cmd: 0, Args: []string{"A"}}

	if cmd := ParseCmd(cmdZero); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdZero!")
		t.Error("Expected:", expected)
		t.Error("Got:")
	}
}

func TestParseCmdOnly(t *testing.T) {
	cmdOnly := []byte{3, 0}
	expected := Command{Cmd: 3, Args: make([]string, 0)}

	if cmd := ParseCmd(cmdOnly); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOnly!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdLF(t *testing.T) {
	cmdLF := []byte{10, 0}
	expected := Command{Cmd: 10, Args: make([]string, 0)}

	if cmd := ParseCmd(cmdLF); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdLF!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdOneArg(t *testing.T) {
	cmdOneArg := []byte{3, 10, 65, 65, 65, 0}
	expected := Command{Cmd: 3, Args: []string{"AAA"}}

	if cmd := ParseCmd(cmdOneArg); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOneArg!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdOneArgFEnd(t *testing.T) {
	cmdOneArgLFEnd := []byte{3, 10, 65, 10}
	expected := Command{Cmd: 3, Args: []string{"A"}}

	if cmd := ParseCmd(cmdOneArgLFEnd); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOneArgLFEnd!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdTwoArg(t *testing.T) {
	cmdTwoArg := []byte{3, 10, 65, 10, 65, 0}
	expected := Command{Cmd: 3, Args: []string{"A", "A"}}

	if cmd := ParseCmd(cmdTwoArg); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdTwoArg!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdUnexpectedArg(t *testing.T) {
	cmdUnexpectedArg := []byte{3, 10, 65, 10, 65}
	expected := Command{Cmd: 3, Args: []string{"A"}}

	if cmd := ParseCmd(cmdUnexpectedArg); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdUnexpectedArg!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

