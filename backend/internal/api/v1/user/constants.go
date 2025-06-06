package user

type RoleType string

const (
	Regular   RoleType = "regular"
	Admin     RoleType = "admin"
	SuperUser RoleType = "superUser"
)

type AuthMethod string

const (
	Credentials AuthMethod = "credentials"
	Google      AuthMethod = "google"
	Github      AuthMethod = "github"
)
