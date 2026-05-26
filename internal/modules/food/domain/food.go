package domain

type Food struct {
	ID        string
	ImageURL  string
	Nutrients Nutrients
}

type Nutrients struct {
	Calories  string
	Protein   string
	Carbs     string
	Fat       string
}
