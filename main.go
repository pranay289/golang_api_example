package main

import (
	"github.com/joho/godotenv"
)

func main() {

	// This is use to load env file to get environment variables
	godotenv.Load(".env")

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
