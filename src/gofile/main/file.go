package main

import (
	"os"
	"fmt"
)

/**
目录操作
文件操作的大多数函数都是在os包里面,下面列举了几个目录操作的:
	1. func Mkdir(name string, perm FileMode) error
	创建名称为name的目录,权限设置是perm,例如0777
	2. func MkdirAll(path string, perm FileMode) error
	根据path创建多级子目录,例如astaxie/test1/test2。
	3. func Remove(name string) error
	删除名称为name的目录,当目录下有文件或者其他目录是会出错
	4. func RemoveAll(path string) error
	根据path删除多级子目录,如果path是单个名称,那么该目录不删除。
 */

 // 目录操作
func dirOpt() {
	os.Mkdir("testdir", 0777)
	os.MkdirAll("testdir/subdir1/subdir2", 0777)
	err := os.Remove("testdir")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("testdir")
}

/*
新建文件可以通过如下两个方法:
	1. func Create(name string) (file *File, err Error)
	根据提供的文件名创建新的文件,返回一个文件对象,默认权限是0666的文件,返回的文件对象是可读写
	的。
 	2. func NewFile(fd uintptr, name string) *File
	根据文件描述符创建相应的文件,返回一个文件对象
通过如下两个方法来打开文件:
	1. func Open(name string) (file *File, err Error)
	该方法打开一个名称为name的文件,但是是只读方式,内部实现其实调用了OpenFile。
	2. func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
 	打开名称为name的文件,flag是打开的方式,只读、读写等,perm是权限
写文件函数:
	1. func (file *File) Write(b []byte) (n int, err Error)
	写入byte类型的信息到文件
	2. func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
	在指定位置开始写入byte类型的信息
	3. func (file *File) WriteString(s string) (ret int, err Error)
	写入string信息到文件
读文件函数:
	1. func (file *File) Read(b []byte) (n int, err Error)
	读取数据到b中
	2. func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
	从off开始读取数据到b中
删除文件
	Go语言里面删除文件和删除文件夹是同一个函数:
	1. func Remove(name string) Error
	调用该函数就可以删除文件名为name的文件
*/
// 写文件操作
func wrfileOpt() {
	userfile := "frank.txt"
	fout, err := os.Create(userfile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a []byte\r\n"))
	}
}

// 读文件操作
func rdfileOpt() {
	userfile := "frank.txt"
	fl, err := os.Open(userfile)
	defer fl.Close()
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	buf := make([]byte, 1024) //缓冲区
	for {
		n, _ := fl.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
		//fmt.Println(buf)
	}
}

func main() {
	dirOpt()	//目录操作
 	wrfileOpt()	//写文件操作
 	rdfileOpt()	//读文件操作
}