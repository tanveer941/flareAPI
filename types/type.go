package types

type ApiError struct {
	Err    string
	Status int
}

func (ae *ApiError) Error() string {
	return ae.Err
}

type User struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
