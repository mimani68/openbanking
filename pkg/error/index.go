package error

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}
