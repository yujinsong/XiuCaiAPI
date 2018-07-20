package model

type (
	CheckCodeModel struct{
		Sessionid string `json:"sessionid"`
		Uuid string `json:"uuid"`
		Devicetype string `json:"devicetype"`
		Softversion string `json:"softversion"`
		Tel string `json:"tel"`
	}
	
	SuperUser struct {
		Id int64 `json:"id"`
		Uuid int64 `json:"uuid"`
		Kind int `json:"kind"`
	}

	User struct {
		Id int64 `json:"id"`
		Flag int `json:"flag"`
		Sex int `json:"sex"`
		Age int `json:"age"`
		Tel string `json:"tel"`
		Regtime string `json:"regtime"`
		Longitude float64 `json:"longitude"`
		Latitude float64 `json:"latitude"`
		Username string `json:"username"`
		Srcimgurl string `json:"srcimgurl"`
		Thumbimgurl string `json:"thumbimgurl"`
		Userdesc string `json:"userdesc"`
		Skills string `json:"skills"`
		Province string `json:"province"`
		City string `json:"city"`
		Backimgurl string `json:"backimgurl"`
		Interest string `json:"interest"`
		Emotion string `json:"emotion"`
		Role string `json:"role"`
		Expvalue int `json:"expvalue"`
		Gradelevel int `json:"gradelevel"`
		Birthday string `json:"birthday"`
		Tags string `json:"tags"`
		Weight int `json:"weight"`
		Score int `json:"score"`
		Fillinfoflag int `json:"fillinfoflag"`
		Extraids string `json:"extraids"`
	}

	SessionModel struct {
		Pushfeedtotopdesc string `json:"pushfeedtotopdesc"`
		Pushfeedtohomedesc string `json:"pushfeedtohomedesc"`
		Feedpraiseno int `json:"feedpraiseno"`
		Feedcommentno int `json:"feedcommentno"`
		Sysserverno int `json:"sysserverno"`
		Orderserviceno int `json:"orderserviceno"`
		Tipserviceno int `json:"tipserviceno"`
		Taskserviceno int `json:"taskserviceno"`
		Atserviceno int `json:"atserviceno"`
		Pushmessageno int `json:"pushmessageno"`
		Invitfriendurl string `json:"invitfriendurl"`
		Displayflag int `json:"displayflag"`
		Xtcostprice int `json:"xtcostprice"`
		Xtminunits int `json:"xtminunits"`
		Ioscurversion string `json:"ioscurversion"`
		Androidcurversion string `json:"androidcurversion"`
		Intervaldays int `json:"intervaldays"`
		Updatemsg string `json:"updatemsg"`
		Costpricemap string `json:"costpricemap"`
		Maxxtprice int `json:"maxxtprice"`
		Maxskillprice int `json:"maxskillprice"`
		Xtshopurl string `json:"xtshopurl"`
		Feedtitle string `json:"feedtitle"`
		Tipsharefee int `json:"tipsharefee"`
		Genpricelist int `json:"genpricelist"`
		Superuserflag int `json:"superuserflag"`
		Thirdpartflag int `json:"thirdpartflag"`
		Rechargeflag int `json:"rechargeflag"`
		Level int `json:"level"`
		Fillimgurlfee float64 `json:"fillimgurlfee"`
		Fillusernamefee float64 `json:"fillusernamefee"`
		Fillsexfee float64 `json:"fillsexfee"`
		Fillbirthdayfee float64 `json:"fillbirthdayfee"`
		Filladdressfee float64 `json:"filladdressfee"`
		Filltagfee float64 `float64:"filltagfee"`
		Fillpersonaltag float64 `json:"fillpersonaltag"`
		Xtautogenprice float64 `json:"xtautogenprice"`
		Xtautogenbackprice float64 `json:"xtautogenbackprice"`
		Autogenxtname string `json:"autogenxtname"`
		Persigninmoney string `json:"persigninmoney"`
		Applink string `json:"apklink"`
		Videomaxsize int `json:"videomaxsize"`
		Videolength int `json:"ideolength"`
		Userbackimglist string `json:"userbackimglist"`
		Circlebackimglist string `json:"circlebackimglist"`
		Cusname string `json:"cusname"`
		Cusimgurl string `json:"cusimgurl"`
	}

	LoginModel struct {
		Sessionid string `json:"sessionid"`
		Uuid int64 `json:"uuid"`
		Tel string `json:"tel"`
		Deviceid string `json:"deviceid"`
		Devicetype string `json:"devicetype"`
		Unittype string `json:"unittype"`
		Sysversion string `json:"sysversion"`
		Softversion string `json:"softversion"`
		Code string `json:"code"`
	}

	LoginResponseModel struct {
		Uuid string `json:"uuid"`
		Imgurl string `json:"imgurl"`
		Imgsize int `json:"imgsize"`
		Thumbimgurl string `json:"thumbimgurl"`
		Username string `json:"username"`
		Sex string `json:"sex"`
		Birthday string `json:"birthday"`
		Needfillinfo string `json:"needfillinfo"`
		Needfillsellinfo string `json:"needfillsellinfo"`
	}

)
