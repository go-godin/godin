package godin

import (
	"os"
	"strings"

	"github.com/emicklei/proto"
)

func Parse(protoPath string) (*Context, error) {
	reader, err := os.Open(protoPath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	def, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	ctx := &Context{}

	proto.Walk(
		def,
		proto.WithService(parseService(ctx)),
		proto.WithEnum(parseEnum(ctx)),
		proto.WithMessage(parseMessage(ctx)),
		WithPackage(parsePackage(ctx)),
	)

	return ctx, nil
}

func WithPackage(apply func(*proto.Package)) proto.Handler {
	return func(v proto.Visitee) {
		if s, ok := v.(*proto.Package); ok {
			apply(s)
		}
	}
}

func parsePackage(ctx *Context) func(p *proto.Package) {
	return func(p *proto.Package) {
		ctx.Package = p.Name
	}
}

func parseMessage(ctx *Context) func(m *proto.Message) {
	return func(m *proto.Message) {
		msg := Message{
			Name: m.Name,
		}

		if m.Comment != nil {
			for _, line := range m.Comment.Lines {
				line = strings.TrimLeft(line, "// ")
				line = strings.TrimRight(line, "// ")
				msg.Comments = append(msg.Comments, line)
			}
		}

		for _, elem := range m.Elements {
			switch elem.(type) {
			default:
				continue
			case *proto.Comment:
				continue
			case *proto.NormalField:
				f := elem.(*proto.NormalField)
				msg.Fields = append(msg.Fields, MessageField{
					Name:     f.Name,
					Type:     f.Type,
					Order:    f.Sequence,
					Repeated: f.Repeated,
				})
			}
		}
		ctx.Messages = append(ctx.Messages, msg)
	}
}

func parseEnum(ctx *Context) func(e *proto.Enum) {
	return func(e *proto.Enum) {
		enum := Enum{
			Name: e.Name,
		}

		if e.Comment != nil {
			for _, line := range e.Comment.Lines {
				line = strings.TrimLeft(line, "// ")
				line = strings.TrimRight(line, "// ")
				enum.Comments = append(enum.Comments, line)
			}
		}

		for _, elem := range e.Elements {
			switch elem.(type) {
			default:
				continue
			case *proto.Comment:
				continue
			case *proto.EnumField:
				f := elem.(*proto.EnumField)
				field := EnumField{
					Name:  f.Name,
					Order: f.Integer,
				}

				enum.Fields = append(enum.Fields, field)
			}
		}
		ctx.Enums = append(ctx.Enums, enum)
	}
}

func parseService(ctx *Context) func(s *proto.Service) {
	return func(s *proto.Service) {
		svc := Service{
			Name: s.Name,
		}

		if s.Comment != nil {
			for _, line := range s.Comment.Lines {
				line = strings.TrimLeft(line, "// ")
				line = strings.TrimRight(line, "// ")
				svc.Comments = append(svc.Comments, line)
			}
		}

		for _, elem := range s.Elements {
			switch elem.(type) {
			default:
				continue
			case *proto.Comment:
				continue
			case *proto.RPC:
				meth := elem.(*proto.RPC)
				rpc := RPC{
					Name:         meth.Name,
					RequestType:  meth.RequestType,
					ResponseType: meth.ReturnsType,
				}

				if meth.Comment != nil {
					for _, line := range meth.Comment.Lines {
						line = strings.TrimLeft(line, "// ")
						line = strings.TrimRight(line, "// ")
						rpc.Comments = append(rpc.Comments, line)
					}
				}
				svc.RPCs = append(svc.RPCs, rpc)
			}
		}
		ctx.Services = append(ctx.Services, svc)
	}
}
