package models

import (
	"errors"
	"time"
)

type Employee struct {
	Id       int
	Name     string  // Василий
	Phone    string  // +71234567890
	Birthday *string // RFC3339
}

func (e *Employee) Validate() error {
	if e.Id <= 0 {
		return errors.New("Id должен быть положительным числом")
	}
	if e.Name == "" {
		return errors.New("Имя не может быть пустым")
	}
	if e.Phone == "" {
		return errors.New("Номер телефона не может быть пустым")
	}
	if e.Birthday != nil {
		_, err := time.Parse(time.RFC3339, *e.Birthday)
		if err != nil {
			return errors.New("Некорректный формат даты рождения. Используйте RFC3339")
		}
	}
	return nil
}
