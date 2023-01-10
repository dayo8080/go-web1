package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("The server is running on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 6)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page, the addition of 2 + 6 = %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 100.0, 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		err := errors.New("CANNOT DIVIDE BY ZERO")
		return 0, err
	}
	result := x / y

	return result, nil

}
