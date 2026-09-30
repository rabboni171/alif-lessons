package main

import "fmt"

type Song struct {
	Title    string
	Duration int
}

type Playlist struct {
	songs []Song
}

func (p *Playlist) Add(s Song) {
	p.songs = append(p.songs, s)
}

func (p Playlist) TotalDuration() int {
	total := 0
	for _, s := range p.songs {
		total += s.Duration
	}
	return total
}

func (p Playlist) Longest() Song {
	longest := p.songs[0]
	for _, s := range p.songs {
		if s.Duration > longest.Duration {
			longest = s
		}
	}
	return longest
}

func (p Playlist) String() string {
	result := ""
	for _, s := range p.songs {
		result += fmt.Sprintf("%s (%d сек)\n", s.Title, s.Duration)
	}
	return result
}

func main() {
	playlist := &Playlist{}

	playlist.Add(Song{Title: "Skyfall", Duration: 286})
	playlist.Add(Song{Title: "Yesterday", Duration: 125})
	playlist.Add(Song{Title: "Bohemian Rhapsody", Duration: 355})
	playlist.Add(Song{Title: "Hotel California", Duration: 391})
	playlist.Add(Song{Title: "Imagine", Duration: 183})

	fmt.Print(playlist)
	fmt.Println("Общая длительность:", playlist.TotalDuration(), "сек")
	fmt.Println("Самый длинный трек:", playlist.Longest().Title)
}
