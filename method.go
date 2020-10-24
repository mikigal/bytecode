package bytecode

import (
	"strconv"
	"strings"
)

type Method struct {
	bytecode   *Bytecode
	Name       string
	Descriptor string

	AccessFlags     uint16
	NameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16

	// Code attribute
	codeAttributeNameIndex uint16
	codeAttributeLength    uint32
	maxStack               uint16
	maxLocals              uint16
	codeBytes            *Buffer
	exceptionTableLength uint16
	codeAttributesCount  uint16
}

func (method *Method) GetReference() uint16 {
	return method.bytecode.ConstPool.AddMethodRef(method.bytecode.ClassName, method.Name, method.Descriptor)
}

func (method *Method) AddFieldInstruction(instruction Instruction, class string, name string, descriptor string) {
	if instruction != Getstatic && instruction != Putstatic && instruction != Getfield && instruction != Putfield {
		panic("AddFieldInstruction can be used only for instructions that need FieldRef as parameter")
	}

	method.AddInstruction(instruction, method.bytecode.ConstPool.AddFieldRef(class, name, descriptor))
}

func (method *Method) AddInvokeInstruction(instruction Instruction, class string, name string, descriptor string) {
	if !strings.HasPrefix(instruction.Name, "Invoke") {
		panic("AddInvokeInstruction can be used only for invoke instructions")
	}

	method.AddInstruction(instruction, method.bytecode.ConstPool.AddMethodRef(class, name, descriptor))
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
	}

	if int(instruction.ParameterLength) != parametersLength {
		panic("could not create instruction " + instruction.Name + " with parameters length: " + strconv.Itoa(parametersLength) + ". It require " + strconv.Itoa(int(instruction.ParameterLength)) + " bytes as parameters")
	}
}

func (method *Method) toBytes() []byte {
	buffer := CreateBuffer()
	buffer.Write(method.AccessFlags)
	buffer.Write(method.NameIndex)
	buffer.Write(method.descriptorIndex)
	buffer.Write(method.attributesCount)

	if method.AccessFlags&Flag_Static != Flag_Static {
		method.maxLocals++
	}

	// MaxStack(u2) + MaxLocals(u2) + CodeLength(u4) + len(CodeBytes) + ExceptionTableLength(u2) + CodeAttributesCount(u2)
	method.codeAttributeLength = uint32(2 + 2 + 4 + len(method.codeBytes.bytes) + 2 + 2)

	buffer.Write(method.codeAttributeNameIndex)
	buffer.Write(method.codeAttributeLength)
	buffer.Write(method.maxStack)
	buffer.Write(method.maxLocals)
	buffer.Write(uint32(len(method.codeBytes.bytes)))
	buffer.Write(method.codeBytes.bytes)
	buffer.Write(method.exceptionTableLength)
	buffer.Write(method.codeAttributesCount)

	return buffer.bytes
}