package main

import (
	"fmt"
	"github.com/beevik/etree"
	"os"
	"os/exec"
	"time"
)

var (
	Xmlfile = "./xml/gameserver.xml"
)

// 参考：https://github.com/beevik/etree

type DBXml struct {
	index  string
	name       string
	host       string
	port       string
	user       string
	pwd       string
	dbname       string
}


func main() {
	region := ""
	timeS := ""
	fmt.Println("命令行参数数量:",len(os.Args))
	for k,v:= range os.Args{
		fmt.Printf("args[%v]=[%v]\n",k,v)
		if k == 1 {
			region = v
		}
		if k == 2 {
			timeS = v
		}
	}
	// 从OSS将配置文件拷下来
	copyXml(region,timeS)
	time.Sleep(time.Duration(2)*time.Second)
	// 用一个结构体数组保存信息
	var dbIntance []DBXml
	// 读取xml配置
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("xml/gameserver/conf/gameserver.xml"); err != nil {
		panic(err)
	}
	// 必须先从根节点找起 不然会出问题 容易
	root := doc.SelectElement("Root")
	fmt.Println("ROOT element:", root.Tag)

	MysqlConf := root.SelectElement("MysqlConf")
	fmt.Println("MysqlConf element:", MysqlConf.Tag)

	for _, db := range MysqlConf.SelectElements("Db") {
		var singleDb DBXml
		for _, attr := range db.Attr {
			// fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
			if attr.Key == "index"{
				singleDb.index = attr.Value
			}
			if attr.Key == "name"{
				singleDb.name = attr.Value
			}
			if attr.Key == "host"{
				singleDb.host = attr.Value
			}
			if attr.Key == "port"{
				singleDb.port = attr.Value
			}
			if attr.Key == "user"{
				singleDb.user = attr.Value
			}
			if attr.Key == "pwd"{
				singleDb.user = attr.Value
			}
			if attr.Key == "dbname"{
				singleDb.user = attr.Value
			}
		}
		// 添加到结构体数组
		dbIntance  = append(dbIntance,singleDb)
	}

	// 遍历结构体数组
	for _, dbid := range dbIntance{
		fmt.Println(dbid.index," ",dbid.user," ",dbid.port," ",dbid.host," ",dbid.name," ",dbid.pwd," ",dbid.dbname)
	}
}

func copyXml(region string,timeS string)  {
	cmd1 := exec.Command("/bin/bash", "-c", `rm -rf xml/*`)
	cmd2 := exec.Command("/bin/bash", "-c", `ossutil cp -f oss://hk4e-config/`+region+`/`+timeS+`/config-`+timeS+`.tgz ./`)
	cmd3 := exec.Command("/bin/bash", "-c", `tar -xvf config-`+timeS+`.tgz -C xml`)
	//执行命令
	if err := cmd1.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	//执行命令
	if err := cmd2.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	//执行命令
	if err := cmd3.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

}
