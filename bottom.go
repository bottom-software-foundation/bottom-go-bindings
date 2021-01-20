package bottom

import (
	"fmt"
	"os"

	"github.com/NCPlayz/go-python3"
)

func encodeBytes(value string) string {
	module := importBottom()

	encodeFunc := module.GetAttrString("encode")

	args := python3.PyTuple_New(1)
	valueAsUnicode := python3.PyBytes_FromString(value)
	python3.PyTuple_SetItem(args, 0, valueAsUnicode)

	returnValue := encodeFunc.CallObject(args)
	return python3.PyUnicode_AsUTF8(returnValue)
}

func encode(value string) string {
	module := importBottom()

	encodeFunc := module.GetAttrString("encode")

	args := python3.PyTuple_New(1)
	valueAsUnicode := python3.PyUnicode_FromString(value)
	python3.PyTuple_SetItem(args, 0, valueAsUnicode)

	returnValue := encodeFunc.CallObject(args)
	return python3.PyUnicode_AsUTF8(returnValue)
}

func decode(value string) (string, error) {
	module := importBottom()

	decodeFunc := module.GetAttrString("decode")

	args := python3.PyTuple_New(1)
	valueAsUnicode := python3.PyUnicode_FromString(value)
	python3.PyTuple_SetItem(args, 0, valueAsUnicode)

	returnValue := decodeFunc.CallObject(args)

	if returnValue == nil {
		python3.PyErr_Print()
		return "", fmt.Errorf("failed to decode.")
	}

	return python3.PyUnicode_AsUTF8(returnValue), nil
}

func importBottom() *python3.PyObject {
	python3.Py_Initialize()

	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}

	bottom := python3.PyImport_ImportModule("bottom")

	if bottom == nil {
		fmt.Println("Could not find bottom module.")
		os.Exit(1)
	}

	return bottom
}
