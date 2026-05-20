package validation

import (
	"reflect"
	"strings"
)

type TagValidator struct {
	validators map[string]ValidatorFunc
}

func (t *TagValidator) Validate(tgt any) (bool, []ValidationError) {
	errs := []ValidationError{}
	tgtVal := reflect.ValueOf(tgt)
	if tgtVal.Kind() == reflect.Pointer {
		tgtVal = tgtVal.Elem()
	}
	if tgtVal.Kind() != reflect.Struct {
		panic("Only structs can be validated")
	}
	for i := range tgtVal.NumField() {
		fType := tgtVal.Type().Field(i)
		vTag, found := fType.Tag.Lookup("validation")
		if found {
			for _, v := range strings.Split(vTag, ",") {
				var name, arg string
				if strings.Contains(v, ":") {
					nameAndArgs := strings.SplitN(v, ":", 2)
					name = nameAndArgs[0]
					arg = nameAndArgs[1]
				} else {
					name = v
				}

				if validator, ok := t.validators[name]; ok {
					if valid, err := validator(fType.Name, tgtVal.Field(i).Interface(), arg); !valid {
						errs = append(errs, ValidationError{
							FieldName: fType.Name,
							Error:     err,
						})
					}
				} else {
					panic("unknown validator")
				}
			}
		}
	}
	return len(errs) == 0, errs
}

func NewTagValidator(v map[string]ValidatorFunc) Validator {
	return &TagValidator{validators: v}
}
