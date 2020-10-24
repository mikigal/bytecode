package bytecode

type Bytecode struct {
	ClassName       string
	magic           uint32
	MinorVersion    uint16
	MajorVersion    uint16
	ConstPool       *ConstantPool
	AccessFlags     uint16
	ThisClass       uint16
	SuperClass      uint16
	InterfacesCount uint16
	Interfaces      *Buffer
	fieldsCount     uint16
	fields          *Buffer
	methodsCount    uint16
	Methods         []Method
	attributesCount uint16
	attributes      *Buffer
}

func CreateClass(className string, superClass string) Bytecode {
	bytecode := Bytecode{
		ClassName:       className,
		magic:           Properties_MagicValue,
		MinorVersion:    Properties_MinorVersion,
		MajorVersion:    Properties_MajorVersion,
		ConstPool:       CreateConstantInfo(),
		AccessFlags:     Flag_Public,
		InterfacesCount: 0,
		Interfaces:      CreateBuffer(),
		fieldsCount:     0,
		fields:          CreateBuffer(),
		methodsCount:    0,
		Methods:         make([]Method, 0),
		attributesCount: 0,
		attributes:      CreateBuffer(),
	}

	bytecode.ThisClass = bytecode.ConstPool.AddClass(className)
	bytecode.SuperClass = bytecode.ConstPool.AddClass(superClass)

	return bytecode
}

func (bytecode *Bytecode) AddInterface(class string) {
	bytecode.Interfaces.Write(bytecode.ConstPool.AddClass(class))
	bytecode.InterfacesCount++
}

func (bytecode *Bytecode) SourceFile(fileName string) {
	bytecode.attributes.Write(bytecode.ConstPool.AddUtf8(Attribute_SourceFile))
	bytecode.attributes.Write(uint32(2))
	bytecode.attributes.Write(bytecode.ConstPool.AddUtf8(fileName))
	bytecode.attributesCount++
}

func (bytecode *Bytecode) AddMainMethod(maxStack uint16) Method {
	return bytecode.AddMethod(Flag_Public | Flag_Static | Flag_Varargs, "main", "([Ljava/lang/String;)V", maxStack)
}

func (bytecode *Bytecode) AddMethod(accessFlags uint16, name string, descriptor string, maxStack uint16) Method {
	method := Method{
		bytecode:               bytecode,
		Name:                   name,
		Descriptor:             descriptor,
		AccessFlags:            accessFlags,
		NameIndex:              bytecode.ConstPool.AddUtf8(name),
		descriptorIndex:        bytecode.ConstPool.AddUtf8(descriptor),
		attributesCount:        1,
		codeAttributeNameIndex: bytecode.ConstPool.AddUtf8(Attribute_Code),
		codeAttributeLength:    0,
		maxStack:               maxStack,
		maxLocals:              countParameters(descriptor),
		codeBytes:              CreateBuffer(),
		exceptionTableLength:   0,
		codeAttributesCount:    0,
	}

	bytecode.Methods = append(bytecode.Methods, method)
	bytecode.methodsCount++
	return method
}

func (bytecode *Bytecode) AddField(accessFlags uint16, name string, descriptor string, constantValue interface{}) uint16 {
	nameIndex := bytecode.ConstPool.AddUtf8(name)
	descriptorIndex := bytecode.ConstPool.AddUtf8(descriptor)

	bytecode.fields.Write(accessFlags)
	bytecode.fields.Write(nameIndex)
	bytecode.fields.Write(descriptorIndex)

	if constantValue != nil {
		if accessFlags&Flag_Final != Flag_Final {
			panic("could not set constant value to non-final field")
		}

		bytecode.fields.Write(uint16(1)) // AttributesCount
		bytecode.fields.Write(bytecode.ConstPool.AddUtf8(Attribute_ConstantValue))
		bytecode.fields.Write(uint32(2)) // AttributeLength

		switch value := constantValue.(type) {
		case string:
			bytecode.fields.Write(bytecode.ConstPool.AddString(value))
			break
		default:
			bytecode.fields.Write(bytecode.ConstPool.AddAuto(value))
		}
	} else {
		bytecode.fields.Write(uint16(0))
	}

	bytecode.fieldsCount++

	return bytecode.ConstPool.AddFieldRef(bytecode.ClassName, name, descriptor)
}

func (bytecode *Bytecode) ToBytes() []byte {
	buffer := CreateBuffer()
	buffer.Write(bytecode.magic)
	buffer.Write(bytecode.MinorVersion)
	buffer.Write(bytecode.MajorVersion)
	buffer.Write(bytecode.ConstPool.lastIndex + 1)
	buffer.Write(bytecode.ConstPool.buffer.bytes)
	buffer.Write(bytecode.AccessFlags)
	buffer.Write(bytecode.ThisClass)
	buffer.Write(bytecode.SuperClass)
	buffer.Write(bytecode.InterfacesCount)
	buffer.Write(bytecode.Interfaces.bytes)
	buffer.Write(bytecode.fieldsCount)
	buffer.Write(bytecode.fields.bytes)
	buffer.Write(bytecode.methodsCount)
	for _, method := range bytecode.Methods {
		buffer.Write(method.toBytes())
	}

	buffer.Write(bytecode.attributesCount)
	buffer.Write(bytecode.attributes.bytes)
	return buffer.bytes
}