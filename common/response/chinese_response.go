package response

import "github.com/jiangrx816/wechat/model"

type ChineseBookNavNameResponse struct {
	List []model.SBookName `json:"list"`
}

type ChineseBookResponse struct {
	Page      int                   `json:"page"`
	Total     int64                 `json:"total"`
	TotalPage int                   `json:"total_page"`
	List      []ResponseChineseBook `json:"list"`
}
type ResponseBookInfoCount struct {
	BookId    string `json:"book_id"`
	BookCount string `json:"book_count"`
}

type ResponseChineseBook struct {
	Id        int    `json:"-"`
	BookId    string `json:"book_id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Level     int8   `json:"-"`
	Position  int    `json:"-"`
	BookCount string `json:"book_count"`
}

type ChineseBookInfoResponse struct {
	Info []ResponseChineseBookInfo `json:"info"`
}
type ResponseChineseBookInfo struct {
	Id       int    `json:"id"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
}

type EnglishBookInfoResponse struct {
	Data ResponseEnglishBookInfoData `json:"data"`
}
type ResponseEnglishBookInfo struct {
	Ret  int                         `json:"ret"`
	Data ResponseEnglishBookInfoData `json:"data"`
}
type ResponseEnglishBookInfoData struct {
	Ent struct {
		Items []struct {
			Bookid  int64 `json:"bookid"`
			Pageid  int64 `json:"pageid"`
			Index   int   `json:"index"`
			Picture struct {
				Tiny   string `json:"tiny"`
				Origin string `json:"origin"`
				W      int    `json:"w"`
				H      int    `json:"h"`
			} `json:"picture"`
			Ct               int         `json:"ct"`
			Ut               int         `json:"ut"`
			State            int         `json:"state"`
			Text             string      `json:"text"`
			Picturev2        interface{} `json:"picturev2"`
			Flag             int         `json:"flag"`
			Topicstate       int         `json:"topicstate"`
			Ext              string      `json:"ext"`
			Recordtext       string      `json:"recordtext"`
			Translation      string      `json:"translation"`
			Hasteach         bool        `json:"hasteach"`
			Hasrecord        bool        `json:"hasrecord"`
			TextPinyin       string      `json:"textPinyin"`
			RecordtextPinyin string      `json:"recordtextPinyin"`
			IsOpenLp         bool        `json:"is_open_lp"`
			IsOpenRp         bool        `json:"is_open_rp"`
			BgPicture        string      `json:"bg_picture"`
			Scoretext        string      `json:"scoretext"`
		} `json:"items"`
		Offset         int  `json:"offset"`
		More           bool `json:"more"`
		Readpagecn     int  `json:"readpagecn"`
		Isshowvipguide bool `json:"isshowvipguide"`
	} `json:"ent"`
	Ext struct {
		Adpic    string `json:"adpic"`
		Adrouter string `json:"adrouter"`
		Bookinfo struct {
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
			Topicid           int64       `json:"topicid"`
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
			Learnlink         []int       `json:"learnlink"`
			Reprice           int         `json:"reprice"`
			Bookprice         string      `json:"bookprice"`
			Difficultyname    string      `json:"difficultyname"`
			Explainvideo      string      `json:"explainvideo"`
			Listenvideo       string      `json:"listenvideo"`
			Explainvideostate int         `json:"explainvideostate"`
			Listenvideostate  int         `json:"listenvideostate"`
			BgPicture         string      `json:"bg_picture"`
		} `json:"bookinfo"`
		Cwrouter    string `json:"cwrouter"`
		Endtext     string `json:"endtext"`
		Evaluations []struct {
			Recordid   int64 `json:"recordid"`
			Totalscore struct {
				Content string `json:"content"`
				Score   int    `json:"score"`
				Rank    int    `json:"rank"`
				Index   int    `json:"index"`
				Startts int    `json:"startts"`
				Endts   int    `json:"endts"`
			} `json:"totalscore"`
			Wordscore []struct {
				Content string `json:"content"`
				Score   int    `json:"score"`
				Rank    int    `json:"rank"`
				Index   int    `json:"index"`
				Startts int    `json:"startts"`
				Endts   int    `json:"endts"`
			} `json:"wordscore"`
			Productid int64 `json:"productid"`
			Ut        int   `json:"ut"`
		} `json:"evaluations"`
		Isconfirm      bool `json:"isconfirm"`
		Isshowvipguide bool `json:"isshowvipguide"`
		Productinfo    struct {
			Productid   int64 `json:"productid"`
			Bookid      int64 `json:"bookid"`
			UID         int   `json:"uid"`
			State       int   `json:"state"`
			Ct          int   `json:"ct"`
			Ut          int   `json:"ut"`
			Playcount   int   `json:"playcount"`
			Likecount   int   `json:"likecount"`
			Islike      bool  `json:"islike"`
			Publishtime int   `json:"publishtime"`
			Producttype int   `json:"producttype"`
			Score       int   `json:"score"`
			Rank        int   `json:"rank"`
			Iscollect   bool  `json:"iscollect"`
			Dt          int   `json:"dt"`
			Scene       int   `json:"scene"`
		} `json:"productinfo"`
		Readpagecn int `json:"readpagecn"`
		Records    []struct {
			Recordid  int64  `json:"recordid"`
			Productid int64  `json:"productid"`
			URL       string `json:"url"`
			Ct        int    `json:"ct"`
			Ut        int    `json:"ut"`
			Pageid    int64  `json:"pageid"`
			Duration  int    `json:"duration"`
			Bookid    int64  `json:"bookid"`
			Rawurl    string `json:"rawurl"`
		} `json:"records"`
		Route        string `json:"route"`
		Sharecontent string `json:"sharecontent"`
		Users        []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			Gender      int    `json:"gender"`
			Sign        string `json:"sign"`
			Regtype     int    `json:"regtype"`
			Cate        int    `json:"cate"`
			Birthday    int    `json:"birthday"`
			Origavatar  string `json:"origavatar"`
			Audiobrief  string `json:"audiobrief"`
			Audiolength int    `json:"audiolength"`
			Ct          int    `json:"ct"`
			Rt          int    `json:"rt"`
			Country     string `json:"country"`
			Govold      int    `json:"govold"`
			Rmk         string `json:"rmk"`
			Gov         int    `json:"gov"`
			Title       string `json:"title"`
			Source      int    `json:"source"`
			Juniortitle string `json:"juniortitle"`
			Enname      string `json:"enname"`
			Iseligible  bool   `json:"iseligible"`
			State       int    `json:"state"`
			Fullname    struct {
				Firstname  string `json:"firstname"`
				Familyname string `json:"familyname"`
				Middlename string `json:"middlename"`
				Status     int    `json:"status"`
			} `json:"fullname"`
			Puid     string `json:"puid"`
			Agelevel string `json:"agelevel"`
		} `json:"users"`
	} `json:"ext"`
}

