package main

import (
    "os"
    "fmt"
)

func main()  {
	//func Mkdir(name string, perm FileMode) error
	//创建名称为name的目录，权限设置是perm，例如0777。
	os.Mkdir("tmp", 0755)
	
	//func MkdirAll(path string, perm FileMode) error
	//根据path创建多级子目录。
	os.MkdirAll("tmp/test/test2", 0755)
	
	//删除名称为name的目录，当目录下有文件或者其他目录时会出错。
    err := os.Remove("tmp")
    if err != nil{
        fmt.Println(err)
	}
	
	//根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除。
    os.RemoveAll("tmp")
}