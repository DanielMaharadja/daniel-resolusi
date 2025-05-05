package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Konten HTML untuk tampilan depan
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Jual Ceramic</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					margin: 0;
					padding: 0;
					text-align: center;
				}
				header {
					background-color: #333;
					color: white;
					padding: 20px;
				}
				main {
					margin: 40px;
				}
				footer {
					background-color: #f1f1f1;
					padding: 10px;
					position: fixed;
					width: 100%;
					bottom: 0;
				}
			</style>
		</head>
		<body>
			<header>
				<h1>Selamat Datang di Jual Ceramic</h1>
			</header>
			<main>
				<p>Kami menyediakan keramik berkualitas tinggi untuk rumah dan proyek Anda.</p>
				<p><a href="/produk">Lihat Produk</a></p>
			</main>
			<footer>
				<p>&copy; 2025 Jual Ceramic. All rights reserved.</p>
			</footer>
		</body>
		</html>
	`
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server berjalan di http://localhost:8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
