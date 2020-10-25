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

	main := code.AddMainMethod(10)
	integer := main.AddLocalVar(900)
	long := main.AddLocalVar(int64(999999))
	float := main.AddLocalVar(float32(2.5))
	double := main.AddLocalVar(float64(5.6))
	text := main.AddLocalVar("test")
	char := main.AddLocalVar('c')
	boolean := main.AddLocalVar(true)

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(integer)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(I)V")

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(double)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(D)V")

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(float)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(F)V")

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(long)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(J)V")

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(text)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(Ljava/lang/String;)V")

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(char)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(C)V")

	main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.LoadLocalVar(boolean)
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(Z)V")

	/*main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
	main.AddInvokeInstruction(bytecode.Invokestatic, "java/util/UUID", "randomUUID", "()Ljava/util/UUID;")
	main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(Ljava/lang/Object;)V")*/

	main.AddInstruction(bytecode.Return)

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
