package functions

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
)





func SendEmail(from, password, smtpServer, smtpPort, to, subject, body string, fileNames ...string) error {
	auth := smtp.PlainAuth("", from, password, smtpServer)

	mimeParts := []string{}
	mimeParts = append(mimeParts, "From: "+from)
	mimeParts = append(mimeParts, "To: "+to)
	mimeParts = append(mimeParts, "Subject: "+subject)
	mimeParts = append(mimeParts, "MIME-Version: 1.0")
	mimeParts = append(mimeParts, "Content-Type: multipart/mixed; boundary=boundarystring")
	mimeParts = append(mimeParts, "")

	// Add the body of the email
	mimeParts = append(mimeParts, "--boundarystring")
	mimeParts = append(mimeParts, "Content-Type: text/plain; charset=UTF-8")
	mimeParts = append(mimeParts, "")
	mimeParts = append(mimeParts, body)

	// Add attachments
	for _, fileName := range fileNames {
		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
			return err
		}

		base64Content := base64.StdEncoding.EncodeToString(fileContent)

		mimeParts = append(mimeParts, "--boundarystring")
		mimeParts = append(mimeParts, "Content-Type: application/octet-stream")
		mimeParts = append(mimeParts, "Content-Disposition: attachment; filename=\""+filepath.Base(fileName)+"\"")
		mimeParts = append(mimeParts, "Content-Transfer-Encoding: base64")
		mimeParts = append(mimeParts, "")
		mimeParts = append(mimeParts, base64Content)
	}

	// End of MIME message
	mimeParts = append(mimeParts, "--boundarystring--")

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(strings.Join(mimeParts, "\r\n")))
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully.")
	return nil
}

