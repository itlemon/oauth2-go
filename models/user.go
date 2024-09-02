package models

type UserDetails struct {
	// 用户标识
	UserId int

	// 用户名 唯一
	Username string

	// 用户密码
	Password string

	// 权限列表
	Authorities []string
}
