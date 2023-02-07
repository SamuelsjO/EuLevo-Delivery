

type Route struct {
	ID string
	ClientID string
	Positions []Positions
}

type Positions struct {
	Lat float64
	Long float64
}

