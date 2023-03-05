package validation

import (
	"fmt"
)

type User struct {
	Name     string
	Password string
	Age      int
}

type UserRule func(u User) error

type UserValidator struct {
	rules []UserRule
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) Add(r UserRule) {
	v.rules = append(v.rules, r)
}

func (v *UserValidator) ValidateUser(u User) []error {
	var errors []error
	for _, rule := range v.rules {
		if err := rule(u); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func (v *UserValidator) SelectUser(users []User) []User {
	var res []User
	for _, u := range users {
		e := v.ValidateUser(u)
		if len(e) == 0 {
			res = append(res, u)
		}
	}
	return res
}

func UserPasswordMinLength(min int) UserRule {
	return func(u User) error {
		if len(u.Password) < min {
			return fmt.Errorf("password must be greater than or equal to %d characters", min)
		}
		return nil
	}
}

func UserPasswordMaxLength(max int) UserRule {
	return func(u User) error {
		if len(u.Password) > max {
			return fmt.Errorf("password must be less than or equal to %d characters", max)
		}
		return nil
	}
}

func UserNameMinLength(min int) UserRule {
	return func(u User) error {
		if len(u.Name) < min {
			return fmt.Errorf("name must be greater than or equal to %d characters", min)
		}
		return nil
	}
}

func UserNameMaxLength(max int) UserRule {
	return func(u User) error {
		if len(u.Name) > max {
			return fmt.Errorf("name must be less than or equal to %d characters", max)
		}
		return nil
	}
}

func UserAge(min int) UserRule {
	return func(u User) error {
		if u.Age < min {
			return fmt.Errorf("age must be greater than or equal to %d", min)
		}
		return nil
	}
}
