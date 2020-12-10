package log

import (
	"math"
	"time"
)

// FieldType represent D value type
type FieldType int32

// DType enum
const (
	UnknownType FieldType = iota
	StringType
	IntTpye
	Int64Type
	UintType
	Uint64Type
	Float32Type
	Float64Type
	DurationType
)

// Field is for encoder
type Field struct {
	Key   string
	Value interface{}
	Type  FieldType
}

//KVString ...
func KVString(k, v string) D {
	return D{
		Key:   k,
		Value: v,
		Type:  StringType,
	}
}

// KVInt64 construct D with int64 value.
func KVInt64(key string, value int64) D {
	return D{Key: key, Type: Int64Type, Value: value}
}

// KVUint construct Field with uint value.
func KVUint(key string, value uint) D {
	return D{Key: key, Type: UintType, Value: int64(value)}
}

// KVUint64 construct Field with uint64 value.
func KVUint64(key string, value uint64) D {
	return D{Key: key, Type: Uint64Type, Value: int64(value)}
}

// KVFloat32 construct Field with float32 value.
func KVFloat32(key string, value float32) D {
	return D{Key: key, Type: Float32Type, Value: int64(math.Float32bits(value))}
}

// KVFloat64 construct Field with float64 value.
func KVFloat64(key string, value float64) D {
	return D{Key: key, Type: Float64Type, Value: int64(math.Float64bits(value))}
}

// KVDuration construct Field with Duration value.
func KVDuration(key string, value time.Duration) D {
	return D{Key: key, Type: DurationType, Value: int64(value)}
}

// KV return a log kv for logging field.
// NOTE: use KV{type name} can avoid object alloc and get better performance. []~(￣▽￣)~*干杯
func KV(key string, value interface{}) D {
	return D{Key: key, Value: value}
}
