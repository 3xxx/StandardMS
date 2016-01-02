package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Person_20150330_210449 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Person_20150330_210449{}
	m.Created = "20150330_210449"
	migration.Register("Person_20150330_210449", m)
}

// Run the migrations
func (m *Person_20150330_210449) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE person(`id` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL)")
}

// Reverse the migrations
func (m *Person_20150330_210449) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	m.Sql("DROP TABLE `person`")
}
