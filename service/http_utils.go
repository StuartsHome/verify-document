package service

import (
	"fmt"
	"net/http"
)

// Extract an optional string param from an HTTP request
func queryParamStringOptional(r *http.Request, paramName string, defaultValue *string) (string, error) {

	value, hasValue := r.Form[paramName]
	if !hasValue {
		if defaultValue == nil {
			return "", fmt.Errorf("no URL parameter %s", paramName)
		}
		return *defaultValue, nil
	}

	if len(value) > 1 {
		return "", fmt.Errorf("Multiple values specified for URL parameter %s", paramName)
	}

	return value[0], nil

}

// Attemp to extract a string param from an HTTP request, otherwise use the specified default
func queryParamStringDefault(r *http.Request, paramName string, defaultValue string) (string, error) {
	return queryParamStringOptional(r, paramName, &defaultValue)
}

func queryParamString(r *http.Request, paramName string) (string, error) {
	return queryParamStringOptional(r, paramName, nil)
}
