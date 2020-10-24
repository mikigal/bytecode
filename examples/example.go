package main

import (
    "bytecode"
    "io/ioutil"
    "os/exec"
)

func main() {
    code := bytecode.CreateClass("Test", "java/lang/Object")
    code.AddInterface("java/io/Serializable")
    code.SourceFile("Test.java")

    main := code.AddMainMethod(2)
    main.AddFieldInstruction(bytecode.Getstatic, "java/lang/System", "out", "Ljava/io/PrintStream;")
    main.AddInvokeInstruction(bytecode.Invokestatic, "java/util/UUID", "randomUUID", "()Ljava/util/UUID;")
    main.AddInvokeInstruction(bytecode.Invokevirtual, "java/io/PrintStream", "println", "(Ljava/lang/Object;)V")
    main.AddInstruction(bytecode.Return)

    writeFile(&code)
}

func writeFile(code *bytecode.Bytecode) {
    err := ioutil.WriteFile("examples/Test.class", code.ToBytes(), 644)
    bytecode.Check(err)

    output, err := exec.Command("javap", "-c", "-p", "-verbose", "examples/Test.class").CombinedOutput()
    bytecode.Check(err)
    bytecode.Log("javap -c -p -verbose \n%s", string(output))

    command := exec.Command("java", "Test")
    command.Dir = "examples"
    output, err = command.CombinedOutput()
    bytecode.Check(err)
    bytecode.Log("java Test \n%s", string(output))
}