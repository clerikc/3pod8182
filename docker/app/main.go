package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// HTML с надписью "Привет медведь" по центру, шрифт 20px
		podName := os.Getenv("POD_NAME")
		html := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>K8s Demo</title>
				<style>
					body {
						display: flex;
						justify-content: center;
						align-items: center;
						height: 100vh;
						margin: 0;
						font-size: 20px;
					}
				</style>
			</head>
			<body>
				<div>Привет медведь (from %s)</div>
			</body>
			</html>
		`, podName)
		fmt.Fprint(w, html)
	})

	http.ListenAndServe(":8080", nil)
}
