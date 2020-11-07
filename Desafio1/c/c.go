package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

type Result struct {
	Status string
}

var coupons Coupons

func main() {

	coupon := Coupon{
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

// func callMymicroservice() Result {

// 	retryClient := retryablehttp.NewClient()
// 	retryClient.RetryMax = 5

// 	response, err := retryClient.Get("https://avancadev-mymicroservice.herokuapp.com/MyService")
// 	//response, err := retryClient.Get("https://localhost:5001/MyService")

// 	if err != nil {
// 		result := Result{Status: "Servidor fora do ar!"}
// 		return result
// 	}

// 	defer response.Body.Close()

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	result := Result{}
// 	json.Unmarshal(responseData, &result)

// 	return result
// }

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	valid := coupons.Check(coupon)

	result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)

	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}
