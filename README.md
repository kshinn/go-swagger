Swagger 2.0
===========

[![Build Status](https://travis-ci.org/casualjim/go-swagger.svg?branch=master)](https://travis-ci.org/casualjim/go-swagger)[![Coverage Status](https://coveralls.io/repos/casualjim/go-swagger/badge.svg?branch=master)](https://coveralls.io/r/casualjim/go-swagger?branch=master)[![GoDoc](https://godoc.org/github.com/casualjim/go-swagger?status.svg)](http://godoc.org/github.com/casualjim/go-swagger)[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/swagger-api/swagger-spec/master/LICENSE)

This API is not stable yet, when it is stable it will be distributed over gopkg.in

Contains an implementation of Swagger 2.0. It knows how to serialize and deserialize swagger specifications.

Swagger is a simple yet powerful representation of your RESTful API.  
With the largest ecosystem of API tooling on the planet, thousands of developers are supporting Swagger in almost every modern programming language and deployment environment.

With a Swagger-enabled API, you get interactive documentation, client SDK generation and discoverability. We created Swagger to help fulfill the promise of APIs.

Swagger helps companies like Apigee, Getty Images, Intuit, LivingSocial, McKesson, Microsoft, Morningstar, and PayPal build the best possible services with RESTful APIs.Now in version 2.0, Swagger is more enabling than ever. And it's 100% open source software.

Docs
----

http://godoc.org/github.com/casualjim/go-swagger

Install:

		go get -u github.com/casualjim/go-swagger/cmd/swagger

The implementation also provides a number of command line tools to help working with swagger.

Currently there is a spec validator tool:

		swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json

You can also serve a swagger document with the swagger UI

		swagger ui ./swagger.json

To generate a server for a swagger spec document:

		swagger generate server [-f ./swagger.json] [--principal [principal-name]] [--with-ui]

There are several other sub commands available for the generate command

Sub command | Description
------------|----------------------------------------------------------------------------------
operation   | generates one or more operations specified in the swagger definition
model       | generates model files for one or more models specified in the swagger definition
support     | generates the api builder and the main method
server      | generates an entire server application

Design
------

For now what exists of documentation on how all the pieces fit together, is described in this [doc](design.md)


What's inside?
--------------

For a V1 I want to have this feature set completed:

-	[x] An object model that serializes to swagger yaml or json
-	[x] A tool to work with swagger:
	-	[x] validate a swagger spec document:
    -	[x] validate against jsonschema
    -	[ ] validate extra rules outlined [here](https://github.com/apigee-127/swagger-tools/blob/master/docs/Swagger_Validation.md)
      - [ ] :boom: definition can't declare a property that's already defined by one of its ancestors (Error)
      - [ ] :boom: definition's ancestor can't be a descendant of the same model (Error)
      - [x] :boom: each api path should be non-verbatim (account for path param names) unique per method (Error)
      - [ ] :warning: each security reference should contain only unique scopes (Warning)
      - [ ] :warning: each security scope in a security definition should be unique (Warning)
      - [x] :boom: each path parameter should correspond to a parameter placeholder and vice versa (Error)
      - [ ] :warning: each referencable definition must have references (Warning)
      - [x] :boom: each definition property listed in the required array must be defined in the properties of the model (Error)
      - [x] :boom: each parameter should have a unique `name` and `type` combination (Error)
      - [x] :boom: each operation should have only 1 parameter of type body (Error)
      - [ ] :boom: each reference must point to a valid object (Error)
      - [ ] :boom: every default value that is specified must validate against the schema for that property (Error)
      - [x] :boom: items property is required for all schemas/definitions of type `array` (Error)
	-	[x] serve swagger UI for any swagger spec file
  - code generation
    -	[x] generate api based on swagger spec
    -	[ ] generate go client from a swagger spec
    -	[ ] generate "sensible" random data based on swagger spec
    -	[ ] generate tests based on swagger spec for client
    -	[ ] generate tests based on swagger spec for server
-	[x] Middlewares:
	-	[x] serve spec
	-	[x] routing
	-	[x] validation
	-	[x] additional validation through an interface
	-	[x] authorization
		-	[x] basic auth
		-	[x] api key auth
	-	[x] swagger docs UI
-	[x] Typed JSON Schema implementation
	-	[x] JSON Pointer that knows about structs
	-	[x] JSON Reference that knows about structs
	-	[x] Passes current json schema test suite
-	[x] extended string formats
	-	[x] uuid, uuid3, uuid4, uuid5
	-	[x] email
	-	[x] uri (absolute)
	-	[x] hostname
	-	[x] ipv4
	-	[x] ipv6
	-	[x] credit card
	-	[x] isbn, isbn10, isbn13
	-	[x] social security number
	-	[x] hexcolor
	-	[x] rgbcolor
	-	[x] date
	-	[x] date-time
	-	[x] duration
	-	[x] custom string formats

Later
-----

After the v1 implementation extra transports are on the roadmap

- Formats:
	- [ ] custom serializer for XML to support namespaces and prefixes
	- [ ] optimized serializer for JSON
	- [ ] optimized serializer for YAML
- Middlewares:
	- [ ] swagger editor
  - [ ] authorization:
		-	[ ] oauth2
			-	[ ] implicit
			-	[ ] access code
			-	[ ] password
			-	[ ] application
-	Tools:
	-	[ ] generate spec document based on the code
	-	[ ] watch swagger spec file and regenerate when modified
-	Transports:
	-	[ ] swagger socket (swagger over tcp sockets)
	-	[ ] swagger websocket (swagger over websockets)
	- [ ] swagger proxy (assemble several backend apis into a single swagger spec and route the requests)
	- [ ] swagger discovery (repository for swagger specifications)
