// Code generated by protoc-gen-zap-marshaler. DO NOT EDIT.
// Generated conservatively: no ObjectMarshaler implementations for messages
// without any logged fields.

package auditv1

import (
	fmt "fmt"
	zapcore "go.uber.org/zap/zapcore"
)

func (m *Event) MarshalLogObject(oe zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "Event" // field event = 1
	oe.AddString(keyName, m.Event)

	keyName = "Reason" // field reason = 3
	oe.AddString(keyName, m.Reason)

	keyName = "Timestamp" // field timestamp = 5
	t := m.Timestamp.AsTime()
	oe.AddTime(keyName, t)

	return nil
}

func (m *Resource) MarshalLogObject(oe zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "Id" // field id = 1
	oe.AddString(keyName, m.Id)

	keyName = "DisplayName" // field displayName = 2
	oe.AddString(keyName, m.DisplayName)

	keyName = "Kind" // field kind = 4
	oe.AddString(keyName, m.Kind)

	keyName = "Labels" // field labels = 10
	oe.AddObject(keyName, zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
		for mk, mv := range m.Labels {
			key := fmt.Sprint(mk)
			oe.AddString(key, mv)
		}
		return nil
	}))

	return nil
}