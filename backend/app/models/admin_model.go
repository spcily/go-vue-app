package models

type Admin struct {
	Username string
	Password string
}

var Admins []Admin

func AddAdmin(admin Admin) {
	Admins = append(Admins, admin)
}
