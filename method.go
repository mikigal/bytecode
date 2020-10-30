package bytecode

import (
	"strconv"
	"strings"
)

type Method struct {
	class      *Class
	Name       string
	Descriptor string
	LocalVars  map[uint8]JavaType

	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	attributesCount uint16
	attributes      *Buffer

	// Code attribute
	codeAttributeNameIndex uint16
	codeAttributeLength    uint32
	maxStack               uint16
	maxLocals              uint16
	codeBytes              *Buffer
	exceptionTableLength   uint16
	codeAttributesCount    uint16
}

func (method *Method) GetReference() uint16 {
	return method.class.ConstPool.AddMethodRef(method.class.ClassName, method.Name, method.Descriptor)
}

func (method *Method) AddFieldInstruction(instruction Instruction, class string, name string, descriptor string) {
	if instruction != Getstatic && instruction != Putstatic && instruction != Getfield && instruction != Putfield {
		panic("AddFieldInstruction can be used only for instructions that need FieldRef as parameter")
	}

	method.AddInstruction(instruction, method.class.ConstPool.AddFieldRef(class, name, descriptor))
}

func (method *Method) AddInvokeInstruction(instruction Instruction, class string, name string, descriptor string) {
	if !strings.HasPrefix(instruction.Name, "Invoke") {
		panic("AddInvokeInstruction can be used only for invoke instructions")
	}

	method.AddInstruction(instruction, method.class.ConstPool.AddMethodRef(class, name, descriptor))
}

func (method *Method) AddInstruction(instruction Instruction, parameters ...interface{}) {
	method.codeBytes.Write(instruction.Value)

	parametersLength := 0
	for i := 0; i < len(parameters); i++ {
		parametersLength += len(Parse(parameters[i]))
		method.codeBytes.Write(parameters[i])
	}

	if instruction.CreateLocalVar {
		method.maxLocals++

		// long and double take 2 indexes in locals
		if strings.HasPrefix(instruction.Name, "Dstore") || strings.HasPrefix(instruction.Name, "Lstore") {
			method.maxLocals++
		}
	}

	if int(instruction.ParameterLength) != parametersLength {
		panic("could not create instruction " + instruction.Name + " with parameters length: " + strconv.Itoa(parametersLength) + ". It require " + strconv.Itoa(int(instruction.ParameterLength)) + " bytes as parameters")
	}
}

func (method *Method) LoadLocalVar(index uint8) {
	varType, exists := method.LocalVars[index]
	if !exists {
		panic("there's no local var with index " + strconv.Itoa(int(index)))
	}

	if instruction, basic := varType.Load[int(index)]; basic {
		method.AddInstruction(instruction)
		return
	}

	method.AddInstruction(varType.Load[-1], index)
}

func (method *Method) AddLocalVar(data interface{}) uint8 {
	if data == nil {
		panic("could not create local var with null value")
	}
	javaType := AsJavaType(data)
	if constInstruction, ok := javaType.Const[data]; ok {
		method.AddInstruction(constInstruction)
	} else {
		pushedInt := false

		if javaType.Name == Type_Int.Name {
			asInt := -1
			if temp, ok := data.(int); ok {
				asInt = temp
			} else if temp, ok := data.(int32); ok {
				asInt = int(temp)
			}

			if asInt >= -128 && asInt <= 127 {
				pushedInt = true
				method.AddInstruction(Bipush, uint8(asInt))
			} else if asInt >= -32768 && asInt <= 32767 {
				method.AddInstruction(Sipush, int16(asInt))
				pushedInt = true
			}
		}

		index := method.class.ConstPool.AddAuto(data)

		if javaType.Wide && !pushedInt {
			method.AddInstruction(Ldc2_w, index)
		} else if index <= 255 && !pushedInt {
			method.AddInstruction(Ldc, uint8(index))
		} else if !pushedInt {
			method.AddInstruction(Ldc_w, index)
		}
	}

	if storeInstruction, ok := javaType.Store[int(method.maxLocals)]; ok {
		method.AddInstruction(storeInstruction)
	} else {
		method.AddInstruction(javaType.Store[-1], uint8(method.maxLocals))
	}

	varIndex := uint8(method.maxLocals - Ternary(javaType.Wide, uint16(2), uint16(1)).(uint16))
	method.LocalVars[varIndex] = javaType
	return varIndex
}

func (method *Method) toBytes() []byte {
	codeBuffer := CreateBuffer()
	codeBuffer.Write(method.maxStack)
	codeBuffer.Write(method.maxLocals)
	codeBuffer.Write(uint32(len(method.codeBytes.bytes)))
	codeBuffer.Write(method.codeBytes.bytes)
	codeBuffer.Write(method.exceptionTableLength)
	codeBuffer.Write(method.codeAttributesCount)
	method.codeAttributeLength = uint32(len(codeBuffer.bytes))

	buffer := CreateBuffer()
	buffer.Write(method.AccessFlags)
	buffer.Write(method.NameIndex)
	buffer.Write(method.DescriptorIndex)
	buffer.Write(method.attributesCount)
	buffer.Write(method.attributes.bytes)
	buffer.Write(method.codeAttributeNameIndex)
	buffer.Write(method.codeAttributeLength)
	buffer.Write(codeBuffer.bytes)

	return buffer.bytes
}
