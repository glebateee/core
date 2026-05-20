package validation

type Validator interface {
	Validate(tgt any) (bool, []ValidationError)
}

type ValidationError struct {
	FieldName string
	Error     error
}

type ValidatorFunc func(fieldName string, val any, arg string) (bool, error)

func DefaultValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required": required,
		"min":      min,
	}
}
