package models

type ClientDetails struct {
	// 客户端标识
	ClientId string

	// 客户端密钥
	ClientSecret string

	// 访问令牌的有效时间，秒
	AccessTokenValiditySeconds int

	// 刷新令牌的有效时间，秒
	RefreshTokenValiditySeconds int

	// 重定向地址，用于授权码类型
	RegisteredRedirectUri string

	// 可用的授权类型
	AuthorizedGrantTypes []string
}
