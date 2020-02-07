package contract

// EntryRequetContext defines request body to save an entry
type EntryRequetContext struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
