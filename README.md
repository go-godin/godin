# Godin - opinionated go-kit toolkit 

> **Note** This is only a draft!

# What?
An awesome toolkit to support go developers (me) when writing microservices which leverage go-kit.
Godin takes care of generating as much code as possible while being opinionated.
Generally speaking, Godin will provide templates within your project, 
specifically for your project. It will then always use these templates 
to generate code, giving the developer a lot of freedom. 

# Why?
After a while of writing microservices the boilerplate code to write go-kit services gets annoying.
Godin is my attempt of solving that problem leveraging go-kit and trying to obey the clean architecture principles.

## Premises
1. Shared models are always specified using Protobuf
2. The service API is always specified using gRPC services
3. Every gRPC service represents it's own subdomain (DDD)
4. Godin will not restrict the developer with it's generated files 

## What does it do exactly?
###  Firstly, what are the components of a go-kit service? 
In order to know which work godin aims on taking over for you, we need to define what a service (roughly) looks like:

* **transport layer**
  + grpc
    - server
    - client
    - encode_decode
  + amqp
    - publisher
    - subscriber
     - encode_decode
  + http
    - server
    - client
    - encode_decode
* **endpoint layer**
  + middleware
* **interface layer**
  + repositories
  + publishers
  + mappings (domain <-> DAO)
* **use-case / service layer**
  + use-case implementations
  + domain models
  + unittests

Considering the premises, Godin tries to take care of the **transport and endpoint layer**. 
It should also support the developer inside the **interface layer** in form of generating middleware boilerplates, repositories and all the other fancy things one might require in order to fulfil the use-cases.

In terms of *shared* and *hidden* models, Godin will take care of the *shared* models as they are directly specified via Protobuf. By generating a `service.go`, Godin establishes the contract between protobuf and the implementation (go-kit style).
That way the developer can focus on what's really important...business functionality.

### The role of templates
Have you ever written a code generator before? I did, quite a few in recent times tbh. And I've always found it very
tempting to just ship one large binary which simply included all available templates ([like I did in a previous godin version](https://github.com/lukasjarosch/godin)). 

If you start going down that path, you must be confident, that your templates are covering **all cases**. Because as soon as a developer
hits an uncovered edge-case, he is forced to stop using your tool and first make the necessary changes to the tool. Although that
might work, I don't really see why one would try to provide centralized templates only.

**Don't force the developer to use your templates**  
So, what if you instead provided all the necessary templates inside the project and let the developer modify them to his needs.
That's way more convenient. Because most of the time, you don't even need to adjust templates (well, depends on your templates :)).
Even better, *IF* a developer hits an edge-case, he can just keep editing his current project and adjust it to his needs.
`godin generate` will continue working => no more frustation.

I personally think that this is a very good path to go. You can ship your services, based on protobuf, including the
templates which lead to most of the code. Everything is under version control, great!

**But what if upstream changes?**  
Consider this: You have custom tempalte modifications on a microservice project. Then you upgrade (git pull) godin (or w/e tool) to the 
newest version. Everything that needs to be done at this point is to compare the local (project) templates with the
upstream templates. Essentially, this is a `git diff` taking place. If a local modification is found, simply force the developer
to properly merge the files.   
That way each project of yours can still have it's own template modifications while keeping godin up to date.

**Go only?**  
Heck, no! This might be completely written in Go, but by no means can godin be used with Go applications only.
As long as the premises apply, you can use any programming language. All you need to do is write templates for it
and maybe introduce some modification into godin.

## Commands
### `godin init`
+ checks whether a `godin.yaml` is already present in the CWD. If so, init has been called already.
+ create `godin.json` in CWD
+ prompt user: 
  - `project_name`: The lowercase name of the project (usually the service name)
  - `project_namespace`: The namespace is used for the kubernetes manifests
  - `template_dir`: That's the target directory where godin will generate the project-templates into
  - `proto_path`: Used as include-path for the **protoc** command. In godin.json, this is actually an array `proto_paths`.
  - `output_dir`: That's where all generated code will go, default is `./` 
+ create **template_dir** and ensure permissions
+ check all **proto_paths**, at least one must exist and be a directory. 
+ find all `*.proto` files in the proto_path and save them in `proto_files`

  
### `godin generate`
Calling the *generate* command will trigger godin to call protoc and render all project templates into
the configured target directory. 

+ read the `godin.yaml` and verify that all prerequisites are fulfilled
+ (temporary solution) call `protoc` directly to generate all files
+ the `protoc-gen-gotemplates` protoc plugin is invoked. It's passed the `template_dir` as well as the path to `godin.yaml`
+ assemble the template context based based on the parsed protobuf file(s)
+ enrich the context with the configuration data.
+ render all templates to the desired `outpud_dir`. The folder structure is preserved.

> Note: You will call `godin generate` quite often. Maybe spend an alias to godin? Something like `gg`? 

### `godin add <module>`
Add a module to the project templates. Upon calling `godin generate`, the module templates are rendered.
More about the modules, see below.

### `godin remove <module>`
### `godin update [module]`
Update is responsible of keeping the project-specific template up to date with upstream godin.  
Hereby you can either update all modules or specify which module to update.

### `godin help`
Prints the full help for godin which (should ideally) have all the info and more which is in this README.

### `godin version`
Prints the currently installed godin version. Godin currently does not have a remote backend server
for version checks, so it's your task to keep it up to date manually.


## Modules
### `transport.grpc.server`
### `transport.grpc.client`
### `transport.grpc.client_library`
### `transport.http.server`
### `transport.http.client`
### `transport.http.debug_server`
### `transport.client.client_library`
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
