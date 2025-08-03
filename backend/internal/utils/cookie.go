package utils

import (
	"net/http"
    "time"
)
//		function set cookie
func SetCookie(w http.ResponseWriter, token string){
	cookie := &http.Cookie{
        Name:     "jwt",
        Value:    token,
        Path:     "/",
        MaxAge:   60 * 20,
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode, 
    }
    http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request, name string) (string, error){
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ClearCookie(w http.ResponseWriter){
    cookie := &http.Cookie{
        Name:     "jwt",
        Value:    "",
        Path:     "/",
        Expires: time.Now().Add(-24 * time.Hour),
        MaxAge:   -1,
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode, 
    }
    http.SetCookie(w, cookie)
}
