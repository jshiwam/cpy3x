package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jshiwam/cpy3x/pycore"
)

var (
	output     = flag.String("o", "variadic", "output file (without extension)")
	caseNumber = flag.Int("n", pycore.MaxVariadicLength, "number of case in the switch statement")
)

func main() {
	flag.Parse()

	headerName := *output + ".h"

	headerFile, err := os.Create(headerName)
	if err != nil {
		fmt.Printf("Error opening file %s: %s", headerFile, err)
		os.Exit(1)
	}
	defer headerFile.Close()

	defn := strings.ToUpper(*output) + "_H"
	headerFile.WriteString(fmt.Sprintf("#ifndef %s\n", defn))
	headerFile.WriteString(fmt.Sprintf("#define %s\n\n", defn))
	headerFile.WriteString("#include \"Python.h\"\n\n")

	headerFile.WriteString(renderHeaderTemplate("PyObject_CallFunctionObjArgs", "callable"))
	headerFile.WriteString(renderHeaderTemplate("PyObject_CallMethodObjArgs", "obj", "name"))
	headerFile.WriteString("\n#endif\n")

	fileName := *output + ".c"
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %s", fileName, err)
		os.Exit(1)
	}
	defer outFile.Close()

	outFile.WriteString(fmt.Sprintf("#include \"%s\"\n\n", headerName))
	outFile.WriteString(renderTemplate(*caseNumber, "PyObject_CallFunctionObjArgs", "callable"))
	outFile.WriteString(renderTemplate(*caseNumber, "PyObject_CallMethodObjArgs", "obj", "name"))
}

func renderHeaderTemplate(functionName string, pyArgs ...string) string {
	template := `PyObject* _go_%s(%sint argc, PyObject **argv);
`
	args := ""
	for _, arg := range pyArgs {
		args += fmt.Sprintf(`PyObject *%s, `, arg)
	}
	return fmt.Sprintf(template, functionName, args)
}

func renderTemplate(n int, functionName string, pyArgs ...string) string {
	template := `PyObject* _go_%s(%sint argc, PyObject **argv) {

	PyObject *result = NULL;
	switch (argc){
	%s
	}
	return result;
}

`
	args := ""
	for _, arg := range pyArgs {
		args += fmt.Sprintf("PyObject *%s, ", arg)
	}
	switchTemplate := ""
	for i := 0; i < n+1; i++ {
		switchTemplate += fmt.Sprintf("		case %d:\n", i)
		args := ""
		for _, arg := range pyArgs {
			args += fmt.Sprintf("%s, ", arg)
		}
		for j := 0; j < i; j++ {
			args += fmt.Sprintf("argv[%d], ", j)
		}
		args += "NULL"
		switchTemplate += fmt.Sprintf("			return %s(%s);\n", functionName, args)
	}
	return fmt.Sprintf(template, functionName, args, switchTemplate)
}
