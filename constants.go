package bytecode

type u1 = uint8
type u2 = uint16
type u4 = uint32

const Properties_MinorVersion u2 = 0
const Properties_MajorVersion u2 = 49
const Properties_MagicValue u4 = 0xCAFEBABE

const Flag_Public u2 = 0x0001
const Flag_Private u2 = 0x0002
const Flag_Protected u2 = 0x0004
const Flag_Static u2 = 0x0008
const Flag_Final u2 = 0x0010
const Flag_Synchronized u2 = 0x0020
const Flag_Bridge u2 = 0x0040
const Flag_Varargs u2 = 0x0080
const Flag_Native u2 = 0x0100
const Flag_Abstract u2 = 0x0400
const Flag_Strict u2 = 0x0800
const Flag_Synthetic u2 = 0x1000
const Flag_Super u2 = 0x0020
const Flag_Interface u2 = 0x0200
const Flag_Annotation u2 = 0x2000
const Flag_Enum u2 = 0x4000

const Constant_Class u1 = 7
const Constant_Fieldref u1 = 9
const Constant_Methodref u1 = 10
const Constant_InterfaceMethodref u1 = 11
const Constant_String u1 = 8
const Constant_Integer u1 = 3
const Constant_Float u1 = 4
const Constant_Long u1 = 5
const Constant_Double u1 = 6
const Constant_NameAndType u1 = 12
const Constant_Utf8 u1 = 1
const Constant_MethodHandle u1 = 15
const Constant_MethodType u1 = 16
const Constant_InvokeDynamic u1 = 18

const Attribute_ConstantValue string = "ConstantValue"
const Attribute_Code string = "Code"
const Attribute_StackMapTable string = "StackMapTable"
const Attribute_Exceptions string = "Exceptions"
const Attribute_InnerClasses string = "InnerClasses"
const Attribute_EnclosingMethod string = "EnclosingMethod"
const Attribute_Synthetic string = "Synthetic"
const Attribute_Signature string = "Signature"
const Attribute_SourceFile string = "SourceFile"
const Attribute_SourceDebugExtension string = "SourceDebugExtension"
const Attribute_LineNumberTable string = "LineNumberTable"
const Attribute_LocalVariableTable string = "LocalVariableTable"
const Attribute_LocalVariableTypeTable string = "LocalVariableTypeTable"
const Attribute_Deprecated string = "Deprecated"
const Attribute_RuntimeVisibleAnnotations string = "RuntimeVisibleAnnotations"
const Attribute_RuntimeInvisibleAnnotations string = "RuntimeInvisibleAnnotations"
const Attribute_RuntimeVisibleParameterAnnotations string = "RuntimeVisibleParameterAnnotations"
const Attribute_RuntimeInvisibleParameterAnnotations string = "RuntimeInvisibleParameterAnnotations"
const Attribute_AnnotationDefault string = "AnnotationDefault"

const Type_Byte = "B"
const Type_Char = "C"
const Type_Double = "D"
const Type_Float = "F"
const Type_Int = "I"
const Type_Long = "J"
const Type_Reference = "L"
const Type_Short = "S"
const Type_Boolean = "Z"
const Type_Array = "["

type Instruction struct {
	Name            string
	Value           u1
	ParameterLength u1
	CreateLocalVar  bool
}

