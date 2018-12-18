package banner

import (
	"fmt"
	"github.com/kataras/iris"
	"community-go/utility"
)

type Banner struct {
	//banner类型id
	Gid int `json:"gid"`
	//banner位置id
	Idx int `json:"idx"`
	//图片地址
	Uri string `json:"uri"`
	//文字内容
	Txt string `json:"txt"`
	//跳转类型
	Scheme string `json:"scheme"`
}

func (banner *Banner) creatBannerByParams(ctx iris.Context) {
	//banner.Gid
}

func GetBanner(ctx iris.Context) (map[string][]Banner, error) {
	rows, err := utility.DB.Query("SELECT * FROM ads")
	defer rows.Close()
	//rows, err := utility.DB.Query("SELECT * FROM ads WHERE gid=?",1)
	if err != nil {
		//fmt.Println(rows)
		fmt.Printf("rows ===== %p", rows)
		fmt.Println("查询错误")
		return map[string][]Banner{}, err
	} else {
		banner := Banner{}
		bannerList := make([]Banner, 0)
		for rows.Next() {
			rows.Scan(&banner.Gid, &banner.Idx, &banner.Uri, &banner.Txt, &banner.Scheme)
			//println(banner)
			err = nil
			bannerList = append(bannerList, banner)
		}

		//fmt.Println(rows)
		fmt.Println("查询正确")

		return map[string][]Banner{"bannerlist": bannerList}, err
	}
}
