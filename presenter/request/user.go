package request

type NewUser struct {
	UID         string  `json:"uid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
