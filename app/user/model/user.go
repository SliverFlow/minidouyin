package model

type (
	userRepo interface {
	}

	// User 用户表
	User struct {
	}

	// defaultUserModel 默认用户表模型
	defaultUserModel struct {
	}
)

func newUserModel() *defaultUserModel {
	return nil
}

func (u *User) TableName() string {
	return "d_user"
}
