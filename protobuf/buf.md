# Flagd build spec

### Managed by [OpenFeature](https://github.com/open-feature). 

This repository contains grpc service definitions that developers can use for interacting with [flagd](https://github.com/open-feature/flagd).

Given below are the modules exposed through this repository,

- [Flag evaluation](#Flag-evaluation)
- [Flag sync](#Flag-sync)

## Flag evaluation

The module `schema.v1` contains a single `Service`, describing the 5 core `rpcs` for feature flag evaluation (`ResolveBoolean`,
`ResolveString`, `ResolveFloat`, `ResolveInt` and `ResolveObject`), each with their type specific request and
response objects(`ResolveXXXRequest` and `ResolveXXXResponse`). The final `rpc` on the `Service` is a streamed response
named `EventStream`, this is used to pass internal events to the client, such as `configuration_change` and `provider_ready`.

Internally flagd uses the connect protocol, meaning it is compatible with grpc interfaces. If your desired language has 
a supported plugin for generating connect stubs then it is recommended to use these over grpc.

## Flag sync

The module `sync.v1` is a grpc server streaming service definition to provide flagd with feature flag configurations.
This service expose a single method `SyncFlags`. Flagd acts as the client and initiate the streaming with `SyncFlagsRequest`.

The server implementation will then stream feature flag configurations through `SyncFlagsResponse`. The response contains
`SyncState` which can be utilized to provide flagd with flexible configuration updates.

## Code generate

Easiest way to generate grpc code for your language of choice is using provided [buf](https://buf.build/) templates.
For example, with required binaries in your path, java code generation can be done using 
`buf generate --template buf.gen.java.yaml` command.

