package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func RemoveRepeatedElement(arr []string) []string {

	result := []string{}

	for i := 0; i < len(arr); i++ {

		exist := false

		for j := 0; j < len(result); j++ {

			if result[j] == arr[i] {

				exist = true
				break
			}
		}

		if !exist {

			result = append(result, arr[i])
		}
	}
	return result
}


func ReadFile(Path string) (result string, err error) {

	f, err1 := os.Open(Path)
	if err1 != nil {
		fmt.Println("os.Open err", err1)
		err = err1
		return
	}
	defer f.Close()


	var content []byte
	buf := make([]byte, 4096)
	for {
		n, err2 := f.Read(buf)
		if err2 != nil && err2 != io.EOF {
			fmt.Println("f.Read err:", err2)
			err = err2
			return
		}
		if err2 == io.EOF {
			break
		}

		content = append(content, buf[:n]...)
	}

	result = string(content)
	return
}


func CreateFile(path, data string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()


	_, err = f.WriteString(data)
	if err != nil {
		fmt.Println("f.WriteString err:", err)
		return
	}
	return
}


func BackupFile(dst, src string) (err error) {

	fSrc, err := os.Open(src)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer fSrc.Close()


	fDst, err := os.Create(dst)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer fDst.Close()

	buf := make([]byte, 4096)
	for {
		n, err := fSrc.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("f.Read err", err)
			return  err
		}
		if err == io.EOF {
			break
		}


		if _, err = fDst.Write(buf[:n]); err != nil {
			return err
		}
	}
	return
}

func main() {

	path := `list.txt`


	Backup := `list_backup.txt`


	if err := BackupFile(Backup, path); err != nil {
		fmt.Println("BackupFile err", err)
		return
	}

	str, err := ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}


	str = strings.Replace(str, "\r\n", "\n", -1)


	strSlice := strings.Split(str, "\n")


	strSlice = RemoveRepeatedElement(strSlice)


	sort.Strings(strSlice)


	str = strings.Join(strSlice, "\n")

	err = CreateFile(path, str)
	if err != nil {
		fmt.Println("CreateFile err:", err)
		return
	}

	fmt.Println("Eroare creare fisier temporar！")
}
