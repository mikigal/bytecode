package main

import (
	"bytecode"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	code := bytecode.CreateClass("Test", "java/lang/Object")
	code.AddInterface("java/io/Serializable")
	code.SourceFile("Test.java")

	// Add two global fields
	code.AddField(bytecode.Flag_Public|bytecode.Flag_Static|bytecode.Flag_Final, "firstNumber", "I", 5)
	code.AddField(bytecode.Flag_Public|bytecode.Flag_Static|bytecode.Flag_Final, "secondNumber", "I", 10)

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
	add := code.AddMethod(bytecode.Flag_Private|bytecode.Flag_Static, "add", "()I", 2)
	add.AddFieldInstruction(bytecode.Getstatic, "Test", "firstNumber", "I")
	add.AddFieldInstruction(bytecode.Getstatic, "Test", "secondNumber", "I")
	add.AddInstruction(bytecode.Iadd)
	add.AddInstruction(bytecode.Ireturn)

	writeFile(code)
}
func writeFile(code *bytecode.Class) {
	err := ioutil.WriteFile("examples/Test.class", code.ToBytes(), 644)
	Check(err)

	output, err := exec.Command("javap", "-c", "-p", "-verbose", "examples/Test.class").CombinedOutput()
	Check(err)
	bytecode.Log("javap -c -p -verbose \n%s", string(output))

	command := exec.Command("java", "Test")
	command.Dir = "examples"
	output, err = command.CombinedOutput()
	Check(err)
	bytecode.Log("java Test \n%s", string(output))
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
