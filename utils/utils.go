package utils

import (
	"backend/types"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func ParsePayload(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body empty")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteResponse(w, status, map[string]string{"error": err.Error()})
}

func GetEmptyJSONFields(payload any) []string {
	emptyFields := []string{}
	val := reflect.ValueOf(payload)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("json")
		if fieldName == "" {
			fieldName = typ.Field(i).Name
		}
		fieldName = strings.Split(fieldName, ",")[0]

		if field.IsZero() {
			emptyFields = append(emptyFields, fieldName)
		}
	}
	return emptyFields
}

func CalculatePoints(payload types.ReceiptPayload) uint64 {
	var points uint64 = 0
	// Get points for retailer name
	points += CalculateNamePoints(payload.Retailer)
	log.Println("Points from name:", points)
	// Get points for purchase date
	purchaseDate := payload.PurchaseDate
	day := purchaseDate[strings.LastIndex(purchaseDate, "-")+1:]
	day_num, err := strconv.Atoi(day)
	if err != nil {
		log.Fatal(err)
	}
	if day_num%2 != 0 {
		points += 6
	}
	log.Println("Points from date:", points)
	// Get points for purchase time on the day
	purchaseTime := payload.PurchaseTime
	if idx := strings.IndexByte(purchaseTime, ':'); idx >= 0 {
		hours := purchaseTime[:idx]
		minutes := purchaseTime[idx+1:]
		hour_limit, err := strconv.Atoi(hours)
		if err != nil {
			log.Fatal("time parsing failed")
		}
		hour_limit = hour_limit * 100
		minute_limit, err := strconv.Atoi(minutes)
		if err != nil {
			log.Fatal("time parsing failed")
		}
		hour_limit += minute_limit
		if 1400 < hour_limit && hour_limit < 1600 {
			points += 10
		}
	}
	log.Println("Points from time:", points)
	// Get points for number of items
	// log.Println(len(payload.Items) / 2)
	points += uint64((len(payload.Items))/2) * 5
	log.Println("Points from item #:", points)
	for _, item := range payload.Items {
		points += CalculateItemPoints(item)
	}
	log.Println("Points from item:", points)
	total, err := strconv.ParseFloat(payload.Total, 64)
	if err != nil {
		log.Fatal(err)
	}
	if math.Floor(total) == total {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}
	return points
}

func CalculateNamePoints(retailer string) uint64 {
	var retail_points uint64 = 0
	for _, r := range retailer {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			retail_points += 1
		}
	}
	return retail_points
}

func CalculateItemPoints(item types.Item) uint64 {
	description := item.ShortDescription
	description = strings.TrimSpace(description)
	price, err := strconv.ParseFloat(item.Price, 64)
	if err != nil {
		log.Fatal("price parsing in item failed")
	}
	var item_points uint64 = 0
	if len(description)%3 == 0 {
		println(description, len(description))
		item_points = uint64(math.Ceil(price * 0.2))
	}
	return item_points
}
