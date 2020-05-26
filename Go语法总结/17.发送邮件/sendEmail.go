package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	user := "fuqiang@mgv.ai"
	password := "password"
	host := "smtp.exmail.qq.com:25"
	to := "1078649116@qq.com"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: discount Gophers!" + "\r\n\r\n This is the email body.\r\n")
	auth := smtp.PlainAuth("", user, password, "smtp.exmail.qq.com")
	err := smtp.SendMail(host, auth, user, []string{"1078649116@qq.com"}, msg)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
