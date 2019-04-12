package external_command

import (
	"reflect"
	"testing"
)

func TestBytesToUint32(t *testing.T) {
	testBytes := [][]byte{{136, 17}, {136, 136, 136, 136}, {153, 102, 102, 153}}
	testResults := []uint32{4488, 2290649224, 2573624985}

	passed := true
	for k, v := range testBytes {
		if o := bytesToUint32(v...); o != testResults[k] {
			t.Error("FAILED TestBytesToUint32!")
			t.Error("Expected:", testResults[k])
			t.Error("Got:", o)

			passed = false
		}
	}

	if !passed {
		t.Fail()
	}
}

func TestParseCmdZero(t *testing.T) {
	cmdZero := []byte{0,
		1, 0, 0, 0,
		65}
	expected := Command{Cmd: 0, Params: []string{"A"}}

	if cmd := ParseCmd(cmdZero); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdZero!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdOnly(t *testing.T) {
	cmdOnly := []byte{3}
	expected := Command{Cmd: 3, Params: make([]string, 0)}

	if cmd := ParseCmd(cmdOnly); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOnly!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdOneArg(t *testing.T) {
	cmdOneArg := []byte{3,
		3, 0, 0, 0,
		65, 65, 65}
	expected := Command{Cmd: 3, Params: []string{"AAA"}}

	if cmd := ParseCmd(cmdOneArg); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOneArg!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdOneArgNEArg(t *testing.T) {
	cmdOneArgLFEnd := []byte{3,
		1, 0, 0, 0,
		65,
		10, 0, 0, 0}
	expected := Command{Cmd: 3, Params: []string{"A"}}

	if cmd := ParseCmd(cmdOneArgLFEnd); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOneArgNEArg! (1 arg + nonexistent arg)")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdOneArgTLArg(t *testing.T) {
	cmdOneArgLFEnd := []byte{3,
		1, 0, 0, 0,
		65,
		10, 0, 0, 0,
		65}
	expected := Command{Cmd: 3, Params: []string{"A"}}

	if cmd := ParseCmd(cmdOneArgLFEnd); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdOneArgTLArg! (1 arg + too long arg, not accepted)")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}

func TestParseCmdTwoArg(t *testing.T) {
	cmdTwoArg := []byte{3,
		1, 0, 0, 0,
		65,
		1, 0, 0, 0,
		65}
	expected := Command{Cmd: 3, Params: []string{"A", "A"}}

	if cmd := ParseCmd(cmdTwoArg); !reflect.DeepEqual(cmd, expected) {
		t.Error("FAILED cmdTwoArg!")
		t.Error("Expected:", expected)
		t.Error("Got:", cmd)

		t.Fail()
	}
}
