package missionqueue

import "testing"

func TestNewMission(t *testing.T) {
	cookie := "_T_WM=cbd8e7f48d959f407df78c57138c72c7; ALF=1523513664; SCF=AoCfn5Asc3ogMMm4c8GQBsx_2V560kzRWycdOOW72WM5PXa-pfzI7xsLMgOYcj98eQs93Jg-NxtIoy6z-eFWOug.; SUB=_2A253oxwRDeRhGeVO4lcV9SrFyTuIHXVVb6RZrDV6PUNbktAKLWnskW1NTXO5Qk3p4i2P9Vv7eNAIRlPaNoOK325M; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WFyP-Tpli0Z3Xai11a52FpJ5JpX5KMhUgL.Foe71K-XSKB4eoM2dJLoI0qLxK-L12qLBKBLxK-LBKBLBK.LxKBLBo.L1K5LxKML1-2L1hBLxK-LBo5L12qLxK-L12qLBoMt; SUHB=0D-Fw9eB2CaCun; SSOLoginState=1520921665"

	url := "https://weibo.cn/3095454927/profile?page=1"

	mysql := "root:123456@tcp(localhost:3306)/monika?charset=utf8"

	mission, _ := NewMission(cookie, mysql)

	mission.setMission(url)

	//url = "https://weibo.cn/3095454927/profile?page=100"
	//mission.setMission(url)
	mission.DoMission()
}
