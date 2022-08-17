package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type Health struct {
	Health  string
	Version int
}

// 4
func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	// 들어오는 요청의 Response Header에 Content-Type을 추가
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		// 전달 받은 http.Handler를 호출한다.
		next.ServeHTTP(rw, r)
	})
}

func main() {
	// 1
	mux := http.NewServeMux()
	now := time.Now()
	// 2
	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		custom := now.Format("2006-01-02 15:04:05")
		fmt.Printf(custom + " " + r.Method + " " + r.UserAgent() + " " + r.RequestURI + "\n")
		switch r.Method {
		case http.MethodGet: // 조회
			json.NewEncoder(wr).Encode(users)
		case http.MethodPost: // 등록
			var user User
			json.NewDecoder(r.Body).Decode(&user)

			users[user.Email] = &user

			json.NewEncoder(wr).Encode(user)
		}
	})

	healthHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		custom := now.Format("2006-01-02 15:04:05")
		fmt.Printf(custom + " " + r.Method + " " + r.UserAgent() + " " + r.RequestURI + "\n")

		u := Health{
			Health:  "up",
			Version: 4,
		}

		json.NewEncoder(wr).Encode(u)
	})

	// 3
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	mux.Handle("/health", jsonContentTypeMiddleware(healthHandler))
	http.ListenAndServe(":8080", mux)
}
