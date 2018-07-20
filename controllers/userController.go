package controllers

import (
	"net/http"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"github.com/XiuCai/XiuCaiAPI/common"
	"github.com/XiuCai/XiuCaiAPI/model"
	"fmt"
	"github.com/XiuCai/XiuCaiAPI/services"
	"github.com/XiuCai/XiuCaiAPI/utils"
	"encoding/json"
	"strconv"
)

func CheckCode(rw http.ResponseWriter, req *http.Request)  {
	body, _ := ioutil.ReadAll(req.Body)

	sessionId := gjson.Get(string(body), "sessionid").Str
	uuid := gjson.Get(string(body), "uuid").Int()
//	deviceType := gjson.Get(string(body), "devicetype").Str
//	softVersion := gjson.Get(string(body), "softversion").Str
	tel := gjson.Get(string(body), "tel").Str

	if len(tel) == 0 || !common.CheckPhone(tel) {
		response := common.Response(2, "参数错误", "")
		rw.Write([]byte(response))
		return
	}

	logStr := fmt.Sprintf("---------获取验证码成功，sessionId:%s, uuid:%d, tel:%s", sessionId, uuid, tel)
	fmt.Println(logStr)

	code := common.GenerateCode()

	response := common.Response(0, "获取验证码成功", model.CodeModel{Code: code})
	rw.Write([]byte(response))

}


func CheckSession(rw http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
//	sessionId := gjson.Get(string(body), "sessionid").Str
	uuid := gjson.Get(string(body), "uuid").Int()
	curVersion := gjson.Get(string(body), "curversion").Str
	deviceType := gjson.Get(string(body), "devicetype").Str

	displayFlag := services.GetDisplayFlag(curVersion)
	thirdPartyFlag := 1
	rechargeFlag := 1
	if deviceType == "iOS" && curVersion == "2.4" {
		thirdPartyFlag = common.ThirdPartyFlag
		rechargeFlag = common.RechargeFlag
	}

	userLevel := services.UserLevel(uuid)
	sessionModel := model.SessionModel{
		Pushfeedtotopdesc: common.PushFeedToTopDesc,
		Pushfeedtohomedesc: common.PushFeedToHomeDesc,
		Feedpraiseno: common.FeedPraiseNo,
		Feedcommentno: common.FeedCommentNo,
		Sysserverno: common.SysServiceNo,
		Orderserviceno: common.OrderServiceNo,
		Tipserviceno: common.TipServiceNo,
		Taskserviceno: common.TaskServiceNo,
		Atserviceno: common.AtServiceNo,
		Pushmessageno: common.FansAlertNo,
		Invitfriendurl: common.InviteFriendUrl,
		Displayflag: displayFlag,
		Xtcostprice: common.XTCostPrice,
		Xtminunits: common.XTMinutes,
		Ioscurversion: common.IOSLastVersion,
		Androidcurversion: common.AndroidLastVersion,
		Intervaldays: common.IntervalDays,
		Updatemsg: common.UpdateMessage,
		Maxxtprice: 50,
		Maxskillprice: 100,
		Feedtitle: "",
		Tipsharefee: 1,
		Superuserflag: 1,
		Thirdpartflag: thirdPartyFlag,
		Rechargeflag: rechargeFlag,
		Level: userLevel,
		Fillimgurlfee: 0.1,
		Fillusernamefee: 0.1,
		Fillsexfee: 0.2,
		Fillbirthdayfee: 0.3,
		Filladdressfee: 0.3,
		Filltagfee: 1,
		Fillpersonaltag: 0.2,
		Xtautogenprice: 45,
		Xtautogenbackprice: 1,
		Autogenxtname: "变现",

		Persigninmoney: "0.15",
		Applink: "http://exiucai.com/xiucai.apk",

		Videomaxsize: 30,
		Videolength: 180,
	}

	var cusName string
	var cusImgurl string
	userInfo, err := services.GetUserById(uuid)
	if err != nil {
		cusName = userInfo.Username
		cusImgurl = common.BuildImgUrl(common.ImagePrefix, userInfo.Srcimgurl)
	}

	sessionModel.Cusname = cusName
	sessionModel.Cusimgurl = cusImgurl

	response := common.Response(0, "成功", sessionModel)
	rw.Write([]byte(response))

}

func Login(rw http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)

	var loginModel model.LoginModel
	err := json.Unmarshal(body, &loginModel)
	if err != nil || !common.CheckPhone(loginModel.Tel) {
		response := common.Response(1, "登录失败", "")
		rw.Write([]byte(response))
		return
	}

	userInfo, err := services.GetUserByTel(loginModel.Tel)
	if err != nil {
		utils.Logger.Error("---- login error, tel:", loginModel.Tel, " error:", err)
		response := common.Response(1, "登录失败", "")
		rw.Write([]byte(response))
		return
	}

	if userInfo.Id == 0 {
		// 注册
		uuid := services.Register(loginModel)
		if uuid < 0 {
			response := common.Response(1, "登录失败", "")
			rw.Write([]byte(response))
			return
		}

		loginResponse := model.LoginResponseModel{
			Uuid: strconv.Itoa(int(uuid)),
			Imgurl: "",
			Imgsize: common.ImageSize,
			Thumbimgurl: "",
			Username: "",
			Sex: "0",
			Birthday: "",
			Needfillinfo: "1",
			Needfillsellinfo: "1",
		}
		response := common.Response(0, "登录成功", loginResponse)
		rw.Write([]byte(response))
		return
	} else {
		loginResponse := model.LoginResponseModel{
			Uuid: strconv.Itoa(int(userInfo.Id)),
			Imgurl: common.BuildImgUrl(common.ImagePrefix, userInfo.Srcimgurl),
			Imgsize: common.ImageSize,
			Thumbimgurl: common.BuildImgUrl(common.ImagePrefix, userInfo.Thumbimgurl),
			Username: userInfo.Username,
			Sex: strconv.Itoa(userInfo.Sex),
			Birthday: userInfo.Birthday,
			Needfillinfo: "0",
			Needfillsellinfo: "0",
		}

		response := common.Response(0, "登录成功", loginResponse)
		rw.Write([]byte(response))
		return
	}
}