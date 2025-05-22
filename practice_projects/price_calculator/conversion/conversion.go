package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64,error){
	floats := make([]float64,len(strings))
	for stringValIndex,stringVal := range strings{
		floatString,err := strconv.ParseFloat(stringVal,64)
		if err != nil {
			return nil,errors.New("string to float conversion failed")
		}
		floats[stringValIndex] = floatString
	}

	return floats,nil
}