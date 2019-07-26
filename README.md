# Godin - opinionated go-kit toolkit 

> **Note** This is only a draft!

# What?
An awesome toolkit to support go developers (me) when writing microservices which leverage go-kit.
Godin takes care of generating as much code as possible while being opinionated.

# Why?
After a while of writing microservices the boilerplate code to write go-kit services gets annoying.
Godin is my attempt of solving that problem leveraging go-kit and trying to obey the clean architecture principles.

## Premises
1. Shared models are always specified using Protobuf
2. The service API is always specified using gRPC services
3. Every gRPC service represents it's own subdomain (DDD)

## What does it do exactly?
####  Firstly, what are the components of a go-kit service? 
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

