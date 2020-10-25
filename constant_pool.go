package bytecode

import (
	"fmt"
	"reflect"
)

type ConstantPool struct {
	buffer    *Buffer
	lastIndex uint16
	cache     map[string]uint16
}

func CreateConstantInfo() *ConstantPool {
	return &ConstantPool{
		buffer:    CreateBuffer(),
		lastIndex: 0,
		cache:     make(map[string]uint16, 0),
	}
}

func (constPool *ConstantPool) AddAuto(data interface{}) uint16 {
	switch value := data.(type) {
	case int8:
		return constPool.AddNumber(int(value))
	case int16:
		return constPool.AddNumber(int(value))
	case int32:
		return constPool.AddNumber(int(value))
	case int:
		return constPool.AddNumber(value)
	case int64:
		return constPool.AddLong(value)
	case float32:
		return constPool.AddFloat(value)
	case float64:
		return constPool.AddDouble(value)
	case string:
		return constPool.AddString(value)
	}

	panic("Can't add unknown type to const pool: " + reflect.TypeOf(data).Name())
}

func (constPool *ConstantPool) AddNumber(value int) uint16 {
	return constPool.AddConst(Constant_Integer, int32(value))
}

func (constPool *ConstantPool) AddFloat(value float32) uint16 {
	return constPool.AddConst(Constant_Float, value)
}

func (constPool *ConstantPool) AddDouble(value float64) uint16 {
	return constPool.AddConst(Constant_Double, value)
}

func (constPool *ConstantPool) AddLong(value int64) uint16 {
	return constPool.AddConst(Constant_Long, value)
}

func (constPool *ConstantPool) AddUtf8(value string) uint16 {
	return constPool.AddConst(Constant_Utf8, uint16(len(value)), []byte(value))
}

func (constPool *ConstantPool) AddString(value string) uint16 {
	return constPool.AddConst(Constant_String, constPool.AddUtf8(value))
}

func (constPool *ConstantPool) AddClass(name string) uint16 {
	return constPool.AddConst(Constant_Class, constPool.AddUtf8(name))
}

func (constPool *ConstantPool) AddNameAndType(name string, descriptor string) uint16 {
	return constPool.AddConst(Constant_NameAndType, constPool.AddUtf8(name), constPool.AddUtf8(descriptor))
}

func (constPool *ConstantPool) AddFieldRef(class string, name string, descriptor string) uint16 {
	return constPool.AddConst(Constant_Fieldref, constPool.AddClass(class), constPool.AddNameAndType(name, descriptor))
}

func (constPool *ConstantPool) AddMethodRef(class string, name string, descriptor string) uint16 {
	return constPool.AddConst(Constant_Methodref, constPool.AddClass(class), constPool.AddNameAndType(name, descriptor))
}

func (constPool *ConstantPool) AddConst(tag uint8, values ...interface{}) uint16 {
	bytes := Parse(tag)
	for _, value := range values {
		bytes = append(bytes, Parse(value)...)
	}

	hex := fmt.Sprintf("%x", bytes)
	if index, found := constPool.cache[hex]; found {
		return index
	}

	constPool.lastIndex++
	constPool.cache[hex] = constPool.lastIndex
	constPool.buffer.Write(bytes)

	if tag == Constant_Long || tag == Constant_Double {
		constPool.lastIndex++ // 8-byte variables takes 2 indexes in Const Pool, but it must return first index
		return constPool.lastIndex - 1
	}

	return constPool.lastIndex
}
