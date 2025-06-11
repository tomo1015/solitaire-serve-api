package leaderboard

import (
	"solitaire-serve-api/storage"
	"sort"
)

type Entry struct {
	ID        string
	Resources int
}

func GetLeaderboard() []Entry {
	entries := []Entry{}
	for id, p := range storage.Players {
		entries = append(entries, Entry{ID: id, Resources: p.Resources})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Resources > entries[j].Resources
	})

	return entries
}
