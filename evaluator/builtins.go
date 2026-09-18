package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/SegniAT/monkey-language-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if length := len(arr.Elements); length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if length := len(arr.Elements); length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"readFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			fileName := args[0]
			if fileName.Type() != object.STRING_OBJ {
				return newError("argument to `readFile` must be STRING, got %s", fileName.Type())
			}

			strObj, _ := fileName.(*object.String)

			absPath, err := filepath.Abs(strObj.Value)
			if err != nil {
				return newError("could not resolve path %q: %s", strObj.Value, err.Error())
			}

			byteContent, err := os.ReadFile(absPath)
			if err != nil {
				return newError("could not read file %q: %s", absPath, err.Error())
			}

			return &object.String{Value: string(byteContent)}
		},
	},
	"splitString": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			str, sep := args[0], args[1]
			if str.Type() != object.STRING_OBJ || sep.Type() != object.STRING_OBJ {
				return newError("argument to `splitString` must be (STRING, STRING), got %s,%s", str.Type(), sep.Type())
			}

			strObj, _ := str.(*object.String)
			sepObj, _ := sep.(*object.String)
			splitRes := strings.Split(strObj.Value, sepObj.Value)

			result := &object.Array{
				Elements: make([]object.Object, len(splitRes)),
			}

			for i, s := range splitRes {
				result.Elements[i] = &object.String{
					Value: s,
				}
			}

			return result
		},
	},
	"atoi": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			numStr := args[0]
			if numStr.Type() != object.STRING_OBJ {
				return newError("argument to `atoi` must be STRING, got %s", numStr.Type())
			}

			numStrObj, _ := numStr.(*object.String)
			num, err := strconv.Atoi(numStrObj.Value)
			if err != nil {
				return newError("could not convert %q to integer: %s", numStr.Inspect(), err.Error())
			}

			return &object.Integer{
				Value: int64(num),
			}
		},
	},
	"sortInts": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			arr := args[0]
			if arr.Type() != object.ARRAY_OBJ {
				return newError("argument to `sortInts` must be ARRAY, got %s", arr.Type())
			}

			arrObj, ok := arr.(*object.Array)
			if !ok {
				return newError("argument to `sortInts` must be ARRAY")
			}

			// Each element should be integer
			for i, el := range arrObj.Elements {
				if el.Type() == object.INTEGER_OBJ {
					continue
				}

				return newError("argument to `sortInts` must be ARRAY of integers, got %s at index %d", el.Type(), i)
			}

			goArr := make([]int64, len(arrObj.Elements))
			for i, el := range arrObj.Elements {
				inObj, _ := el.(*object.Integer)
				goArr[i] = inObj.Value
			}

			slices.Sort(goArr)

			resArr := &object.Array{
				Elements: make([]object.Object, len(goArr)),
			}

			for i, el := range goArr {
				resArr.Elements[i] = &object.Integer{
					Value: el,
				}
			}

			return resArr
		},
	},
	"trimSpace": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			str, ok := args[0].(*object.String)
			if !ok {
				return newError("argument to `trimSpace` must be STRING, got %s", args[0].Type())
			}

			return &object.String{Value: strings.TrimSpace(str.Value)}
		},
	},
	"trimSuffix": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("argument to `trimSuffix` must be (STRING, STRING), got (%s, %s)", args[0].Type(), args[1].Type())
			}

			str, ok := args[0].(*object.String)
			if !ok {
				return newError("first argument to `trimSuffix` must be (STRING, STRING), got %s", args[0].Type())
			}

			suffix, ok := args[1].(*object.String)
			if !ok {
				return newError("second argument to `trimSuffix` must be STRING, got %s", args[1].Type())
			}

			return &object.String{Value: strings.TrimSuffix(str.Value, suffix.Value)}
		},
	},
}

func BuiltinNames() []string {
	names := make([]string, 0, len(builtins))
	for name := range builtins {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}
