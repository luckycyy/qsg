package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"qsg/models"
	"encoding/json"
	"time"
	"strconv"
)

type TodayListController struct {
	beego.Controller
}
type SaveArriveZwController struct {
	beego.Controller
}
type SaveArriveWsController struct {
	beego.Controller
}
type SaveArriveNfcController struct {
	beego.Controller
}

func (c *SaveArriveNfcController) Get() {
	stuId := c.GetString("stuId")
	fmt.Println( "stuId is:"+stuId)
	if stuId == "" {
		fmt.Println("get data err")
		return
	}
	//查询数据库获取姓名，组装数据入点名表。不判断学生是几顿，只要打卡了就代表吃了以便记录实际次数
	iStuId, _ := strconv.Atoi(stuId)
	v, err := models.GetStuById(iStuId)
	if err != nil {
		c.Data["json"] = "no_stu_id"
		fmt.Println("stu id is not in database.stuId is:" + stuId)
	} else {
		var arriveStu models.Callname
		o := orm.NewOrm()
		arriveStu.StuId = iStuId
		arriveStu.Name = v.Name
		//1 到达 2未到达
		arriveStu.Arrive = 1
		arriveStu.ArriveTime = time.Now()
		arriveStu.Date = time.Now()

		banci := "zw"
		if time.Now().Hour() > 14 {
			banci = "ws"
		}
		fmt.Printf("now hour: %d,banci:%s",time.Now().Hour(),banci)
		arriveStu.Banci = banci
		num, err := o.QueryTable("callname").Filter("date", time.Now().Format("2006-01-02")).Filter("stu_id", arriveStu.StuId).Filter("banci", arriveStu.Banci).Update(orm.Params{
			"arrive":      arriveStu.Arrive,
			"arrive_time": arriveStu.ArriveTime,
		})
		fmt.Printf("Returned Rows Num: %s, %s", num, err)
		if num == 0 {
			id, err := o.Insert(&arriveStu)
			if err == nil {
				fmt.Println(id)
			}
		}
		c.Data["json"] = arriveStu.Name
	}
	//app界面从未接到名单移除这个人，
	c.ServeJSON()
}

func (c *TodayListController) Get() {

	date := time.Now().Format("2006-01-02")
	banci := c.GetString("banci")
	fmt.Println("date is :"+date, "banci is:"+banci)
	if date == "" || banci == "" {
		fmt.Println("get data err")
		return
	}

	o := orm.NewOrm()
	var todayNames []models.Today
	var condition string
	if "zw" == banci {
		condition = "count='1' or count='2'"
	} else if "ws" == banci {
		condition = "count='2' or count='3'"
	}
	num, err := o.Raw("select id,name,namesx,grade,class,tel from stu where " + condition).QueryRows(&todayNames)
	if err == nil {
		fmt.Println("today nums: ", num)
	} else {
		fmt.Print(err.Error())
	}
	var callnames []models.Callname
	num, err = o.Raw("select stu_id,arrive,arrive_time from callname where banci=? and date=?", banci, time.Now().Format("2006-01-02")).QueryRows(&callnames)
	if err == nil {
		fmt.Println("callnames nums: ", num)
	} else {
		fmt.Print(err.Error())
	}

	for i := 0; i < len(todayNames); i++ {
		todayNames[i].Arrive = 2 //先将名单中的学生设置为没到达，匹配中了设置为到达.1是到达 2是没到达
		todayNames[i].ArriveTime = time.Now()
		for j := 0; j < len(callnames); j++ {
			if callnames[j].StuId == todayNames[i].Id && callnames[j].Arrive == 1 {
				todayNames[i].Arrive = 1
				todayNames[i].ArriveTime = callnames[j].ArriveTime
				break
			}
		}
	}

	c.Data["json"] = todayNames
	c.ServeJSON()

}
func (c *SaveArriveZwController) Post() {

	var arriveStu models.Today
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &arriveStu); err == nil {
		fmt.Println(&arriveStu)
		o := orm.NewOrm()
		arriveStu.ArriveTime = time.Now()
		num, err := o.QueryTable("callname").Filter("date", time.Now().Format("2006-01-02")).Filter("stu_id", arriveStu.Id).Filter("banci", "zw").Update(orm.Params{
			"arrive":      arriveStu.Arrive,
			"arrive_time": arriveStu.ArriveTime,
		})
		fmt.Printf("Returned Rows Num: %s, %s", num, err)
		if num == 0 {
			var stu models.Callname
			stu.StuId = arriveStu.Id
			stu.Name = arriveStu.Name
			stu.Date = time.Now()
			stu.ArriveTime = arriveStu.ArriveTime
			stu.Arrive = arriveStu.Arrive
			stu.Banci = "zw"
			id, err := o.Insert(&stu)
			if err == nil {
				fmt.Println(id)
			}
		}
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}
func (c *SaveArriveWsController) Post() {

	var arriveStu models.Today
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &arriveStu); err == nil {
		fmt.Println(&arriveStu)
		o := orm.NewOrm()
		arriveStu.ArriveTime = time.Now()
		num, err := o.QueryTable("callname").Filter("date", time.Now().Format("2006-01-02")).Filter("stu_id", arriveStu.Id).Filter("banci", "ws").Update(orm.Params{
			"arrive":      arriveStu.Arrive,
			"arrive_time": arriveStu.ArriveTime,
		})
		fmt.Printf("Returned Rows Num: %s, %s", num, err)
		if num == 0 {
			var stu models.Callname
			stu.StuId = arriveStu.Id
			stu.Name = arriveStu.Name
			stu.Date = time.Now()
			stu.ArriveTime = arriveStu.ArriveTime
			stu.Arrive = arriveStu.Arrive
			stu.Banci = "ws"
			id, err := o.Insert(&stu)
			if err == nil {
				fmt.Println(id)
			}
		}
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}
