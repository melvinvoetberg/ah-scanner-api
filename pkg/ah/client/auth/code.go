package ahauth

import (
	"io/ioutil"
  "strings"
  "unicode"
)

func GetCode() (code string, err error) {
  b, err := ioutil.ReadFile("code.txt")

  if err != nil {
    return "", err
  }

  return stripSpaces(string(b)), nil
}

// https://stackoverflow.com/a/32082217
func stripSpaces(str string) string {
  return strings.Map(func(r rune) rune {
    if unicode.IsSpace(r) {
      return -1
    }
    return r
  }, str)
}
