package entity


type UserSession struct {
	openId string
}

func NewUserSession(openId string) *UserSession {
	return &UserSession{
		openId:openId,
	}
}

func (session *UserSession)OpenId()string{
	return session.openId
}

func (session *UserSession)Verify()bool{
	// todo 这里实现用户身份校验
	return true
}