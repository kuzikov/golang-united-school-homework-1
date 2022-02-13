package solution

import "github.com/kyokomi/emoji"

// GetMessage responses with emoji string.
func GetMessage() string {
	return emoji.Sprint("Hello 🗺️ !")
}
