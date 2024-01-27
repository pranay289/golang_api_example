package user_recommendation

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// recomm
func GetAllRecipies() (RecipeModel, error) {
	var recipes RecipeModel
	baseUrl := os.Getenv("baseUrl")
	response, responseErr := http.Get(baseUrl + "alldata")

	// handling error
	if responseErr != nil {
		log.Fatal(responseErr)
		return nil, responseErr
	}

	defer response.Body.Close()
	// To read bytes
	body, readError := io.ReadAll(response.Body) // response body is []byte

	if readError != nil {
		return nil, readError
	}

	// To decode json and to handle errors
	if err := json.Unmarshal(body, &recipes); err != nil {
		return nil, err
	}

	// myString := string(body)

	// fmt.Println(prettylogger.JsonPrettyPrint(string(body)))

	return recipes, nil

}
