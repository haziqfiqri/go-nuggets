package methods

import (
	"fmt"
	"time"
)

type Label struct {
	Name    string
	Artists []*Artist
}

type Artist struct {
	Name      string
	Country   string
	Active    bool
	DebutDate time.Time
	Songs     []*Song
}

type Song struct {
	Name        string
	Duration    time.Duration
	ReleaseDate time.Time
}

func NewArtist(name, country string, active bool) *Artist {
	return &Artist{
		Name:      name,
		Country:   country,
		Active:    active,
		DebutDate: time.Now(),
	}
}

func NewLabel(name string) *Label {
	return &Label{Name: name}
}

func (l *Label) Sign(a *Artist) {
	l.Artists = append(l.Artists, a)
}

func Methods() {
	label := NewLabel("Universal Music")
	fmt.Println("label", label)

	artist := NewArtist("Taylor Swift", "USA", true)
	fmt.Println("artist", artist)

	label.Sign(artist)
	fmt.Println("label", label)
}
