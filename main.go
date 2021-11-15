package main

import (
	"os"
	"s3bench/cmd"
	"s3bench/object"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

type config struct {
	object *object.Object
}

func init() {
	log = logrus.New()
	log.Out = os.Stdout
}

func main() {
	cmd.Execute()
}
