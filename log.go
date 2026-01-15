package geonames

import "log"

// Panicf logs a formatted message and then panics
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}
