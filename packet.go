package main

import "encoding/json"

type Packet struct {
	Cmd string `json:"cmd"`
	Arg string `json:"arg"`
}

func (p Packet) ToBytePack() []byte {
	bytepck, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return append(bytepck, []byte("\n")[0])
}

func FromBytes(b []byte) (Packet, error) {
	p := Packet{}
	err := json.Unmarshal(b, &p)
	return p, err
}

func ErrorPacket(s string) Packet {
	return Packet{
		Cmd: "error",
		Arg: s,
	}
}

func ResponsePacket(s string) Packet {
	return Packet{
		Cmd: "response",
		Arg: s,
	}
}
