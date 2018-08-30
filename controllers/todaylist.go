package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"qsg/models"
	"encoding/json"
)

type TodayListController struct {
	beego.Controller
}
type SaveArriveController struct {
	beego.Controller
}
func (c *TodayListController) Get() {

	date:=c.GetString("date")
	fmt.Println("date is :"+date)
	if date == ""{
		fmt.Println("get date err")
		return
	}

	//var todayNames []*models.Callname
	o := orm.NewOrm()
	//
	var todayNames []models.Today
	num, err := o.Raw("select stu.id,stu.name,namesx,grade,class,tel,arrive,arrive_time,date from stu,callname where date=? and stu.id=callname.id", date).QueryRows(&todayNames)
	if err == nil {
		fmt.Println("today nums: ", num)
	}else{
		fmt.Print(err.Error())
	}

	//num, err := o.QueryTable("callname").Filter("date", date).All(&todayNames)
	//fmt.Printf("Returned Rows Num: %d, %s", num, err)

	c.Data["json"] = todayNames
	c.ServeJSON()


}
func (c *SaveArriveController) Post() {

	var arriveArr[] models.Today
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &arriveArr); err == nil {
		//if _, err := models.AddCallname(&v); err == nil {
		//	c.Ctx.Output.SetStatus(201)
		//	c.Data["json"] = v
		//} else {
		//	c.Data["json"] = err.Error()
		//}
		fmt.Println(&arriveArr)
		fmt.Println(arriveArr[0].Name,arriveArr[1].Name)
		for i:=0;i<len(arriveArr);i++ {
			arriveArr[i].Arrive=1
		}
		var callnameArr[] models.Callname
		//todo 把arriveArr转换成callname才能写入数据库
		o := orm.NewOrm()
		num, err := o.QueryTable("callname").Filter("date", "2018-08-28").All(&arriveArr)
		fmt.Printf("Returned Rows Num: %s, %s", num, err)
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}
