package google

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type ResponseGoogle struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"verifiedEmail"`
	Name     string `json:"name"`

	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Picture    string `json:"picture"`
}
