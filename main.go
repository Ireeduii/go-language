package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// Туршилтын санах ой (In-memory database)
var users = []User{
	{ID: 1, Name: "Bat-Ireedui", Role: "Developer"},
	{ID: 2, Name: "Anand", Role: "Designer"},
}

func main() {
	// API routes
	http.HandleFunc("/api/users", handleUsers)

	fmt.Println("🚀 Go Server ажиллаж байна: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Сервер асахад алдаа гарлаа:", err)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. GET хүсэлт ирвэл жагсаалтаа буцаана
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(users)
		return
	}

	// 2. POST хүсэлт ирвэл шинэ хэрэглэгч нэмнэ
	if r.Method == http.MethodPost {
		var newUser User

		// Хэрэглэгчийн явуулсан JSON-ийг уншиж newUser уруу хөрвүүлнэ
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Буруу JSON өгөгдөл байна", http.StatusBadRequest)
			return
		}

		// Шинэ ID олгоод array руугаа нэмнэ
		newUser.ID = len(users) + 1
		users = append(users, newUser)

		// Амжилттай нэмэгдсэнийг буцаана (Status 201 Created)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
		return
	}

	// GET болон POST-оос бусад хүсэлтэнд алдаа буцаана
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}