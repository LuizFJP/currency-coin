package main

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (a CoinItem) UpdateValidate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
	)
}

func (a CoinItem) CreateValidate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.StringOrBytes()),
		validation.Field(&a.Price, validation.Required),
	)
}
