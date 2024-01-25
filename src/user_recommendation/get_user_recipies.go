package user_recommendation

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

func GetUserRecipies(title string) (RecipeModel, error) {

	var recipes RecipeModel
	// This will provide baseUrl
	baseUrl := os.Getenv("baseUrl")

	titleArray := []string{title}

	requestMap := url.Values{
		"title": titleArray,
	}

	// bytes.NewBuffer(jsonReq)
	// jsonReq, err := json.Marshal(requestMap)

	resp, errrr := http.PostForm(baseUrl+"recomm", requestMap)

	if errrr != nil {
		return nil, errrr
	}

	defer resp.Body.Close()

	data, errorr := io.ReadAll(resp.Body)
	// To decode json and to handle errors
	if err := json.Unmarshal(data, &recipes); err != nil {
		// return nil, err
	}
	if errorr != nil {
		return nil, errorr
	}

	return recipes, nil
	// f, _ := json.Marshal(recipes[0])
	// fmt.Println(prettylogger.JsonPrettyPrint(string(f)))

}
