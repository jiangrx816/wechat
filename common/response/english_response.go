package response

type BYJSONData struct {
	Ret  int `json:"ret"`
	Data struct {
		Ent struct {
			Items []struct {
				Bookid int64  `json:"bookid"`
				Level  int    `json:"level"`
				Title  string `json:"title"`
				Cover  struct {
					Tiny   string `json:"tiny"`
					Origin string `json:"origin"`
					W      int    `json:"w"`
					H      int    `json:"h"`
				} `json:"cover"`
				Ct                int         `json:"ct"`
				Ut                int         `json:"ut"`
				State             int         `json:"state"`
				Domain            string      `json:"domain"`
				Downloads         int         `json:"downloads"`
				Resid             int         `json:"resid"`
				Screen            int         `json:"screen"`
				Top               int         `json:"top"`
				Flag              int         `json:"flag"`
				Playcount         int         `json:"playcount"`
				Pagecount         int         `json:"pagecount"`
				Score             int         `json:"score"`
				Version           int         `json:"version"`
				Lock              int         `json:"lock"`
				Recordstate       int         `json:"recordstate"`
				Topicstate        int         `json:"topicstate"`
				Vocabulary        int         `json:"vocabulary"`
				Difficulty        int         `json:"difficulty"`
				Isvip             bool        `json:"isvip"`
				Isstoproll        bool        `json:"isstoproll"`
				Topicid           int         `json:"topicid"`
				Topicbookindex    int         `json:"topicbookindex"`
				Vipbubble         bool        `json:"vipbubble"`
				Newbubble         bool        `json:"newbubble"`
				Recordbubble      bool        `json:"recordbubble"`
				Onlinetime        int         `json:"onlinetime"`
				Readbubble        bool        `json:"readbubble"`
				Themeinfos        interface{} `json:"themeinfos"`
				Notinpicgallery   bool        `json:"notinpicgallery"`
				Tags              []string    `json:"tags"`
				Words             []string    `json:"words"`
				Introduction      string      `json:"introduction"`
				Voicetype         int         `json:"voicetype"`
				Pressid           int         `json:"pressid"`
				Booktype          int         `json:"booktype"`
				Illustrator       string      `json:"illustrator"`
				Translator        string      `json:"translator"`
				Writer            string      `json:"writer"`
				Bgm               string      `json:"bgm"`
				Firstclassifyid   int         `json:"firstclassifyid"`
				Secondclassifyid  int         `json:"secondclassifyid"`
				Paytype           int         `json:"paytype"`
				Sort              int         `json:"sort"`
				Levellist         interface{} `json:"levellist"`
				Createrid         int         `json:"createrid"`
				Lastupdateid      int         `json:"lastupdateid"`
				Avescore          int         `json:"avescore"`
				Learnlink         interface{} `json:"learnlink"`
				Reprice           int         `json:"reprice"`
				Bookprice         string      `json:"bookprice"`
				Difficultyname    string      `json:"difficultyname"`
				Explainvideo      string      `json:"explainvideo"`
				Listenvideo       string      `json:"listenvideo"`
				Explainvideostate int         `json:"explainvideostate"`
				Listenvideostate  int         `json:"listenvideostate"`
				BgPicture         string      `json:"bg_picture"`
				Commoninfo        struct {
					Difficulty int    `json:"difficulty"`
					Name       string `json:"name"`
					Level      int    `json:"level"`
					Color      string `json:"color"`
					Status     int    `json:"status"`
					Count      int    `json:"count"`
					Bgimg      string `json:"bgimg"`
				} `json:"commoninfo"`
			} `json:"items"`
			More   bool `json:"more"`
			Offset int  `json:"offset"`
		} `json:"ent"`
		Ext struct {
			Difficultyinfos []struct {
				Difficulty int    `json:"difficulty"`
				Name       string `json:"name"`
				Level      int    `json:"level"`
				Color      string `json:"color"`
				Status     int    `json:"status"`
				Count      int    `json:"count"`
				Bgimg      string `json:"bgimg"`
			} `json:"difficultyinfos"`
			Guidetext     string `json:"guidetext"`
			Isshowexplain bool   `json:"isshowexplain"`
			Levelinfo     struct {
				Level int    `json:"level"`
				Count int    `json:"count"`
				Name  string `json:"name"`
			} `json:"levelinfo"`
			Pkroute            string      `json:"pkroute"`
			Publishids         interface{} `json:"publishids"`
			Readids            interface{} `json:"readids"`
			Totalbookcn        int         `json:"totalbookcn"`
			Userdifficultyinfo struct {
				UID          int `json:"uid"`
				Difficulty   int `json:"difficulty"`
				Readcount    int `json:"readcount"`
				Publishcount int `json:"publishcount"`
			} `json:"userdifficultyinfo"`
		} `json:"ext"`
	} `json:"data"`
}
