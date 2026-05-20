package params

import (
	"errors"
	"reflect"
)

func getParametersFromURLValues(
	funcType reflect.Type,
	urlVals []string,
) ([]reflect.Value, error) {
	if len(urlVals)+1 != funcType.NumIn() {
		return nil, errors.New("Parameter number mismatch")
	}
	var err error
	params := make([]reflect.Value, len(urlVals))
	for i := range params {
		params[i], err = parseValueToType(funcType.In(i+1), urlVals[i])
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}
