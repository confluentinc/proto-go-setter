package main

import "github.com/gogo/protobuf/vanity/command"

func main() {
	req := command.Read()
	p := NewGenerator()
	rep := command.GeneratePlugin(req, p, "_setter.go")
	if p.ShouldWrite() {
		command.Write(rep)
	}
}
