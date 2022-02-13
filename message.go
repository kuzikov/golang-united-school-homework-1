package solution

import "github.com/kyokomi/emoji"

// GetMessage responses with emoji containing string.
func GetMessage() string {
	return emoji.Sprint("Hello 🗺️!")
}
