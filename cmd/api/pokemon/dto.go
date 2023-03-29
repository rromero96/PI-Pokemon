package pokemon

type Pokemon struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	Image   string `json:"image"`
	Speed   int    `json:"speed"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Created bool   `json:"created"`
	Types   []Type `json:"types"`
}

type Type struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
