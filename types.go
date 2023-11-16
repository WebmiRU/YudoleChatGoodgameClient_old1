package main

type Config struct {
	Servers struct {
		Host struct {
			Http struct {
				Address string `json:"address"`
				Port    int16  `json:"port"`
			}
			Server struct {
				Address string `json:"address"`
				Port    int16  `json:"port"`
			}
		} `json:"host"`
		Twitch struct {
			Address  string   `json:"address"`
			Port     int16    `json:"port"`
			Login    string   `json:"login"`
			Password string   `json:"password"`
			Channels []string `json:"channels"`
		} `json:"twitch"`
	} `json:"servers"`
}

type TypeMeta struct {
	Badges map[string]string `json:"badges"` // key => Image URL
}

type TypeUser struct {
	Id       string   `json:"id"`
	Nickname string   `json:"nickname"`
	Login    string   `json:"login"`
	Meta     TypeMeta `json:"meta"`
}

type GoodgameMessageData struct {
	ChannelId string `json:"channel_id"`
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Color     string `json:"color"`
	Icon      string `json:"icon"`
	Role      string `json:"role"`
	Mobile    int    `json:"mobile"`
	IsStatus  int    `json:"isStatus"`
	MessageId int64  `json:"message_id"`
	Timestamp int    `json:"timestamp"`
	Text      string `json:"text"`
	Hidden    int    `json:"hidden"`
	Reload    bool   `json:"reload"`
}

type GoodgameMessage struct {
	Type string              `json:"type"`
	Data GoodgameMessageData `json:"data"`
}

type GoodgameJoinRequestData struct {
	ChannelId string `json:"channel_id"`
	Hidden    int    `json:"hidden"`
	Mobile    bool   `json:"mobile"`
	Reload    bool   `json:"reload"`
}

type GoodgameJoinRequest struct {
	Type string                  `json:"type"`
	Data GoodgameJoinRequestData `json:"data"`
}
