package model

// User is user interface which need to be implemented
type User interface {
	GetUserLoginID() string
}

var defaultUser User

// AddUserModel will add the user model that implements the User interface
func AddUserModel(user User) {
	defaultUser = user
}

// defaultUserImp is default user model implementation
// if AddUserModel is not called to implement the User interface then defaultUserImp will be used
type defaultUserImp struct {
	LoginID string
}

func (u *defaultUserImp) GetUserLoginID() string {
	return u.LoginID
}
