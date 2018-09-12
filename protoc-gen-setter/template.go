package main

var tmpl = `
{{ range $message := .Messages }}
{{ range $field := $message.Fields }}
func (t *{{ $message.Name }}) Set{{ $field.Name }}({{ $field.VarName }} {{ $field.VarType }}) {
    t.{{ $field.Name }} = {{ $field.VarName }}
}
{{ end }}

{{ end }}
`
