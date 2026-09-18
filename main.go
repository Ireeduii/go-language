package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PHP API-аас ирэх Task-ийн бүтэц
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/api/php-tasks", getTasksFromPHP)

	fmt.Println("🚀 Go Server ажиллаж байна: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func getTasksFromPHP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. PHP API руу GET хүсэлт илгээнэ
	resp, err := http.Get("http://localhost/php-api/tasks.php")
	if err != nil {
		http.Error(w, "PHP API-тай холбогдож чадсангүй", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 2. Ирсэн JSON-ийг Go-ийн struct руу уншина
	var tasks []Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, "JSON уншихад алдаа гарлаа", http.StatusInternalServerError)
		return
	}

	// 3. Уншсан датагаа буцаана
	json.NewEncoder(w).Encode(tasks)
}