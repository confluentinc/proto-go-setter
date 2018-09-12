package main

import (
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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

func SetterAST(file *generator.FileDescriptor, g *generator.Generator) (*SetterFile, bool) {
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
					VarType: normalizeVarType(field, g),
				}
				m.Fields = append(m.Fields, f)
				hasInclude = true
			}
		}
	}
	return pkg, hasInclude
}

func normalizeVarType(field *descriptor.FieldDescriptorProto, g *generator.Generator) string {
	var varType string
	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		varType = normalizeMessageVarType(field, g)
	default:
		varType = strings.ToLower(strings.Split(field.Type.String(), "_")[1])
	}
	return varType
}

func normalizeMessageVarType(field *descriptor.FieldDescriptorProto, g *generator.Generator) string {
	desc := g.ObjectNamed(field.GetTypeName())
	if d, ok := desc.(*generator.Descriptor); ok && d.GetOptions().GetMapEntry() {
		// Figure out the Go types for the key and value types.
		keyField, valField := d.Field[0], d.Field[1]
		keyType, _ := g.GoType(d, keyField)
		valType, _ := g.GoType(d, valField)
		return fmt.Sprintf("map[%s]%s", keyType, valType)
	}
	return "" // TODO non-map message types: https://github.com/golang/protobuf/blob/e344474228f55cc8e5f2342b0493b71bcdd258f4/protoc-gen-go/generator/generator.go#L2525
}
