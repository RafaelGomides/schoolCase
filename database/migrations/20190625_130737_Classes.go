package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Classes_20190625_130737 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Classes_20190625_130737{}
	m.Created = "20190625_130737"

	migration.Register("Classes_20190625_130737", m)
}

// Run the migrations
func (m *Classes_20190625_130737) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Classes(`id` int(11) NOT NULL AUTO_INCREMENT,`start_date` varchar(128) NOT NULL,`finish_date` varchar(128) NOT NULL,`workload` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Classes_20190625_130737) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Classes`")
}
