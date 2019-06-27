package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Courses_20190625_130942 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Courses_20190625_130942{}
	m.Created = "20190625_130942"

	migration.Register("Courses_20190625_130942", m)
}

// Run the migrations
func (m *Courses_20190625_130942) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Courses(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`requisites` varchar(128) NOT NULL,`workload` int(11) DEFAULT NULL,`price` float NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Courses_20190625_130942) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Courses`")
}
