package bytecode

type Class struct {
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
	Methods         []*Method
	attributesCount uint16
	attributes      *Buffer
}

func CreateClass(className string, superClass string) *Class {
	class := Class{
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
		Methods:         make([]*Method, 0),
		attributesCount: 0,
		attributes:      CreateBuffer(),
	}

	class.ThisClass = class.ConstPool.AddClass(className)
	class.SuperClass = class.ConstPool.AddClass(superClass)

	return &class
}

func (class *Class) AddInterface(className string) {
	class.Interfaces.Write(class.ConstPool.AddClass(className))
	class.InterfacesCount++
}

func (class *Class) SourceFile(fileName string) {
	class.attributes.Write(class.ConstPool.AddUtf8(Attribute_SourceFile))
	class.attributes.Write(uint32(2))
	class.attributes.Write(class.ConstPool.AddUtf8(fileName))
	class.attributesCount++
}

func (class *Class) AddMainMethod(maxStack uint16) *Method {
	return class.AddMethod(Flag_Public|Flag_Static, "main", "([Ljava/lang/String;)V", maxStack)
}

func (class *Class) AddMethod(accessFlags uint16, name string, descriptor string, maxStack uint16) *Method {
	method := Method{
		class:                  class,
		Name:                   name,
		Descriptor:             descriptor,
		LocalVars:              make(map[uint8]JavaType, 0),
		AccessFlags:            accessFlags,
		NameIndex:              class.ConstPool.AddUtf8(name),
		DescriptorIndex:        class.ConstPool.AddUtf8(descriptor),
		attributesCount:        1,
		attributes:             CreateBuffer(),
		codeAttributeNameIndex: class.ConstPool.AddUtf8(Attribute_Code),
		codeAttributeLength:    0,
		maxStack:               maxStack,
		maxLocals:              countParameters(descriptor) + Ternary(accessFlags&Flag_Static != Flag_Static, uint16(1), uint16(0)).(uint16),
		codeBytes:              CreateBuffer(),
		exceptionTableLength:   0,
		codeAttributesCount:    0,
	}

	class.Methods = append(class.Methods, &method)
	class.methodsCount++
	return &method
}

func (class *Class) AddField(accessFlags uint16, name string, descriptor string, constantValue interface{}) uint16 {
	nameIndex := class.ConstPool.AddUtf8(name)
	descriptorIndex := class.ConstPool.AddUtf8(descriptor)

	class.fields.Write(accessFlags)
	class.fields.Write(nameIndex)
	class.fields.Write(descriptorIndex)

	if constantValue != nil {
		if accessFlags&Flag_Final != Flag_Final {
			panic("could not set constant value to non-final field")
		}

		class.fields.Write(uint16(1)) // AttributesCount
		class.fields.Write(class.ConstPool.AddUtf8(Attribute_ConstantValue))
		class.fields.Write(uint32(2)) // AttributeLength

		switch value := constantValue.(type) {
		case string:
			class.fields.Write(class.ConstPool.AddString(value))
			break
		default:
			class.fields.Write(class.ConstPool.AddAuto(value))
		}
	} else {
		class.fields.Write(uint16(0))
	}

	class.fieldsCount++

	return class.ConstPool.AddFieldRef(class.ClassName, name, descriptor)
}

func (class *Class) ToBytes() []byte {
	buffer := CreateBuffer()
	buffer.Write(class.magic)
	buffer.Write(class.MinorVersion)
	buffer.Write(class.MajorVersion)
	buffer.Write(class.ConstPool.lastIndex + 1)
	buffer.Write(class.ConstPool.buffer.bytes)
	buffer.Write(class.AccessFlags)
	buffer.Write(class.ThisClass)
	buffer.Write(class.SuperClass)
	buffer.Write(class.InterfacesCount)
	buffer.Write(class.Interfaces.bytes)
	buffer.Write(class.fieldsCount)
	buffer.Write(class.fields.bytes)
	buffer.Write(class.methodsCount)
	for _, method := range class.Methods {
		buffer.Write(method.toBytes())
	}

	buffer.Write(class.attributesCount)
	buffer.Write(class.attributes.bytes)
	return buffer.bytes
}
