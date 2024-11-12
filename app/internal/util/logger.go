package utils

import (
	"log"
	"os"
)

// Logger is a simple logger
var Logger = log.New(os.Stdout, "", log.LstdFlags)
