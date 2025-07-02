// main.go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Todo merepresentasikan data dari sebuah to-do item
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// "Database" in-memory kita
var todos = []Todo{
	{ID: 1, Title: "Belajar Go", Completed: true},
	{ID: 2, Title: "Belajar React", Completed: false},
	{ID: 3, Title: "Buat Aplikasi CRUD", Completed: false},
}

// Variabel untuk auto-increment ID
var nextID = 4

func main() {
	router := gin.Default()

	// Setup CORS Middleware
	// Ini mengizinkan semua origin. Untuk produksi, Anda harus lebih spesifik.
	router.Use(cors.Default())

	// Definisikan rute untuk setiap operasi CRUD
	// READ: Mendapatkan semua to-do
	router.GET("/todos", getTodos)
	// CREATE: Menambahkan to-do baru
	router.POST("/todos", addTodo)
	// UPDATE: Mengubah status completed sebuah to-do
	router.PATCH("/todos/:id", toggleTodoStatus)
	// DELETE: Menghapus sebuah to-do
	router.DELETE("/todos/:id", deleteTodo)

	// Jalankan server di port 8080
	router.Run("localhost:8080")
}

// getTodos: Handler untuk mengambil semua data to-do
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// addTodo: Handler untuk menambahkan to-do baru
func addTodo(c *gin.Context) {
	var newTodo Todo

	// Bind JSON yang diterima ke struct newTodo
	if err := c.BindJSON(&newTodo); err != nil {
		return // Jika error, hentikan proses
	}

	// Beri ID baru dan tambahkan ke slice 'todos'
	newTodo.ID = nextID
	nextID++
	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

// toggleTodoStatus: Handler untuk mengubah status 'completed'
func toggleTodoStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid"})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !todos[i].Completed
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "To-do tidak ditemukan"})
}

// deleteTodo: Handler untuk menghapus to-do berdasarkan ID
func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid"})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			// Hapus item dari slice
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "To-do berhasil dihapus"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "To-do tidak ditemukan"})
}