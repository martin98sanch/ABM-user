package user

type (
	Users []DTO
	DTO   struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
)

func (dto DTO) Validate() error {
	if dto.Name == "" || dto.Username == "" || dto.Password == "" || dto.Age == 0 {
		return ErrInvalidBody
	}
	return nil
}
