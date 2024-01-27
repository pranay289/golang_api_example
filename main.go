package main

import (
	"encoding/json"
	"fmt"
	"golang_apis/src/user_recommendation"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// This is use to load env file to get environment variables
	godotenv.Load(".env")
	router := mux.NewRouter()

	// To get all recipies
	router.HandleFunc("/get-all-recipies", getAllrecipies).Methods("GET")

	router.HandleFunc("/get-user-recipies", getUserRecipies).Methods("POST")

	// router := gin.Default()

	// To get all recipies
	// router.GET("/get-all-recipies", getAllrecipies)
	// newRcipes, respError := user_recommendation.GetRecommendations()

	// To handle 404 custom errors
	router.NotFoundHandler = http.HandlerFunc(noRouteFound)

	fmt.Println("server is listining....")
	log.Fatal(http.ListenAndServe(":8080", router))

	// router.Run("localhost:9000")
	// Entry point our app

	// To get all recipies
	// _, err := user_recommendation.GetRecommendations()
	// user_recommendation.GetUserRecipies("dahi sabudana")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, er := json.Marshal(recipesData)
	// if er != nil {

	// }
	// log.Fatal()
}

func getUserRecipies(w http.ResponseWriter, r *http.Request) {
	var params = r.Body
	data, _ := io.ReadAll(params)

	// creating new slice of map
	userRequest := make(map[string]interface{})

	if anyError := json.Unmarshal(data, &userRequest); anyError != nil {
		// TODO
		w.WriteHeader(500)
	}
	if len(userRequest) == 0 {
		w.WriteHeader(400)
	}

	if _, isExist := userRequest["title"]; isExist {
		// TODO
		w.Write([]byte("Ok"))
	}

	// To get all recipies
	// newRcipes, respError := user_recommendation.GetUserRecipies()
}

// To get all recipies
func getAllrecipies(response http.ResponseWriter, request *http.Request) {

	// To get all recipies
	newRcipes, respError := user_recommendation.GetAllRecipies()

	// To handle error
	if respError != nil {
		//
		log.Fatal(respError)
		response.Write([]byte(respError.Error()))
		response.WriteHeader(500)

	}

	response.Header().Set("Content-Type", "application/json")
	// TODO: not in used

	/*

	   To convert model to json
	   	data, err := json.Marshal(newRcipes)

	   	// To handle error
	   	if err != nil {
	   		log.Fatal(respError)
	   		response.Write([]byte(respError.Error()))
	   		response.WriteHeader(500)
	   	}

	*/

	// custom response model for sending the data
	jsonFormat := map[string]interface{}{
		"result": "success",
		"status": 200,
		"data":   newRcipes,
	}

	f, err := json.Marshal(jsonFormat)

	if err != nil {
		log.Fatal(respError)
		response.Write([]byte(respError.Error()))
		response.WriteHeader(500)
	}

	// sending response
	response.Write(f)

}

// To handle not found route error
func noRouteFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(404)
	// custom response model for sending the data
	jsonFormat := map[string]interface{}{
		"result": "failed",
		"status": 404,
		"data":   "No route found please, specify correct route",
	}

	f, _ := json.Marshal(jsonFormat)
	w.Write(f)
}
