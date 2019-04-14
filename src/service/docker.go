package service

import (
	"os"
	"os/exec"
)

func Login(name string, passwd string, path string) (err error) {
	cmd := exec.Command("docker", "login", "-u", name, "-p", passwd, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 运行命令
	err = cmd.Run()
	return
}

func Pull(path string) (err error) {
	// docker pull path
	cmd := exec.Command("docker", "pull",
		path,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 运行命令
	err = cmd.Run()
	return
}

func Tag(name string, tagname string) (err error) {
	cmd := exec.Command("docker", "tag", name,
		tagname,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 运行命令
	err = cmd.Run()
	return
}

func Push() {

}

func Save(path string, dir string, name string) (err error) {
	cmd := exec.Command("docker", "save",
		"-o", dir+"/"+name, path,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 运行命令
	err = cmd.Run()
	return
}
