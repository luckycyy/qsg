// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"qsg/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/static/v1",

		beego.NSNamespace("/callname",
			beego.NSInclude(
				&controllers.CallnameController{},
			),
		),

		beego.NSNamespace("/stu",
			beego.NSInclude(
				&controllers.StuController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/static/todayList", &controllers.TodayListController{})
	beego.Router("/static/saveArrivezw", &controllers.SaveArriveZwController{})
	beego.Router("/static/saveArrivews", &controllers.SaveArriveWsController{})
	beego.Router("/static/saveArriveNFC", &controllers.SaveArriveNfcController{})
	beego.Router("/static/upAll", &controllers.UpAllController{})
	beego.Router("/static/downAll", &controllers.DownAllController{})
	beego.Router("/static/deleteStus", &controllers.DeleteStusController{})

}
