package main

import (
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"

	"github.com/codyaray/proto-go-setter/setter"
)

type SetterFile struct {
	All      bool
	Messages []*SetterMessage
}

type SetterMessage struct {
	Name   string
	All    bool
	Fields []*SetterField
}

type SetterField struct {
	Name    string
	VarName string
	VarType string
}

func SetterAST(file *generator.FileDescriptor) (*SetterFile, bool) {
	hasInclude := false
	pkg := &SetterFile{
		All: proto.GetBoolExtension(file.Options, setter.E_AllMessages, false),
	}
	for _, message := range file.Messages() {
		m := &SetterMessage{
			Name: *message.Name,
			All:  proto.GetBoolExtension(message.Options, setter.E_AllFields, pkg.All),
		}
		pkg.Messages = append(pkg.Messages, m)
		for _, field := range message.GetField() {
			include := m.All
			if proto.GetBoolExtension(field.Options, setter.E_Exclude, false) {
				include = false
			} else if proto.GetBoolExtension(field.Options, setter.E_Include, false) {
				include = true
			}
			if include {
				f := &SetterField{
					Name:    generator.CamelCase(*field.Name),
					VarName: *field.Name,
					VarType: normalizeFieldType(field.Type.String()),
				}
				m.Fields = append(m.Fields, f)
				hasInclude = true
			}
		}
	}
	return pkg, hasInclude
}

func normalizeFieldType(t string) string {
	return strings.ToLower(strings.Split(t, "_")[1])
}
