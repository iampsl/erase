package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var char byte

func init() {
	rand.Seed(time.Now().Unix())
	char = byte(rand.Intn(256))
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage:%s file...\n", os.Args[0])
		return
	}
	for i := 1; i < len(os.Args); i++ {
		ok, err := isDir(os.Args[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if ok {
			fmt.Printf("%s is dir\n", os.Args[i])
			return
		}
	}
	for i := 1; i < len(os.Args); i++ {
		if err := erase(os.Args[i]); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func isDir(filePath string) (bool, error) {
	pfile, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return false, err
	}
	defer pfile.Close()
	pfileInfo, err := pfile.Stat()
	if err != nil {
		return false, err
	}
	return pfileInfo.IsDir(), nil
}

func erase(filePath string) error {
	pfile, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	pfileInfo, err := pfile.Stat()
	if err != nil {
		pfile.Close()
		return err
	}
	writer := bufio.NewWriter(pfile)
	for i := int64(0); i < pfileInfo.Size(); i++ {
		if err := writer.WriteByte(char); err != nil {
			writer.Flush()
			pfile.Close()
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		pfile.Close()
		return err
	}
	return pfile.Close()
}
