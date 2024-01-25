package user_recommendation

type RecipeModel []RecipeModelElement

type RecipeModelElement struct {
	RecipeID     int64      `json:"RecipeID"`
	Recipies     string     `json:"recipies"`
	ImageLink    string     `json:"ImageLink"`
	Description  string     `json:"Description"`
	Tags         string     `json:"Tags"`
	Nutritions   Nutritions `json:"Nutritions"`
	By           string     `json:"By"`
	RequiredTime int64      `json:"RequiredTime"`
	Servings     int64      `json:"Servings"`
	Ingredients  string     `json:"Ingredients"`
	Directions   string     `json:"Directions"`
}

type Nutritions struct {
	Energy     string `json:"energy"`
	Protein    string `json:"protein"`
	Carbs      string `json:"carbs"`
	Fiber      string `json:"fiber"`
	Fat        string `json:"fat"`
	Cholestrol string `json:"cholestrol"`
	Sodium     string `json:"sodium"`
}
