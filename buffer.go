package bytecode

import (
	"math"
	"reflect"
)

type Buffer struct {
	bytes []byte
}

func ParseInt8(value int8) []byte {
	return []byte{byte(value)}
}

func ParseUint8(value uint8) []byte {
	return ParseInt8(int8(value))
}

func ParseInt16(value int16) []byte {
	return []byte{byte(value >> 8), byte(value)}
}

func ParseUint16(value uint16) []byte {
	return ParseInt16(int16(value))
}

func ParseInt32(value int32) []byte {
	return []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

func ParseUint32(value uint32) []byte {
	return ParseInt32(int32(value))
}

func ParseInt64(value int64) []byte {
	return []byte{byte(value >> 56), byte(value >> 48), byte(value >> 40), byte(value >> 32),
		byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

func ParseUint64(value uint64) []byte {
	return ParseInt64(int64(value))
}

func ParseFloat32(value float32) []byte {
	bits := math.Float32bits(value)
	return []byte{byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)}
}
func ParseFloat64(value float64) []byte {
	bits := math.Float64bits(value)
	return []byte{byte(bits >> 56), byte(bits >> 48), byte(bits >> 40), byte(bits >> 32),
		byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)}
}

func ParseBool(value bool) []byte {
	return []byte{byte(Ternary(value, 1, 0).(int))}
}

func (buffer *Buffer) WriteUint8(value uint8) {
	buffer.WriteByteArray(ParseUint8(value))
}

func (buffer *Buffer) WriteInt8(value int8) {
	buffer.Write(ParseInt8(value))
}

func (buffer *Buffer) WriteUint16(value uint16) {
	buffer.Write(ParseUint16(value))
}

func (buffer *Buffer) WriteInt16(value int16) {
	buffer.Write(ParseInt16(value))
}

func (buffer *Buffer) WriteUint32(value uint32) {
	buffer.Write(ParseUint32(value))
}

func (buffer *Buffer) WriteInt32(value int32) {
	buffer.Write(ParseInt32(value))
}

func (buffer *Buffer) WriteUint64(value uint64) {
	buffer.Write(ParseUint64(value))
}

func (buffer *Buffer) WriteInt64(value int64) {
	buffer.Write(ParseInt64(value))
}

func (buffer *Buffer) WriteFloat32(value float32) {
	buffer.Write(ParseFloat32(value))
}

func (buffer *Buffer) WriteFloat64(value float64) {
	buffer.Write(ParseFloat64(value))
}

func (buffer *Buffer) WriteBool(value bool) {
	buffer.Write(ParseBool(value))
}

func (buffer *Buffer) WriteByteArray(value []byte) {
	for i := 0; i < len(value); i++ {
		buffer.writeRaw(value[i])
	}
}

func (buffer *Buffer) Write(data interface{}) {
	switch value := data.(type) {
	case uint8:
		buffer.WriteUint8(value)
		break
	case int8:
		buffer.WriteInt8(value)
		break
	case []uint8:
		for _, v := range value {
			buffer.WriteUint8(v)
		}
		break
	case []int8:
		for _, v := range value {
			buffer.WriteInt8(v)
		}
		break
	case uint16:
		buffer.WriteUint16(value)
		break
	case int16:
		buffer.WriteInt16(value)
		break
	case []uint16:
		for _, v := range value {
			buffer.WriteUint16(v)
		}
		break
	case []int16:
		for _, v := range value {
			buffer.WriteInt16(v)
		}
		break
	case uint32:
		buffer.WriteUint32(value)
		break
	case int32:
		buffer.WriteInt32(value)
		break
	case []uint32:
		for _, v := range value {
			buffer.WriteUint32(v)
		}
		break
	case []int32:
		for _, v := range value {
			buffer.WriteInt32(v)
		}
		break
	case uint64:
		buffer.WriteUint64(value)
		break
	case int64:
		buffer.WriteInt64(value)
		break
	case []uint64:
		for _, v := range value {
			buffer.WriteUint64(v)
		}
		break
	case []int64:
		for _, v := range value {
			buffer.WriteInt64(v)
		}
		break
	case float32:
		buffer.WriteFloat32(value)
		break
	case []float32:
		for _, v := range value {
			buffer.WriteFloat32(v)
		}
		break
	case float64:
		buffer.WriteFloat64(value)
		break
	case []float64:
		for _, v := range value {
			buffer.WriteFloat64(v)
		}
		break
	case bool:
		buffer.WriteBool(value)
		break
	case []bool:
		for _, v := range value {
			buffer.WriteBool(v)
		}
		break
	case *Buffer:
		buffer.WriteByteArray(value.bytes)
		break
	default:
		panic("can not write type: " + reflect.TypeOf(data).Name())
	}
}

func Parse(data interface{}) []byte {
	switch value := data.(type) {
	case uint8:
		return ParseUint8(value)
	case int8:
		return ParseInt8(value)
	case uint16:
		return ParseUint16(value)
	case int16:
		return ParseInt16(value)
	case uint32:
		return ParseUint32(value)
	case int32:
		return ParseInt32(value)
	case uint64:
		return ParseUint64(value)
	case int64:
		return ParseInt64(value)
	case float32:
		return ParseFloat32(value)
	case float64:
		return ParseFloat64(value)
	case bool:
		return ParseBool(value)
	case []byte:
		return value
	default:
		panic("can not parse type: " + reflect.TypeOf(data).Name())
	}
}

func (buffer *Buffer) writeRaw(value byte) {
	buffer.bytes = append(buffer.bytes, value)
}

func CreateBuffer() *Buffer {
	return &Buffer{
		bytes: make([]byte, 0),
	}
}
