package dto

type CreateCourierInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateCourierOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
