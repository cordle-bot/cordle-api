package file

import (
	"errors"
	"os"

	"github.com/cordle-bot/cordle-api/pkg/util"
)

// Create a file.
//
// Params:
//   - `dir` string : directory of file being created.
//
// Returns error.
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

// Write to a file, if it doesn't exist a new file is created.
//
// Params:
//   - `w` string : data being written to file.
//   - `dir` string : directory of file being written to.
//
// Returns error.
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

// Delete a file.
//
// Params:
//   - `dir` string : directory of file being deleted.
//
// Returns error.
func DeleteFile(dir string) error {
	err := os.Remove(dir)
	if err != nil {
		return err
	}

	return nil
}

// Read data from a file.
//
// Params:
//   - `dir` string : directory of file being read.
//
// Returns []string and error.
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

// Check if a file exists.
//
// Params:
//   - `dir` string : directory of file being read.
//
// Retruns bool and error.
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

// Checks if a file contains a given piece of data.
//
// Params:
//   - `c` string : data being compared to file data.
//   - `dir` string : directory of file being checked.
//
// Returns bool and error.
func FileContains(c string, dir string) (bool, error) {
	return false, nil
}
