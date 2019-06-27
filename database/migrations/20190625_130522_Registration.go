package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Registration_20190625_130522 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Registration_20190625_130522{}
	m.Created = "20190625_130522"

	migration.Register("Registration_20190625_130522", m)
}

// Run the migrations
func (m *Registration_20190625_130522) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Registration(`id` int(11) NOT NULL AUTO_INCREMENT,`register_date` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Registration_20190625_130522) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Registration`")
}
