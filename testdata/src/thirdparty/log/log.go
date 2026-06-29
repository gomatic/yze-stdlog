// Package log is a third-party package whose import path ends in "log". It exists
// to prove the analyzer matches the exact path "log", not a suffix or substring.
package log

// Println is a third-party stand-in unrelated to the standard log package.
func Println(...any) {}
