package main

import (
	"html/template"
	"bytes"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type Generator struct {
	*generator.Generator
	generator.PluginImports
	write bool
}

func NewGenerator() *Generator {
	return &Generator{write: false}
}

func (p *Generator) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *Generator) Name() string {
	return "setter"
}

func (p *Generator) ShouldWrite() bool {
	return p.write
}

func (p *Generator) Generate(file *generator.FileDescriptor) {
	t := template.Must(template.New("setter").Parse(tmpl))

	pkg, include := SetterAST(file, p.Generator)
	if include {
		var buf bytes.Buffer
		t.Execute(&buf, pkg)
		p.P(buf.String())
		p.write = true
	}
}

func (p *Generator) GenerateImports(file *generator.FileDescriptor) {
	// nada
}

func init() {
	generator.RegisterPlugin(NewGenerator())
}
