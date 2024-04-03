package models

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

var counter int

func Cookies(w http.ResponseWriter) {
	uuid := GenerateUUID()

	http.SetCookie(w, &http.Cookie{
		Name:  "user",
		Value: uuid,
	})
}

func GetCookies(r *http.Request) string{
	cookie, err := r.Cookie("user")
	if err != nil {
		fmt.Println("l'utilisateur n'est pas connecté")
	}
	value := cookie.Value
	return value
}

func GenerateUUID() string {
	counter++

	data := fmt.Sprintf("%d%d", time.Now().UnixNano(), counter)

	hash := sha1.New()
	hash.Write([]byte(data))
	uuidHash := hash.Sum(nil)

	uuid := hex.EncodeToString(uuidHash)

	return uuid[:16]
}
