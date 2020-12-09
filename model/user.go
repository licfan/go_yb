package model

type User struct {
	ID            string `json:"id" `
	UserName      string `json:"id" `
	NickName      string `json:"id" `
	NickNameTmp   string `json:"id" `
	Email         string `json:"id" `
	Avatar        string `json:"id" `
	AvatarTmp     string `json:"id" `
	Sex           string `json:"id" `
	Password      string `json:"id" `
	Salt          string `json:"id" `
	Phone         string `json:"id" `
	Province      string `json:"id" `
	City          string `json:"id" `
	Constellation string `json:"id" `
	RegIp         int    `json:"id" `
	RegSite       string `json:"id" `
	Type          int8   `json:"id" `
	Status        int8   `json:"id" `
	RegTime       int    `json:"id" `
	UpdateTime    int    `json:"id" `
	LastLoginTime int    `json:"id" `
	LastLoginIp   int    `json:"id" `
}