type EnglishBookResponse struct {
	Data ResponseEnglishBookData `json:"data"`
}

type ResponseEnglishBook struct {
	Ret  int                     `json:"ret"`
	Data ResponseEnglishBookData `json:"data"`
}

type ResponseEnglishBookData struct {
	Ent struct {
		Items  []interface{} `json:"items"`
		More   bool          `json:"more"`
		Offset int           `json:"offset"`
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
}

type PoetryBookResponse struct {
	Data []ResponsePoetryBook `json:"data"`
}
type ResponsePoetryBook struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TypeID     string `json:"type_id"`
	Img        string `json:"img"`
	Difficulty string `json:"difficulty"`
	Author     string `json:"author"`
	FreeRank   string `json:"free_rank"`
	IsRead     int    `json:"is_read"`
	IsFree     int    `json:"is_free"`
	IsVip      int    `json:"is_vip"`
}
type ResponsePoetryBookJson struct {
	Code     int                  `json:"code"`
	Msg      string               `json:"msg"`
	Data     []ResponsePoetryBook `json:"data"`
	TypeList struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Img  string `json:"img"`
	} `json:"type_list"`
}

type PoetryBookInfoResponse struct {
	Data map[string]PoetryBookInfo `json:"data"`
}

type ResponsePoetryBookInfo struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg"`
	Data map[string]PoetryBookInfo `json:"data"`
}

type PoetryBookInfo struct {
	ID           int         `json:"id"`
	BookID       int         `json:"book_id"`
	CnContent    string      `json:"cn_content"`
	EnContent    string      `json:"en_content"`
	SpellContent string      `json:"spell_content"`
	CnArr        string      `json:"cn_arr"`
	EnArr        string      `json:"en_arr"`
	PlayTime     int         `json:"play_time"`
	EnPlayTime   int         `json:"en_play_time"`
	Img          string      `json:"img"`
	Meaning      string      `json:"meaning"`
	MeaningCnArr interface{} `json:"meaning_cn_arr"`
	MeaningEn    string      `json:"meaning_en"`
	MeaningEnArr interface{} `json:"meaning_en_arr"`
	Rank         int         `json:"rank"`
	IsShow       int         `json:"is_show"`
}

type AlbumPictureInfoResponse struct {
	Info ResponseAlbumBookInfo `json:"info"`
}

type AlbumBookIndexResponse struct {
	Page      int                      `json:"page"`
	Total     int64                    `json:"total"`
	TotalPage int                      `json:"total_page"`
	List      []ResponseAlbumBookIndex `json:"list"`
}

type ResponseAlbumBookIndex struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Position uint8  `json:"position"`
}

type AlbumBookListResponse struct {
	Page      int                     `json:"page"`
	Total     int64                   `json:"total"`
	TotalPage int                     `json:"total_page"`
	List      []ResponseAlbumBookInfo `json:"list"`
}

type ResponseAlbumBookInfo struct {
	Id       int    `json:"id"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Duration string `json:"duration"`
}