var Aaload = Instruction{
	Name:            "Aaload",
	Value:           0x32,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Aastore = Instruction{
	Name:            "Aastore",
	Value:           0x53,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Aconst_null = Instruction{
	Name:            "Aconst_null",
	Value:           0x01,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Aload = Instruction{
	Name:            "Aload",
	Value:           0x19,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Aload_0 = Instruction{
	Name:            "Aload_0",
	Value:           0x2a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Aload_1 = Instruction{
	Name:            "Aload_1",
	Value:           0x2b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Aload_2 = Instruction{
	Name:            "Aload_2",
	Value:           0x2c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Aload_3 = Instruction{
	Name:            "Aload_3",
	Value:           0x2d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Anewarray = Instruction{
	Name:            "Anewarray",
	Value:           0xbd,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Areturn = Instruction{
	Name:            "Areturn",
	Value:           0xb0,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Arraylength = Instruction{
	Name:            "Arraylength",
	Value:           0xbe,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Astore = Instruction{
	Name:            "Astore",
	Value:           0x3a,
	ParameterLength: 1,
	CreateLocalVar:  true,
}
var Astore_0 = Instruction{
	Name:            "Astore_0",
	Value:           0x4b,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Astore_1 = Instruction{
	Name:            "Astore_1",
	Value:           0x4c,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Astore_2 = Instruction{
	Name:            "Astore_2",
	Value:           0x4d,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Astore_3 = Instruction{
	Name:            "Astore_3",
	Value:           0x4e,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Athrow = Instruction{
	Name:            "Athrow",
	Value:           0xbf,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Baload = Instruction{
	Name:            "Baload",
	Value:           0x33,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Bastore = Instruction{
	Name:            "Bastore",
	Value:           0x54,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Bipush = Instruction{
	Name:            "Bipush",
	Value:           0x10,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Breakpoint = Instruction{
	Name:            "Breakpoint",
	Value:           0xca,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Caload = Instruction{
	Name:            "Caload",
	Value:           0x34,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Castore = Instruction{
	Name:            "Castore",
	Value:           0x55,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Checkcast = Instruction{
	Name:            "Checkcast",
	Value:           0xc0,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var D2f = Instruction{
	Name:            "D2f",
	Value:           0x90,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var D2i = Instruction{
	Name:            "D2i",
	Value:           0x8e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var D2l = Instruction{
	Name:            "D2l",
	Value:           0x8f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dadd = Instruction{
	Name:            "Dadd",
	Value:           0x63,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Daload = Instruction{
	Name:            "Daload",
	Value:           0x31,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dastore = Instruction{
	Name:            "Dastore",
	Value:           0x52,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Dcmpg = Instruction{
	Name:            "Dcmpg",
	Value:           0x98,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dcmpl = Instruction{
	Name:            "Dcmpl",
	Value:           0x97,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dconst_0 = Instruction{
	Name:            "Dconst_0",
	Value:           0x0e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dconst_1 = Instruction{
	Name:            "Dconst_1",
	Value:           0x0f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ddiv = Instruction{
	Name:            "Ddiv",
	Value:           0x6f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dload = Instruction{
	Name:            "Dload",
	Value:           0x18,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Dload_0 = Instruction{
	Name:            "Dload_0",
	Value:           0x26,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dload_1 = Instruction{
	Name:            "Dload_1",
	Value:           0x27,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dload_2 = Instruction{
	Name:            "Dload_2",
	Value:           0x28,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dload_3 = Instruction{
	Name:            "Dload_3",
	Value:           0x29,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dmul = Instruction{
	Name:            "Dmul",
	Value:           0x6b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dneg = Instruction{
	Name:            "Dneg",
	Value:           0x77,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Drem = Instruction{
	Name:            "Drem",
	Value:           0x73,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dreturn = Instruction{
	Name:            "Dreturn",
	Value:           0xaf,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dstore = Instruction{
	Name:            "Dstore",
	Value:           0x39,
	ParameterLength: 1,
	CreateLocalVar:  true,
}
var Dstore_0 = Instruction{
	Name:            "Dstore_0",
	Value:           0x47,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Dstore_1 = Instruction{
	Name:            "Dstore_1",
	Value:           0x48,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Dstore_2 = Instruction{
	Name:            "Dstore_2",
	Value:           0x49,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Dstore_3 = Instruction{
	Name:            "Dstore_3",
	Value:           0x4a,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Dsub = Instruction{
	Name:            "Dsub",
	Value:           0x67,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dup = Instruction{
	Name:            "Dup",
	Value:           0x59,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dup_x1 = Instruction{
	Name:            "Dup_x1",
	Value:           0x5a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dup_x2 = Instruction{
	Name:            "Dup_x2",
	Value:           0x5b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dup2 = Instruction{
	Name:            "Dup2",
	Value:           0x5c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dup2_x1 = Instruction{
	Name:            "Dup2_x1",
	Value:           0x5d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Dup2_x2 = Instruction{
	Name:            "Dup2_x2",
	Value:           0x5e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var F2d = Instruction{
	Name:            "F2d",
	Value:           0x8d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var F2i = Instruction{
	Name:            "F2i",
	Value:           0x8b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var F2l = Instruction{
	Name:            "F2l",
	Value:           0x8c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fadd = Instruction{
	Name:            "Fadd",
	Value:           0x62,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Faload = Instruction{
	Name:            "Faload",
	Value:           0x30,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fastore = Instruction{
	Name:            "Fastore",
	Value:           0x51,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Fcmpg = Instruction{
	Name:            "Fcmpg",
	Value:           0x96,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fcmpl = Instruction{
	Name:            "Fcmpl",
	Value:           0x95,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fconst_0 = Instruction{
	Name:            "Fconst_0",
	Value:           0x0b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fconst_1 = Instruction{
	Name:            "Fconst_1",
	Value:           0x0c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fconst_2 = Instruction{
	Name:            "Fconst_2",
	Value:           0x0d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fdiv = Instruction{
	Name:            "Fdiv",
	Value:           0x6e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fload = Instruction{
	Name:            "Fload",
	Value:           0x17,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Fload_0 = Instruction{
	Name:            "Fload_0",
	Value:           0x22,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fload_1 = Instruction{
	Name:            "Fload_1",
	Value:           0x23,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fload_2 = Instruction{
	Name:            "Fload_2",
	Value:           0x24,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fload_3 = Instruction{
	Name:            "Fload_3",
	Value:           0x25,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fmul = Instruction{
	Name:            "Fmul",
	Value:           0x6a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fneg = Instruction{
	Name:            "Fneg",
	Value:           0x76,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Frem = Instruction{
	Name:            "Frem",
	Value:           0x72,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Freturn = Instruction{
	Name:            "Freturn",
	Value:           0xae,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Fstore = Instruction{
	Name:            "Fstore",
	Value:           0x38,
	ParameterLength: 1,
	CreateLocalVar:  true,
}
var Fstore_0 = Instruction{
	Name:            "Fstore_0",
	Value:           0x43,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Fstore_1 = Instruction{
	Name:            "Fstore_1",
	Value:           0x44,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Fstore_2 = Instruction{
	Name:            "Fstore_2",
	Value:           0x45,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Fstore_3 = Instruction{
	Name:            "Fstore_3",
	Value:           0x46,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Fsub = Instruction{
	Name:            "Fsub",
	Value:           0x66,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Getfield = Instruction{
	Name:            "Getfield",
	Value:           0xb4,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Getstatic = Instruction{
	Name:            "Getstatic",
	Value:           0xb2,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Goto = Instruction{
	Name:            "Goto",
	Value:           0xa7,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Goto_w = Instruction{
	Name:            "Goto_w",
	Value:           0xc8,
	ParameterLength: 4,
	CreateLocalVar:  false,
}
var I2b = Instruction{
	Name:            "I2b",
	Value:           0x91,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var I2c = Instruction{
	Name:            "I2c",
	Value:           0x92,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var I2d = Instruction{
	Name:            "I2d",
	Value:           0x87,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var I2f = Instruction{
	Name:            "I2f",
	Value:           0x86,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var I2l = Instruction{
	Name:            "I2l",
	Value:           0x85,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var I2s = Instruction{
	Name:            "I2s",
	Value:           0x93,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iadd = Instruction{
	Name:            "Iadd",
	Value:           0x60,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iaload = Instruction{
	Name:            "Iaload",
	Value:           0x2e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iand = Instruction{
	Name:            "Iand",
	Value:           0x7e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iastore = Instruction{
	Name:            "Iastore",
	Value:           0x4f,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Iconst_m1 = Instruction{
	Name:            "Iconst_m1",
	Value:           0x02,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iconst_0 = Instruction{
	Name:            "Iconst_0",
	Value:           0x03,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iconst_1 = Instruction{
	Name:            "Iconst_1",
	Value:           0x04,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iconst_2 = Instruction{
	Name:            "Iconst_2",
	Value:           0x05,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iconst_3 = Instruction{
	Name:            "Iconst_3",
	Value:           0x06,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iconst_4 = Instruction{
	Name:            "Iconst_4",
	Value:           0x07,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iconst_5 = Instruction{
	Name:            "Iconst_5",
	Value:           0x08,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Idiv = Instruction{
	Name:            "Idiv",
	Value:           0x6c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var If_acmpeq = Instruction{
	Name:            "If_acmpeq",
	Value:           0xa5,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_acmpne = Instruction{
	Name:            "If_acmpne",
	Value:           0xa6,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_icmpeq = Instruction{
	Name:            "If_icmpeq",
	Value:           0x9f,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_icmpge = Instruction{
	Name:            "If_icmpge",
	Value:           0xa2,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_icmpgt = Instruction{
	Name:            "If_icmpgt",
	Value:           0xa3,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_icmple = Instruction{
	Name:            "If_icmple",
	Value:           0xa4,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_icmplt = Instruction{
	Name:            "If_icmplt",
	Value:           0xa1,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var If_icmpne = Instruction{
	Name:            "If_icmpne",
	Value:           0xa0,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifeq = Instruction{
	Name:            "Ifeq",
	Value:           0x99,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifge = Instruction{
	Name:            "Ifge",
	Value:           0x9c,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifgt = Instruction{
	Name:            "Ifgt",
	Value:           0x9d,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifle = Instruction{
	Name:            "Ifle",
	Value:           0x9e,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Iflt = Instruction{
	Name:            "Iflt",
	Value:           0x9b,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifne = Instruction{
	Name:            "Ifne",
	Value:           0x9a,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifnonnull = Instruction{
	Name:            "Ifnonnull",
	Value:           0xc7,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ifnull = Instruction{
	Name:            "Ifnull",
	Value:           0xc6,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Iinc = Instruction{
	Name:            "Iinc",
	Value:           0x84,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Iload = Instruction{
	Name:            "Iload",
	Value:           0x15,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Iload_0 = Instruction{
	Name:            "Iload_0",
	Value:           0x1a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iload_1 = Instruction{
	Name:            "Iload_1",
	Value:           0x1b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iload_2 = Instruction{
	Name:            "Iload_2",
	Value:           0x1c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iload_3 = Instruction{
	Name:            "Iload_3",
	Value:           0x1d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Impdep1 = Instruction{
	Name:            "Impdep1",
	Value:           0xfe,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Impdep2 = Instruction{
	Name:            "Impdep2",
	Value:           0xff,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Imul = Instruction{
	Name:            "Imul",
	Value:           0x68,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ineg = Instruction{
	Name:            "Ineg",
	Value:           0x74,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Instanceof = Instruction{
	Name:            "Instanceof",
	Value:           0xc1,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Invokedynamic = Instruction{
	Name:            "Invokedynamic",
	Value:           0xba,
	ParameterLength: 4,
	CreateLocalVar:  false,
}
var Invokeinterface = Instruction{
	Name:            "Invokeinterface",
	Value:           0xb9,
	ParameterLength: 4,
	CreateLocalVar:  false,
}
var Invokespecial = Instruction{
	Name:            "Invokespecial",
	Value:           0xb7,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Invokestatic = Instruction{
	Name:            "Invokestatic",
	Value:           0xb8,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Invokevirtual = Instruction{
	Name:            "Invokevirtual",
	Value:           0xb6,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ior = Instruction{
	Name:            "Ior",
	Value:           0x80,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Irem = Instruction{
	Name:            "Irem",
	Value:           0x70,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ireturn = Instruction{
	Name:            "Ireturn",
	Value:           0xac,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ishl = Instruction{
	Name:            "Ishl",
	Value:           0x78,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ishr = Instruction{
	Name:            "Ishr",
	Value:           0x7a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Istore = Instruction{
	Name:            "Istore",
	Value:           0x36,
	ParameterLength: 1,
	CreateLocalVar:  true,
}
var Istore_0 = Instruction{
	Name:            "Istore_0",
	Value:           0x3b,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Istore_1 = Instruction{
	Name:            "Istore_1",
	Value:           0x3c,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Istore_2 = Instruction{
	Name:            "Istore_2",
	Value:           0x3d,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Istore_3 = Instruction{
	Name:            "Istore_3",
	Value:           0x3e,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Isub = Instruction{
	Name:            "Isub",
	Value:           0x64,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Iushr = Instruction{
	Name:            "Iushr",
	Value:           0x7c,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ixor = Instruction{
	Name:            "Ixor",
	Value:           0x82,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Jsr = Instruction{
	Name:            "Jsr",
	Value:           0xa8,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Jsr_w = Instruction{
	Name:            "Jsr_w",
	Value:           0xc9,
	ParameterLength: 4,
	CreateLocalVar:  false,
}
var L2d = Instruction{
	Name:            "L2d",
	Value:           0x8a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var L2f = Instruction{
	Name:            "L2f",
	Value:           0x89,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var L2i = Instruction{
	Name:            "L2i",
	Value:           0x88,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ladd = Instruction{
	Name:            "Ladd",
	Value:           0x61,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Laload = Instruction{
	Name:            "Laload",
	Value:           0x2f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Land = Instruction{
	Name:            "Land",
	Value:           0x7f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lastore = Instruction{
	Name:            "Lastore",
	Value:           0x50,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Lcmp = Instruction{
	Name:            "Lcmp",
	Value:           0x94,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lconst_0 = Instruction{
	Name:            "Lconst_0",
	Value:           0x09,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lconst_1 = Instruction{
	Name:            "Lconst_1",
	Value:           0x0a,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Ldc = Instruction{
	Name:            "Ldc",
	Value:           0x12,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Ldc_w = Instruction{
	Name:            "Ldc_w",
	Value:           0x13,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ldc2_w = Instruction{
	Name:            "Ldc2_w",
	Value:           0x14,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ldiv = Instruction{
	Name:            "Ldiv",
	Value:           0x6d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lload = Instruction{
	Name:            "Lload",
	Value:           0x16,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Lload_0 = Instruction{
	Name:            "Lload_0",
	Value:           0x1e,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lload_1 = Instruction{
	Name:            "Lload_1",
	Value:           0x1f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lload_2 = Instruction{
	Name:            "Lload_2",
	Value:           0x20,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lload_3 = Instruction{
	Name:            "Lload_3",
	Value:           0x21,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lmul = Instruction{
	Name:            "Lmul",
	Value:           0x69,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lneg = Instruction{
	Name:            "Lneg",
	Value:           0x75,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lookupswitch = Instruction{
	Name:            "Lookupswitch",
	Value:           0xab,
	ParameterLength: 8,
	CreateLocalVar:  false,
}
var Lor = Instruction{
	Name:            "Lor",
	Value:           0x81,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lrem = Instruction{
	Name:            "Lrem",
	Value:           0x71,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lreturn = Instruction{
	Name:            "Lreturn",
	Value:           0xad,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lshl = Instruction{
	Name:            "Lshl",
	Value:           0x79,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lshr = Instruction{
	Name:            "Lshr",
	Value:           0x7b,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lstore = Instruction{
	Name:            "Lstore",
	Value:           0x37,
	ParameterLength: 1,
	CreateLocalVar:  true,
}
var Lstore_0 = Instruction{
	Name:            "Lstore_0",
	Value:           0x3f,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Lstore_1 = Instruction{
	Name:            "Lstore_1",
	Value:           0x40,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Lstore_2 = Instruction{
	Name:            "Lstore_2",
	Value:           0x41,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Lstore_3 = Instruction{
	Name:            "Lstore_3",
	Value:           0x42,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Lsub = Instruction{
	Name:            "Lsub",
	Value:           0x65,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lushr = Instruction{
	Name:            "Lushr",
	Value:           0x7d,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Lxor = Instruction{
	Name:            "Lxor",
	Value:           0x83,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Monitorenter = Instruction{
	Name:            "Monitorenter",
	Value:           0xc2,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Monitorexit = Instruction{
	Name:            "Monitorexit",
	Value:           0xc3,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Multianewarray = Instruction{
	Name:            "Multianewarray",
	Value:           0xc5,
	ParameterLength: 3,
	CreateLocalVar:  false,
}
var New = Instruction{
	Name:            "New",
	Value:           0xbb,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Newarray = Instruction{
	Name:            "Newarray",
	Value:           0xbc,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Nop = Instruction{
	Name:            "Nop",
	Value:           0x00,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Pop = Instruction{
	Name:            "Pop",
	Value:           0x57,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Pop2 = Instruction{
	Name:            "Pop2",
	Value:           0x58,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Putfield = Instruction{
	Name:            "Putfield",
	Value:           0xb5,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Putstatic = Instruction{
	Name:            "Putstatic",
	Value:           0xb3,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Ret = Instruction{
	Name:            "Ret",
	Value:           0xa9,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Return = Instruction{
	Name:            "Return",
	Value:           0xb1,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Saload = Instruction{
	Name:            "Saload",
	Value:           0x35,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Sastore = Instruction{
	Name:            "Sastore",
	Value:           0x56,
	ParameterLength: 0,
	CreateLocalVar:  true,
}
var Sipush = Instruction{
	Name:            "Sipush",
	Value:           0x11,
	ParameterLength: 2,
	CreateLocalVar:  false,
}
var Swap = Instruction{
	Name:            "Swap",
	Value:           0x5f,
	ParameterLength: 0,
	CreateLocalVar:  false,
}
var Tableswitch = Instruction{
	Name:            "Tableswitch",
	Value:           0xaa,
	ParameterLength: 1,
	CreateLocalVar:  false,
}
var Wide = Instruction{
	Name:            "Wide",
	Value:           0xc4,
	ParameterLength: 3,
	CreateLocalVar:  false,
}
