package main

import (
	logg "log"
	"os"
)

var log = logg.New(os.Stderr, "log ", logg.Llongfile)
