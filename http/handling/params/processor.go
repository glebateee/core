package params

import (
	"net/http"
	"reflect"
)

func GetParametersFromRequest(r *http.Request, handlerMethod reflect.Method, urlVals []string) ([]reflect.Value, error) {
	handlerMethodType := handlerMethod.Type

	if handlerMethodType.NumIn() == 1 {
		return []reflect.Value{}, nil
	}

	if handlerMethodType.NumIn() == 2 && handlerMethodType.In(1).Kind() == reflect.Struct {
		dtoPtr := reflect.New(handlerMethodType.In(1))
		if err := r.ParseForm(); err != nil {
			return []reflect.Value{}, err
		}

		// Всегда пытаемся распарсить форму
		if err := populateStructFromForm(dtoPtr, r.Form); err != nil {
			return []reflect.Value{}, err
		}

		// Если Content-Type JSON, сначала парсим тело (JSON), потом форма может перезаписать поля
		if getContentType(r) == "application/json" {
			if err := populateStructFromJSON(dtoPtr, r.Body); err != nil {
				return []reflect.Value{}, err
			}
			// форму уже парсили выше, она остаётся с более высоким приоритетом
		}

		return []reflect.Value{dtoPtr.Elem()}, nil
	}

	return getParametersFromURLValues(handlerMethodType, urlVals)
}

func getContentType(request *http.Request) string {
	headerSlice := request.Header["Content-Type"]
	if len(headerSlice) > 0 {
		return headerSlice[0]
	}
	return ""
}
