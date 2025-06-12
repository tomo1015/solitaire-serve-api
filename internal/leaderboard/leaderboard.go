package leaderboard

type Entry struct {
	ID        string
	Resources int
}

// func GetLeaderboard() []Entry {
// 	// entries := []Entry{}
// 	// for id, p := range storage.Players {
// 	// 	//TODO：リーダーボードは村の発展力に変更
// 	// 	entries = append(entries, Entry{ID: id, Resources: p.Soldiers})
// 	// }

// 	// sort.Slice(entries, func(i, j int) bool {
// 	// 	return entries[i].Resources > entries[j].Resources
// 	// })

// 	// return entries
// }
