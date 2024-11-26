package utilities

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

func GenerateToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	token := make([]byte, 8)
	for i := range token {
		token[i] = letters[r.Intn(len(letters))]
	}
	return string(token)
}

func SendTokenEmail(to string, token string) error {

	corr := os.Getenv("CORREO")
	pass := os.Getenv("PASS_CORREO")

	m := gomail.NewMessage()
	m.SetHeader("From", corr) // Cambiar a tu correo
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Tu Token de Verificación")
	m.SetBody("text/html", fmt.Sprintf("Este es tu token: <b>%s</b>", token))

	d := gomail.NewDialer("smtp.gmail.com", 587, corr, pass) // Cambiar con tus credenciales

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
