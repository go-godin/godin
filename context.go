package godin

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/vetcher/go-astra/types"
	"strings"
)

// Context is a container for everything which can be accessed from godin templates.
type Context struct {
	Package  string
	Services []Service
	Messages []Message
	Enums    []Enum
}

// GetMessage tries to resolve a given message name and return the actual message struct.
// The method will always return a Message, even if no message with the given name exists.
// In that case, an UNDEFINED message is returned in order to ensure that templates can still be processed.
// The casing of the name does not matter.
func (ctx Context) GetMessage(name string) Message {
	for _, msg := range ctx.Messages {
		if strings.ToLower(msg.Name) == strings.ToLower(name) {
			return msg
		}
	}
	return Message{Name: "UNDEFINED", Comments: []string{"Godin was unable to resolve this type"}}
}

func (ctx Context) IsEnum(name string) bool {
	for _, enum := range ctx.Enums {
		if strings.ToLower(enum.Name) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

// Service describes a gRPC service as defined in a protobuf file
type Service struct {
	Name     string
	Comments []string
	RPCs     []RPC
}

// RPC abstracts a RPC defined in a protobuf file
type RPC struct {
	Name         string
	Comments     []string
	RequestType  string
	ResponseType string
}

// Message abstracts an arbitrary protobuf 'message' struct
type Message struct {
	Name     string
	Comments []string
	Fields   []MessageField
}

// IsRequest returns true if the message is a Request message of an RPC.
// In this case, if the message name contains 'Request', the message is considered to be one.
func (m Message) IsRequest() bool {
	if strings.Contains(m.Name, "Request") {
		return true
	}
	return false
}

// IsResponse returns true if the message is a Response message of an RPC.
// In this case, if the message name contains 'Response', the message is considered to be one.
func (m Message) IsResponse() bool {
	if strings.Contains(m.Name, "Response") {
		return true
	}
	return false
}

// IsModel returns true if the current message is neither a request nor a response
func (m Message) IsModel() bool {
	if !m.IsRequest() && !m.IsResponse() {
		return true
	}
	return false
}

// FieldList returns the message fields as a list string which can be used as param, e.g. 'name string, foo int, bar bool'
func (m Message) FieldList(ctx Context) string {
	var list []string
	for _, field := range m.Fields {
		if ctx.IsEnum(field.Type) {
			list = append(list, fmt.Sprintf("%s %s", field.Name, field.Type))
		} else {
			list = append(list, fmt.Sprintf("%s %s", field.Name, field.ResolveType()))
		}
	}
	return strings.Join(list, ", ")
}

// FieldNames returns all field names without types
func (m Message) FieldNames() string {
	var list []string
	for _, field := range m.Fields {
		list = append(list, field.Name)
	}
	return strings.Join(list, ", ")
}

// FieldNamesPrefixed returns all field names without types but with a prefix added.
func (m Message) FieldNamesPrefixed(prefix string) string {
	var list []string
	for _, field := range m.Fields {
		list = append(list, prefix+strcase.ToCamel(field.Name))
	}
	return strings.Join(list, ", ")
}

// FieldStructInit returns a struct init style string: Name: name
func (m Message) FieldStructInit() string {
	var list []string
	for _, field := range m.Fields {
		list = append(list, fmt.Sprintf("%s: %s", strcase.ToCamel(field.Name), field.Name))
	}
	return strings.Join(list, ", ")
}

// FieldStructDeclare returns all fields of the message and their types, ready to use for a struct declaration
func (m Message) FieldStructDeclare(ctx Context) string {
	var list []string
	for _, field := range m.Fields {
		if ctx.IsEnum(field.Type) {
			list = append(list, fmt.Sprintf("%s %s", strcase.ToCamel(field.Name), field.Type))
		} else {
			list = append(list, fmt.Sprintf("%s %s", strcase.ToCamel(field.Name), field.ResolveType()))
		}
	}
	return strings.Join(list, "\n ")
}

// MessageField defines the field of a protobuf message
type MessageField struct {
	Name     string
	Type     string
	Order    int
	Repeated bool
}

func (f *MessageField) ResolveType() string {
	var prefix = ""
	if f.Repeated {
		prefix += "[]"
	}
	if !types.IsBuiltinTypeString(f.Type) {
		prefix += "*"
	}

	return prefix + f.Type
}

// Enum holds the enum definition from a protobuf file
type Enum struct {
	Name     string
	Comments []string
	Fields   []EnumField
}

// EnumField defines a single enum field from a protobuf file.
type EnumField struct {
	Name  string
	Order int
}
