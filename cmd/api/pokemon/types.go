package pokemon

type Pokemon struct {
	ID      string
	Name    string
	HP      int
	Attack  int
	Defense int
	Image   string
	Speed   int
	Height  int
	Weight  int
	Created bool
	Types   []Type
}

type Type struct {
	ID   string
	Name string
}
