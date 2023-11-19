package functions

import (
	utils "love/leno"
	"strconv"
)

func GetProfiles(browserPath string) []string {
	var profiles []string
	if utils.CheckFileExist(browserPath + "\\Default") {
		profiles = append(profiles, browserPath+"\\Default")
	}

	for i := 1; i < 6; i++ {
		if utils.CheckFileExist(browserPath + "\\Profile " + strconv.Itoa(i)) {
			profiles = append(profiles, browserPath+"\\Profile "+strconv.Itoa(i))
		}
	}
	return profiles
}