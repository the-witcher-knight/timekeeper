package tracing

import (
	"go.uber.org/zap"
)

type fieldType uint

const (
	fieldTypeString fieldType = iota
	fieldTypeInt
	fieldTypeBytes
)

// Field is struct contains log field
type Field struct {
	Key       string
	Value     interface{}
	FieldType fieldType
}

func String(key string, value string) Field {
	return Field{
		Key:       key,
		Value:     value,
		FieldType: fieldTypeString,
	}
}

func Bytes(key string, value []byte) Field {
	return Field{
		Key:       key,
		Value:     value,
		FieldType: fieldTypeBytes,
	}
}

func Int(key string, value int) Field {
	return Field{
		Key:       key,
		Value:     value,
		FieldType: fieldTypeInt,
	}
}

func toZapField(f Field) zap.Field {
	switch f.FieldType {
	case fieldTypeInt:
		return zap.Int(f.Key, f.Value.(int))
	case fieldTypeString:
		return zap.String(f.Key, f.Value.(string))
	case fieldTypeBytes:
		return zap.Binary(f.Key, f.Value.([]byte))
	default:
		return zap.Any(f.Key, f.Value)
	}
}

func toZapFieldSlice(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		zapFields[i] = toZapField(field)
	}

	return zapFields
}
