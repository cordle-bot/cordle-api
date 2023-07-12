package file

import (
	"errors"
	"os"

	"github.com/CordleBot/cordle-api/pkg/util"
)

func CreateFile(dir string) error {
	f, err := os.Create(dir)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(w string, dir string) error {
	e, err := CheckFile(dir)
	if err != nil {
		return err
	}
	if !e {
		err = CreateFile(dir)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(dir, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	b := []byte(w + "\n")
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFile(dir string) error {
	err := os.Remove(dir)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(dir string) ([]string, error) {
	e, err := CheckFile(dir)

	if err != nil {
		return nil, err
	}

	if !e {
		return nil, util.ErrFileNotExist
	}

	return nil, nil
}

func CheckFile(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func FileContains(c string, dir string) (bool, error) {
	return false, nil
}
