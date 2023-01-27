# protoc-gen-zap-marshaler

> **Note:** This is a modified fork of https://github.com/kazegusuri/go-proto-zap-marshaler.
>
> - Updated to avoid deprecated dependency usage.
> - Removed non-secure marshaler.
> - Changed naming.
>
> If you're making additional changes to this plugin, I recommend testing using the `make test` target in this directory; tests in ./examples addume the protoc-generated artifacts are up to date.

A protoc plugin whitch generates `MarshalLogObject()` functions for each generated go structs to implement `zapcore.ObjectMarshaler` interface for [uber-go/zap](https://github.com/uber-go/zap). So you can use `zap.Object("proto", someProtoMessage)` to log contents of the proto message.

## Install

There is one protoc plugins to generate marshaler functions for zap:

- `protoc-gen-zap-marshaler`
  - which genearetes marshaler functions for all messages, but only logs fields enabled by proto option.

You can install this plugin with the following command:

```bash
$ go get github.com/indentinc/protoc-gen-zap-marshaler/protoc-gen-zap-marshaler
```

## Usage

### protoc-gen-zap-marshaler

This plugin generates marshaler functions as well as `protoc-gen-zap-marshaler`, but the funtions does not log anything as default. The marshaler only marshal fields enabled by [zap_marshaler](https://github.com/indentinc/protoc-gen-zap-marshaler/blob/master/zap_marshaler.proto) option.

To enable field option for message fields, you need to define proto like this:

```proto
message SimpleMessage {
  string string_value = 1 [(indent.zap.Include) = true];
  bool bool_value = 2 [(indent.zap.Include) = true];
}
```

To generate marshaler functions from proto, use `--zap-marshaler_out` with `protoc` command. That runs `protoc-gen-zap-marshaler` internally and then generates `*.zap.go` files.

```
$ protoc --zap-marshaler-secure_out=. path/to/example.proto
```

See usage in `Makefile`'s `regenerate/...` targets.

#### Example

For this proto without field option, the secure plugin generates a go function:

```proto
syntax = "proto3";
package example;

message SimpleMessage {
  string string_value = 1;
  bool bool_value = 2;
}
```

```go
func (m *SimpleMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	return nil
}
```

The generated function does nothing.

By changing the proto with field option like this, the secure plugin generates a go function:

```proto
syntax = "proto3";
package example;

import "indentinc/protoc-gen-zap-marshaler/zap.proto";

message SimpleMessage {
  string string_value = 1 [(indent.zap.Include) = true];
  bool bool_value = 2 [(indent.zap.Include) = true];
}
```

```go
func (m *SimpleMessage) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	keyName = "string_value" // field string_value = 1
	enc.AddString(keyName, m.StringValue)

	keyName = "bool_value" // field bool_value = 2
	enc.AddBool(keyName, m.BoolValue)

	return nil
}
```
