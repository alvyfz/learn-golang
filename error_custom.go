package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "Alvy" {
		return &notFoundError{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("eko", nil)

	if err != nil {
		// terjadi error
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error", notFoundErr.Error())
		} else {
			fmt.Println("unknown error", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
