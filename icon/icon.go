package icon

import (
	"log"
	"os"
)

func GetIconData(iconPath string) []byte {
	data, err := os.ReadFile(iconPath)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func DefaultIconPath() string {
	return "./icon/default.ico"
}

func BlackAndWhiteIconPath() string {
	return "./icon/black_and_white.ico"
}
