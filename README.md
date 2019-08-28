# Godin - An opinionated go-kit generator

An awesome toolkit to support go developers (me) when writing microservices which leverage go-kit.
Godin takes care of generating as much code as possible while being opinionated.
Generally speaking, Godin will provide templates within your project, 
specifically for your project. It will then always use these templates 
to generate code, giving the developer a lot of freedom. 

# Getting started
> TODO

# Commands
### `godin init` / `godin i`
Initializes a godin project in the current directory. It requires a few options to  be set:
+ `--namespace, -n` The project's namespace (use the kubernetes namespace for convenience), lowercase
+ `--service, -s` The service name which this project implements, lowercase
+ `--module, -m` The module which of the project. Godin will automatically initialize the module for you.
+ `--protobuf-module, -p` The module in which the generated protobuf stubs for this service are located.

Example: 
```bash
godin init --namespace godin --service ticket --module github.com/go-godin/godin/examples/ticket --protobuf-module github.com/go-godin/ticket-service/api
```

This will leave you with an empty project. You can then proceed and add modules to generate code for you.

### `godin add <module>` / `godin a <module>`
Install a module into the current project. A list of available modules is below.
When installing a module, two things happen:
+ The modules configuration is saved into the `godin.yaml`, which indicates that the module is enabled
+ The templates of the module are copied into the project (default folder is `templates`). This allows the developer
to be able to modify them to fit to fit the use-case you're solving.

**Note:** A module can only be added once!

Example:
```bash
godin add transport.grpc.server
```
  
### `godin generate` / `godin g`
This command is probably the most frequently used. It generates all module's templates based on the configuration.
In order to have access to the protobuf types, you need to pass the source proto-file which defines the API of the
service you are implementing (`--protobuf, -p` option).

Godin will render all templates in the project's templates folder into their respective targets (see module info).
If a template is missing, a warning is logged. The generate command does not restore any missing files, it will
only operate on what you  have in the project's templates.

The protobuf path can also be set via the `PROTOBUF_FILE` environment variable. That way you could
use a make-target to easily call the generate command like `make generate`.

Example:
```bash
godin generate -p /some/where/over/the/rainbow.proto
```

### `godin update-templates` / `godin u`
Updates the project templates with the templates of the currently installed godin version.
It will only update templates of enabled modules. It will *not overwrite* existing templates as
godin cannot figure out if you've made changes to it. 

In order to update all templates, regardless whether they already exist, you have two options:
* Simply remove the templates folder and let godin recreate it
* Use the `--force, -f` option (recommended) which will print a warning but overwrite existing templates

**TODO**: partially update the templates by module

# Why?
After a while of writing microservices the boilerplate code to write go-kit services gets annoying.
Godin is my attempt of solving that problem leveraging go-kit and trying to obey the clean architecture principles.

**Yes, I get that, but why yet another generator?** 

Well, first of all - it's fun. Secondly I really wasn't content with existing (and my previous) solutions.
The reason being that - sooner or later - you have to force a developer down a certain path.
There is nothing wrong with having a standard way of doing things. But when writing a code-generator
you need to think about **a lot** of edge cases in order to make it usable.

All previous attempts were abandoned because I was hitting such an edge case and was forced to either
ditch the code generator by editing generated code or modify it first before continuing to solve my actual problem.

This led me to the following premises under which this version of godin was developed.

## Premises
1. Shared models are always specified using Protobuf
2. The service API is always specified using gRPC services
3. Every gRPC service represents it's own subdomain (DDD)
4. Godin will not restrict the developer with it's generated files 

## The role of templates
Have you ever written a code generator before? I did, quite a few recently actually (and you should too). And I've always found it very
tempting to just ship one large binary which simply included all available templates ([like I did in a previous godin version](https://github.com/lukasjarosch/godin)). 

If you start going down that path, you must be confident, that your templates are covering **all cases**. Because as soon as a developer
hits an uncovered edge-case, he is forced to stop using your tool and first make the necessary changes to the tool. Although that
might work, I don't really see why one would try to provide centralized templates only.

**Don't force the developer to use your templates**  
So, what if you instead provided all the necessary templates inside the project and let the developer modify them to his needs.
That's way more convenient. Because most of the time, you don't even need to adjust templates (well, depends on your templates :)).
Even better, *IF* a developer hits an edge-case, he can just keep editing his current project and adjust it to his needs.
`godin generate` will continue working => less frustration.

I personally think that this is a good path to go. You can ship your services, based on protobuf, including the
templates which lead to most of the code. Everything is under version control, great!

**But what if upstream changes?**  
Consider this: You have custom template modifications on a microservice project. Then you upgrade (git pull) godin (or w/e tool) to the 
newest version. Everything that needs to be done at this point is to compare the local (project) templates with the
upstream templates. Essentially, this is a `git diff` taking place. If a local modification is found, simply force the developer
to properly merge the files.   
That way each project of yours can still have it's own template modifications while keeping godin up to date.

**Go only?**  
Heck, no! This might be completely written in Go, but by no means can godin be used with Go applications only.
As long as the premises apply, you can use any programming language. All you need to do is write templates for it
and maybe introduce some modification into godin.


# Modules
### `transport.grpc.server`
| Source | Target | Note | Overwrite |
|--------|--------|------|-----------|
| `transport/grpc/server.go.tmpl`     | `internal/transport/grpc/server.go`     |  Generates the gRPC server transport layer    | YES |

### `transport.grpc.client`
| Source | Target | Note | Overwrite |
|--------|--------|------|-----------|
| `transport/grpc/client.go.tmpl`     | `pkg/grpc/client.go`     |  Generates the gRPC client transport layer for downstream usage    | YES |

### `service.interface`
| Source | Target | Note | Overwrite |
|--------|--------|------|-----------|
| `service/interface.go.tmpl`     | `internal/<serviceName>/service.go`     |  The main service interface   | YES |
| `service/models.go.tmpl`     | `internal/<serviceName>/models.go`     |  Generates the protobuf messages (except request and response) as well as enums as a base for your domain-models.   | NO |

### `service.endpoints`
| Source | Target | Note | Overwrite |
|--------|--------|------|-----------|
| `endpoint/endpoints.go.tmpl`     | `internal/endpoint/endpoints.go`     |  The endpoint set   | YES |
| `endpoint/request_response.go.tmpl`     | `internal/endpoint/request_response.go`     |  The internally used request and response structs   | YES |


#### `transport.amqp.publisher`
#### `transport.amqp.subscriber`
### `store.mysql`
### `store.mongodb`
### `k8s.deployment`
### `k8s.service`
### `k8s.service`
### `project.dockerfile`
### `project.makefile`
### `project.readme`
### `service.middleware.logging`
### `service.middleware.authentication`
### `service.middleware.caching`
