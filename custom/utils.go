package custom

import (
	"reflect"
	"runtime"
	"strings"
	"math/big"
	"bytes"
	"strconv"
)

func GetFuncName(funcPointer interface{},withPackage bool)string{
	funcName := runtime.FuncForPC(reflect.ValueOf(funcPointer).Pointer()).Name()
	if !withPackage{
		allParts := strings.Split(funcName,"/")
		return allParts[len(allParts) - 1]
	}
	return funcName
}


func GetStackString(stack []*big.Int) string{
	buffer := new(bytes.Buffer)
	for index,value := range stack{
		if value == big.NewInt(0){
			buffer.WriteByte('0')
		}
		for _,value2 := range value.Bytes(){
			hex := strconv.FormatInt(int64(value2),16)
			buffer.WriteString(hex)
		}
		if index != len(stack) -1 {
			buffer.WriteByte(' ')
		}

	}
	return buffer.String()
}