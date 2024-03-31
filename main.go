package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Card struct {
	Number string `json:"number"`
}

func detectCardFlag(number string) string {
	switch {
	case strings.HasPrefix(number, "1"):
		return "Visa"
	case strings.HasPrefix(number, "2"):
		return "Mastercard"
	case strings.HasPrefix(number, "3"):
		return "American Express"
	default:
		return "Unknow"
	}
}

func validateCard(number string) (bool, string) {
	number = strings.ReplaceAll(number, " ", "")
	if _, err := strconv.Atoi(number); err != nil {
		return false, ""
	}

	sum := 0
	doubleNext := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')

		if doubleNext {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		doubleNext = !doubleNext
	}

	isValid := sum%10 == 0
	cardFlag := detectCardFlag(number)

	return isValid, cardFlag
}

func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	var card Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valid, flag := validateCard(card.Number)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"Valid": valid, "Flag": flag})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/validate_card", validateCardHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", router))
}
