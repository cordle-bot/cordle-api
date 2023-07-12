package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	file := "test.txt"

	err := CreateFile(file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.True(t, e)

	err = DeleteFile(file)
	assert.NoError(t, err)

	_, err = CheckFile(file)
	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	file := "test.txt"

	err := WriteFile("hello file", file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.True(t, e)

	err = DeleteFile(file)
	assert.NoError(t, err)

	_, err = CheckFile(file)
	assert.NoError(t, err)
}

func TestDeleteFile(t *testing.T) {
	file := "test.txt"

	err := CreateFile(file)
	assert.NoError(t, err)

	err = DeleteFile(file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.False(t, e)
}

func TestReadFile(t *testing.T) {

}

func TestCheckFile(t *testing.T) {
	g := "test.txt"
	b := "bad.txt"

	err := CreateFile(g)
	assert.NoError(t, err)

	e, err := CheckFile(g)
	assert.NoError(t, err)
	assert.True(t, e)

	err = DeleteFile(g)
	assert.NoError(t, err)

	e, err = CheckFile(g)
	assert.NoError(t, err)
	assert.False(t, e)

	e, err = CheckFile(b)
	assert.NoError(t, err)
	assert.False(t, e)
}
