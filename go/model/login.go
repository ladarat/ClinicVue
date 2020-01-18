package	model

// LoginRequest validate requidle
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse modek
type LoginResponse struct {
	UserID string `json:"userID"`
}
