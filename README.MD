# bytecode
Basic tool for JVM bytecode generation written in Go

## Features
 - Generate valid bytecode and Java classes
 - Define classes, methods and global fields
 - API for creating local variables
 - Optimized constant pool generation
 - Automatic locals calculation

## TODO
 - Java 6+ support (missing StackMapTable implementation)
 - Automatic max stack calculation
 - Labels support
 
## Example
 ```go
package main

import (
    "github.com/mikigal/bytecode"
)

func main() {
    code := bytecode.CreateClass("Test", "java/lang/Object")
    code.AddInterface("java/io/Serializable")
    code.SourceFile("Test.java")

    // Add two global fields
    code.AddField(bytecode.Flag_Public | bytecode.Flag_Static | bytecode.Flag_Final, "firstNumber", "I", 5)
    code.AddField(bytecode.Flag_Public | bytecode.Flag_Static | bytecode.Flag_Final, "secondNumber", "I", 10)

    // Create main method
    main := code.AddMainMethod(2)
    localVar := main.AddLocalVar("test") // Define local String var
    main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
    main.LoadLocalVar(localVar)
    main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(Ljava/lang/String;)V")
    
    main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
    main.AddInvokeInstruction(bytecode.Invokestatic, "Test", "add", "()I")
    main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(I)V")
    main.AddInstruction(bytecode.Return)
    
    // Add second method
    add := code.AddMethod(bytecode.Flag_Private | bytecode.Flag_Static, "add", "()I", 2)
    add.AddFieldInstruction(bytecode.Getstatic, "Test", "firstNumber", "I")
    add.AddFieldInstruction(bytecode.Getstatic, "Test", "secondNumber", "I")
    add.AddInstruction(bytecode.Iadd)
    add.AddInstruction(bytecode.Ireturn)

    writeFile(code.ToBytes())
}
```
```
public class Test implements java.io.Serializable
  minor version: 0
  major version: 49
  flags: ACC_PUBLIC
Constant pool:
   #1 = Utf8               Test
   #2 = Class              #1             // Test
   #3 = Utf8               java/lang/Object
   #4 = Class              #3             // java/lang/Object
   #5 = Utf8               java/io/Serializable
   #6 = Class              #5             // java/io/Serializable
   #7 = Utf8               SourceFile
   #8 = Utf8               Test.java
   #9 = Utf8               firstNumber
  #10 = Utf8               I
  #11 = Utf8               ConstantValue
  #12 = Integer            5
  #13 = NameAndType        #9:#10         // firstNumber:I
  #14 = Fieldref           #2.#13         // Test.firstNumber:I
  #15 = Utf8               secondNumber
  #16 = Integer            10
  #17 = NameAndType        #15:#10        // secondNumber:I
  #18 = Fieldref           #2.#17         // Test.secondNumber:I
  #19 = Utf8               main
  #20 = Utf8               ([Ljava/lang/String;)V
  #21 = Utf8               Code
  #22 = Utf8               test
  #23 = String             #22            // test
  #24 = Utf8               java/lang/System
  #25 = Class              #24            // java/lang/System
  #26 = Utf8               out
  #27 = Utf8               Ljava/io/PrintStream;
  #28 = NameAndType        #26:#27        // out:Ljava/io/PrintStream;
  #29 = Fieldref           #25.#28        // java/lang/System.out:Ljava/io/PrintStream;
  #30 = Utf8               java/io/PrintStream
  #31 = Class              #30            // java/io/PrintStream
  #32 = Utf8               println
  #33 = Utf8               (Ljava/lang/String;)V
  #34 = NameAndType        #32:#33        // println:(Ljava/lang/String;)V
  #35 = Methodref          #31.#34        // java/io/PrintStream.println:(Ljava/lang/String;)V
  #36 = Utf8               add
  #37 = Utf8               ()I
  #38 = NameAndType        #36:#37        // add:()I
  #39 = Methodref          #2.#38         // Test.add:()I
  #40 = Utf8               (I)V
  #41 = NameAndType        #32:#40        // println:(I)V
  #42 = Methodref          #31.#41        // java/io/PrintStream.println:(I)V
{
  public static final int firstNumber;
    descriptor: I
    flags: ACC_PUBLIC, ACC_STATIC, ACC_FINAL
    ConstantValue: int 5

  public static final int secondNumber;
    descriptor: I
    flags: ACC_PUBLIC, ACC_STATIC, ACC_FINAL
    ConstantValue: int 10

  public static void main(java.lang.String[]);
    descriptor: ([Ljava/lang/String;)V
    flags: ACC_PUBLIC, ACC_STATIC
    Code:
      stack=2, locals=2, args_size=1
         0: ldc           #23                 // String test
         2: astore_1
         3: getstatic     #29                 // Field java/lang/System.out:Ljava/io/PrintStream;
         6: aload_1
         7: invokevirtual #35                 // Method java/io/PrintStream.println:(Ljava/lang/String;)V
        10: getstatic     #29                 // Field java/lang/System.out:Ljava/io/PrintStream;
        13: invokestatic  #39                 // Method add:()I
        16: invokevirtual #42                 // Method java/io/PrintStream.println:(I)V
        19: return

  private static int add();
    descriptor: ()I
    flags: ACC_PRIVATE, ACC_STATIC
    Code:
      stack=2, locals=0, args_size=0
         0: getstatic     #14                 // Field firstNumber:I
         3: getstatic     #18                 // Field secondNumber:I
         6: iadd
         7: ireturn
}
SourceFile: "Test.java"
```