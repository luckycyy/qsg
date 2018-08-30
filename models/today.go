package models

import "time"

type Today struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(255);null"`
	Namesx     string    `orm:"column(namesx);size(255);null"`
	Grade      string    `orm:"column(grade);size(255);null"`
	Class      string    `orm:"column(class);size(255);null"`
	Tel        string    `orm:"column(tel);size(255);null"`
	Arrive     int       `orm:"column(arrive);null"`
	ArriveTime time.Time `orm:"column(arrive_time);type(datetime);null"`
	Date       time.Time `orm:"column(date);type(date);null"`
}
