# Flagd build spec

## This repository is managed by OpenFeature  


This module contains the core types that developers can use for interacting with [flagd](https://github.com/open-feature/flagd).

Internally flagd uses the connect protocol, meaning it is compatible with grpc interfaces. If your desired language has a supported plugin for generating connect stubs then it is reccomended to use these over grpc.

The package contains a single `Service`, describing the 5 core `rpcs` for feature flag evaluation (`ResolveBoolean`, `ResolveString`, `ResolveFloat`, `ResolveInt` and `ResolveObject`) each with their type specific requests and responses (`ResolveXXXRequest` and `ResolveXXXResponse`).
The final `rpc` on the `Service` is a streamed response called `EventStream`, this is used to pass internal events to the client, such as `configuration_change` or `provider_ready`.

## Build options

The core definitions are in the `schema.v1` package, and contains package name options for the following languages, as a result these can be excluded from build instructions:

- Go: schema/service/v1       
- Java:   dev.openfeature.flagd.grpc        
- C#:  OpenFeature.Flagd.Grpc        
