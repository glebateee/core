package params

import (
	"encoding/json"
	"io"
	"net/url"
	"reflect"
	"strings"
)

func populateStructFromForm(
	structVal reflect.Value,
	formVals url.Values,
) error {
	for i := range structVal.Elem().Type().NumField() {
		field := structVal.Elem().Type().Field(i)
		for key, vals := range formVals {
			if strings.EqualFold(key, field.Name) && len(vals) > 0 {
				valField := structVal.Elem().Field(i)
				if valField.CanSet() {
					valToSet, err := parseValueToType(valField.Type(), vals[0])
					if err != nil {
						return err
					}
					valField.Set(valToSet)
				}
			}
		}
	}
	return nil
}

func populateStructFromJSON(structVal reflect.Value, reader io.ReadCloser) (err error) {
	return json.NewDecoder(reader).Decode(structVal.Interface())
}
