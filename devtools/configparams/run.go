package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/target/goalert/config"
)

func hasType(typeName string, fields []field) bool {
	for _, f := range fields {
		if f.Type == typeName {
			return true
		}
	}
	return false
}

var tmpl = template.Must(
	template.
		New("mapconfig.go").
		Funcs(template.FuncMap{
			"quote":      strconv.Quote,
			"hasBool":    func(fields []field) bool { return hasType("ConfigTypeBoolean", fields) },
			"hasStrList": func(fields []field) bool { return hasType("ConfigTypeStringList", fields) },
			"hasInt":     func(fields []field) bool { return hasType("ConfigTypeInteger", fields) },
		}).
		Parse(`
import (
	"github.com/target/goalert/config"
	"github.com/target/goalert/validation"
)

func MapConfigHints(cfg config.Hints) []ConfigHint {
	return []ConfigHint{
		{{- range .HintFields }}
		{ID: {{quote .ID}}, DisplayName: {{quote .DisplayName}}, Value: {{.Value}}},
		{{- end}}
	}
}

// MapConfigValues will map a Config struct into a flat list of ConfigValue structs.
func MapConfigValues(cfg config.Config) []ConfigValue {
	return []ConfigValue{
		{{- range .ConfigFields }}
		{ID: {{quote .ID}},Type: {{.Type}}, DisplayName: {{quote .DisplayName}}, Description: {{quote .Desc}}, Value: {{.Value}}{{if .Password}}, Password: true{{end}}},
		{{- end}}
	}
}

// MapPublicConfigValues will map a Config struct into a flat list of ConfigValue structs.
func MapPublicConfigValues(cfg config.Config) []ConfigValue {
	return []ConfigValue{
		{{- range .ConfigFields }}
		{{- if .Public}}
		{ID: {{quote .ID}}, Type: {{.Type}}, DisplayName: {{quote .DisplayName}}, Description: {{quote .Desc}}, Value: {{.Value}}{{if .Password}}, Password: true{{end}}},
		{{- end}}
		{{- end}}
	}
}

// ApplyConfigValues will apply a list of ConfigValues to a Config struct.
func ApplyConfigValues(cfg config.Config, vals []ConfigValueInput) (config.Config, error) {
	{{- if hasStrList .ConfigFields}}
	parseStringList := func(v string) []string {
		if v == "" {
			return nil
		}
		return strings.Split(v, "\n")
	}
	{{- end}}
	{{- if hasInt .ConfigFields}}
	parseInt := func(id, v string) (int, error) {
		if v == "" {
			return 0, nil
		}
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, validation.NewFieldError("\""+id+"\".Value", "integer value invalid: " + err.Error())
		}
		return int(val), nil
	}
	{{- end}}
	{{- if hasBool .ConfigFields}}
	parseBool := func(id, v string) (bool, error) {
		switch v {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return false, validation.NewFieldError("\""+id+"\".Value", "boolean value invalid: expected 'true' or 'false'")
		}
	}
	{{- end}}
	for _, v := range vals {
		switch v.ID {
		{{- range .ConfigFields}}
		case {{quote .ID}}:
			{{- if eq .Type "ConfigTypeString"}}
			cfg.{{.ID}} = v.Value
			{{- else if eq .Type "ConfigTypeStringList"}}
			cfg.{{.ID}} = parseStringList(v.Value)
			{{- else if eq .Type "ConfigTypeInteger"}}
			val, err := parseInt(v.ID, v.Value)
			if err != nil {
				return cfg, err
			}
			cfg.{{.ID}} = val
			{{- else if eq .Type "ConfigTypeBoolean"}}
			val, err := parseBool(v.ID, v.Value)
			if err != nil {
				return cfg, err
			}
			cfg.{{.ID}} = val
			{{- end}}
		{{- end}}
		default:
			return cfg, validation.NewFieldError("ID", fmt.Sprintf("unknown config ID '%s'", v.ID))
		}
	}
	return cfg, nil
}
`))

type field struct {
	ID, DisplayName, Type, Desc, Value string
	Public, Password      bool
}

func main() {
	out := flag.String("out", "", "Output file.")
	flag.Parse()

	w := os.Stdout
	if *out != "" {
		fd, err := os.Create(*out)
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		w = fd
	}

	fmt.Fprintln(w, `// Code generated by devtools/configparams DO NOT EDIT.

package graphql2`)

	var input struct {
		ConfigFields []field
		HintFields   []field
	}
	input.ConfigFields = printType("", reflect.TypeOf(config.Config{}), "", "", false, false)
	input.HintFields = printType("", reflect.TypeOf(config.Hints{}), "", "", false, false)

	err := tmpl.Execute(w, input)
	if err != nil {
		panic(err)
	}
}

func printField(prefix string, f reflect.StructField) []field {
	fPrefix := prefix + f.Name + "."
	if f.Type.Kind() == reflect.Slice && f.Type.Elem().Kind() == reflect.Struct {
		fPrefix = prefix + f.Name + "[]."
	}
	return printType(fPrefix, f.Type, f.Tag.Get("displayName"), f.Tag.Get("info"), f.Tag.Get("public") == "true", f.Tag.Get("password") == "true")
}
func printType(prefix string, v reflect.Type, displayName string, info string, public, pass bool) []field {
	var f []field
	key := strings.TrimSuffix(prefix, ".")

	var typ, value string
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).PkgPath != "" {
				// skip unexported fields
				continue
			}
			f = append(f, printField(prefix, v.Field(i))...)
		}
		return f
	case reflect.Bool:
		typ = "ConfigTypeBoolean"
		value = fmt.Sprintf(`fmt.Sprintf("%%t", cfg.%s)`, key)
	case reflect.String:
		typ = "ConfigTypeString"
		value = fmt.Sprintf(`cfg.%s`, key)
	case reflect.Int:
		typ = "ConfigTypeInteger"
		value = fmt.Sprintf(`fmt.Sprintf("%%d", cfg.%s)`, key)
	case reflect.Slice:
		switch v.Elem().Kind() {
		case reflect.String:
			typ = "ConfigTypeStringList"
			value = fmt.Sprintf(`strings.Join(cfg.%s, "\n")`, key)
		default:
			panic(fmt.Sprintf("not implemented for type []%v", v.Elem().Kind()))
		}
	default:
		panic(fmt.Sprintf("not implemented for type %T", v.Kind()))
	}

	f = append(f, field{ID: key, Type: typ, DisplayName: printDisplayName(key, displayName), Desc: info, Value: value, Public: public, Password: pass})
	return f
}

func printDisplayName(key, displayName string) string {
	format := func(str string) string {
		var formattedStr []string
		pattern := regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")
		for _, subStr := range pattern.FindAllStringSubmatch(str, -1) {
			if subStr[1] != "" {
				formattedStr = append(formattedStr, subStr[1])
			}
			if subStr[2] != "" {
				formattedStr = append(formattedStr, subStr[2])
			}
		}
		return strings.Title(strings.Join(formattedStr, " "))
	}

	if displayName == "" {
		fieldKey := strings.Split(key, ".")
		displayName = format(fieldKey[len(fieldKey)-1])
	}

	return displayName
}
