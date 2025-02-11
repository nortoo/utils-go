package http

import "fmt"

// BuildContentDisposition returns the content disposition string of a file.
func BuildContentDisposition(filename string) string {
	return fmt.Sprintf("filename=%s", filename)
}
