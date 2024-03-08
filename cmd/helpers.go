package cmd

import (
	"strconv"
)

type ManUnit struct {
	Name        string
	Description string
}

type NullableInt struct {
	IsSet bool
	Val   int
}

func (i *NullableInt) String() string {
	if !i.IsSet {
		return "<not set>"
	}
	return strconv.Itoa(i.Val)
}

func (i *NullableInt) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	i.IsSet = true
	i.Val = v
	return nil
}
