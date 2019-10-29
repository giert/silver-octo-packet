package main

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/giert/silver-octo-packet/assert"
	"github.com/giert/silver-octo-packet/mock"
)

func Test_Main(t *testing.T) {
	main()

	log.Printf("Look at these mockup variables:\n%s\n%s\n%s\n%s\n%s\n", mock.Addr, mock.Token, mock.Role, mock.File, mock.Path)

	log.Println("we can also assert things for testing")

	_, err := ioutil.ReadFile(mock.File)
	assert.Err(t, err, "open mock.file:")
	assert.Result(t, mock.Path, "mock/path")
}
