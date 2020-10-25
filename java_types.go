package bytecode

type JavaType struct {
	Name  string
	Type  string
	Wide  bool
	Store map[int]Instruction
	Load  map[int]Instruction
	Const map[interface{}]Instruction
}

var numberLoad = map[int]Instruction{
	-1: Iload,
	0:  Iload_0,
	1:  Iload_1,
	2:  Iload_2,
	3:  Iload_3,
}

var numberStore = map[int]Instruction{
	-1: Istore,
	0:  Istore_0,
	1:  Istore_1,
	2:  Istore_2,
	3:  Istore_3,
}

var Type_Byte = JavaType{
	Name:  "byte",
	Type:  "B",
	Wide:  false,
	Store: numberStore,
	Load:  numberLoad,
	Const: map[interface{}]Instruction{
		int8(-1): Iconst_m1,
		int8(0):  Iconst_0,
		int8(1):  Iconst_1,
		int8(2):  Iconst_2,
		int8(3):  Iconst_3,
		int8(4):  Iconst_4,
		int8(5):  Iconst_5,
	},
}

var Type_Short = JavaType{
	Name:  "short",
	Type:  "S",
	Wide:  false,
	Store: numberStore,
	Load:  numberLoad,
	Const: map[interface{}]Instruction{
		int16(-1): Iconst_m1,
		int16(0):  Iconst_0,
		int16(1):  Iconst_1,
		int16(2):  Iconst_2,
		int16(3):  Iconst_3,
		int16(4):  Iconst_4,
		int16(5):  Iconst_5,
	},
}

var Type_Int = JavaType{
	Name:  "int",
	Type:  "I",
	Wide:  false,
	Store: numberStore,
	Load:  numberLoad,
	Const: map[interface{}]Instruction{
		-1: Iconst_m1,
		0:  Iconst_0,
		1:  Iconst_1,
		2:  Iconst_2,
		3:  Iconst_3,
		4:  Iconst_4,
		5:  Iconst_5,
	},
}

var Type_Char = JavaType{
	Name:  "char",
	Type:  "C",
	Wide:  false,
	Store: numberStore,
	Load:  numberLoad,
	Const: make(map[interface{}]Instruction, 0),
}

var Type_Boolean = JavaType{
	Name:  "boolean",
	Type:  "Z",
	Wide:  false,
	Store: numberStore,
	Load:  numberLoad,
	Const: map[interface{}]Instruction{
		true:  Iconst_1,
		false: Iconst_0,
	},
}

var Type_Float = JavaType{
	Name: "float",
	Type: "F",
	Wide: false,
	Store: map[int]Instruction{
		-1: Fstore,
		0:  Fstore_0,
		1:  Fstore_1,
		2:  Fstore_2,
		3:  Fstore_3,
	},
	Load: map[int]Instruction{
		-1: Fload,
		0:  Fload_0,
		1:  Fload_1,
		2:  Fload_2,
		3:  Fload_3,
	},
	Const: map[interface{}]Instruction{
		float32(0): Fconst_0,
		float32(1): Fconst_1,
		float32(2): Fconst_2,
	},
}

var Type_Double = JavaType{
	Name: "double",
	Type: "D",
	Wide: true,
	Store: map[int]Instruction{
		-1: Dstore,
		0:  Dstore_0,
		1:  Dstore_1,
		2:  Dstore_2,
		3:  Dstore_3,
	},
	Load: map[int]Instruction{
		-1: Dload,
		0:  Dload_0,
		1:  Dload_1,
		2:  Dload_2,
		3:  Dload_3,
	},
	Const: map[interface{}]Instruction{
		float64(0): Dconst_0,
		float64(1): Dconst_1,
	},
}

var Type_Long = JavaType{
	Name: "long",
	Type: "J",
	Wide: true,
	Store: map[int]Instruction{
		-1: Lstore,
		0:  Lstore_0,
		1:  Lstore_1,
		2:  Lstore_2,
		3:  Lstore_3,
	},
	Load: map[int]Instruction{
		-1: Lload,
		0:  Lload_0,
		1:  Lload_1,
		2:  Lload_2,
		3:  Lload_3,
	},
	Const: map[interface{}]Instruction{
		int64(0): Lconst_0,
		int64(1): Lconst_1,
	},
}

var Type_Reference = JavaType{
	Name: "reference",
	Type: "L",
	Wide: false,
	Store: map[int]Instruction{
		-1: Astore,
		0:  Astore_0,
		1:  Astore_1,
		2:  Astore_2,
		3:  Astore_3,
	},
	Load: map[int]Instruction{
		-1: Aload,
		0:  Aload_0,
		1:  Aload_1,
		2:  Aload_2,
		3:  Aload_3,
	},
	Const: map[interface{}]Instruction{
		nil: Aconst_null,
	},
}

var Type_Array = JavaType{
	Name:  "array",
	Type:  "[",
	Wide:  false,
	Store: make(map[int]Instruction, 0),         // TODO
	Load:  make(map[int]Instruction, 0),         // TODO
	Const: make(map[interface{}]Instruction, 0), // TODO
}
