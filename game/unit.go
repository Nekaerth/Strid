package game

type Unit struct {
	position *Field
}

type Human struct {
	Health int
}

type Infantry struct {
	Human
}

type Cavalry struct {
	Human
}

type Archery struct {
	Human
}
