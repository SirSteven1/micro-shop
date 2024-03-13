// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	UserName  string `json:"UserName"`  // 用户名
	PassWord  string `json:"PassWord"`  // 用户密码，MD5加密
	UserNick  string `json:"UserNick"`  // 用户昵称
	UserFace  string `json:"UserFace"`  // 用户头像地址
	UserSex   int64  `json:"UserSex"`   // 用户性别：0男，1女，2保密
	UserEmail string `json:"UserEmail"` // 用户邮箱
}

type LoginReq struct {
	UserName string `json:"UserName"` // 用户名
	PassWord string `json:"PassWord"` // 用户密码，MD5加密
}

type UserInfoReq struct {
	//UserIdentity string `json:"UserIdentity"` // 用户唯一标识
}

type UserInfoResply struct {
	Code    int64         `json:"Code"`
	Message string        `json:"Message"`
	Data    *UserInfoItem `json:"Data"`
}

type CommonResply struct {
	Code    int64  `json:"Code"`
	Message string `json:"Message"`
	Data    interface{} `json:"Data"`
}

type UserInfoItem struct {
	UserIdentity string `json:"UserIdentity"` // 用户唯一表哦是
	UserName     string `json:"UserName"`     // 用户名
	UserNick     string `json:"UserNick"`     // 用户昵称
	UserFace     string `json:"UserFace"`     // 用户头像地址
	UserSex      int64  `json:"UserSex"`      // 用户性别：0男，1女，2保密
	UserEmail    string `json:"UserEmail"`    // 用户邮箱
	UserPhone    string `json:"UserPhone"`    // 用户手机号
}