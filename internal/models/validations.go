package models

type validations interface {
	Validate() error
}

func Validate(v validations) error {
	return v.Validate()
}
