package icon

import _ "embed"

//go:embed default.ico
var defaultIconData []byte

//go:embed blackAndWhite.ico
var blackAndWhiteIconData []byte

func GetIconData(iconPath string) []byte {
	switch iconPath {
	case "default.ico":
		return defaultIconData
	case "blackAndWhite.ico":
		return blackAndWhiteIconData
	default:
		return nil
	}
}
