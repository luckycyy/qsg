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
	ns := beego.NewNamespace("/v1",

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
	beego.Router("/static/saveArrive", &controllers.SaveArriveController{})
	//todo 上面修改
}
