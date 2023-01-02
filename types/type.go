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

type Configuration struct {
	PortAddress    string
	Environment    string
	DatabaseConfig struct {
		Dsn string
	}
}
