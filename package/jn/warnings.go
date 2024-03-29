package jn

import "fmt"

var Warnings = map[string]string{
	`doc_ignored`:         `documentation is ignored because object isn't supports documentations`,
	`exist_undefined_doc`: `source code has undefined documentations (some documentations isn't document anything)`,
}

func GetWarning(key string, args ...any) string {
	return fmt.Sprintf(Warnings[key], args...)
}
