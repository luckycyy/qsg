package controllers

import (
	"encoding/json"
	"errors"
	"qsg/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// StuController operations for Stu
type StuController struct {
	beego.Controller
}
type UpAllController struct {
	beego.Controller
}
type DownAllController struct {
	beego.Controller
}
type DeleteStusController struct {
	beego.Controller
}
// URLMapping ...
func (c *StuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Stu
// @Param	body		body 	models.Stu	true		"body for Stu content"
// @Success 201 {int} models.Stu
// @Failure 403 body is empty
// @router / [post]
func (c *StuController) Post() {
	var v models.Stu
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddStu(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Stu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Stu
// @Failure 403 :id is empty
// @router /:id [get]
func (c *StuController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetStuById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Stu
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Stu
// @Failure 403
// @router / [get]
func (c *StuController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllStu(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Stu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Stu	true		"body for Stu content"
// @Success 200 {object} models.Stu
// @Failure 403 :id is not int
// @router /:id [put]
func (c *StuController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Stu{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateStuById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Stu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *StuController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteStu(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UpAllController) Get() {
	fmt.Println("receive up all grade req")
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE stu SET grade=grade+1").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("UpAll row affected nums: ", num)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

func (c *DownAllController) Get() {
	fmt.Println("receive down all grade req")
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE stu SET grade=grade-1").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("UpAll row affected nums: ", num)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

func (c *DeleteStusController) Get() {
	stusId:=c.GetString("stus")
	fmt.Println("stusIdStr is:"+stusId)
	if stusId==""{
		fmt.Println("get data err")
		return
	}
	stusIdArr := strings.Split(stusId,"-")
	fmt.Println(len(stusIdArr))

	c.Data["json"] = "OK"
	for _, stuIdStr := range stusIdArr{
		fmt.Println(stuIdStr)
		stuId, err := strconv.Atoi(stuIdStr)
		if err != nil {
			fmt.Println("stuId can not convert to int: ", stuId)
		}
		if err := models.DeleteStu(stuId); err != nil {
			fmt.Println("delete stu failed,stuid:"+strconv.Itoa(stuId))
			c.Data["json"] = err.Error()
			break
		}
		fmt.Println("delete stu success,stuid:"+strconv.Itoa(stuId))
	}

	c.ServeJSON()
}
