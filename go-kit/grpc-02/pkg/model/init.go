package model

var UM *UserManager

func init() {
	UM = &UserManager{
		Users: []*User{
			&User{
				Name:    "liyanhong",
				Title:   "CEOx",
				Company: "baidu",
			},
			&User{
				Name:    "mayun",
				Title:   "CMOx",
				Company: "ali",
			},
			&User{
				Name:    "pony",
				Title:   "CTOx",
				Company: "tecent",
			},
		},
	}
}
