package dict

import (
	"errors"
	"fmt"
)

type dictionary map[string]string

var (
	errKeyValueNotExist = errors.New("error: Value not found")
	errKeyValueExist    = errors.New("error: key value pair already exists")
	errCannotUpdate     = errors.New("error: cannot update non-existent key value pair")
	errCannotDelete     = errors.New("error: cannot delete non-existent key value pair")
)

func CreateDictionary() (d dictionary) {
	d = make(map[string]string)
	return
}

func (dictionary) PrintDictionary(d dictionary) {
	for k, v := range d {
		fmt.Println("word,", k)
		fmt.Println("Definition,", v, "\n")
	}
}

func (d dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if !exists {
		return "", errKeyValueNotExist
	}
	return value, nil
}

func (d dictionary) Add(key string, value string) error {
	switch _, err := d.Search(key); err {
	case errKeyValueNotExist:
		d[key] = value
	case nil:
		return errKeyValueExist
	}
	return nil
}

func (d dictionary) Update(key string, newValue string) error {
	switch _, err := d.Search(key); err {
	case errKeyValueNotExist:
		return errCannotUpdate
	case nil:
		d[key] = newValue
	}
	return nil
}

func (d dictionary) Delete(key string) error {
	switch _, err := d.Search(key); err {
	case errKeyValueNotExist:
		return errCannotDelete
	case nil:
		delete(d, key)
	}
	return nil
}
