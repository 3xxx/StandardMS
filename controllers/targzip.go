package controllers

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func TarGz(srcDirPath string, destFilePath string) {
	fw, err := os.Create(destFilePath)
	if err != nil {
		panic(err)
	}
	defer fw.Close()

	// Gzip writer NewWriter创建并返回一个Writer。写入返回值的数据都会在压缩后写入w。调用者有责任在结束写入后调用返回值的Close方法。因为写入的数据可能保存在缓冲中没有刷新入下层。
	// 如要设定Writer.Header字段，调用者必须在第一次调用Write方法或者Close方法之前设置。Header字段的Comment和Name字段是go的utf-8字符串，但下层格式要求为NUL中止的ISO 8859-1 (Latin-1)序列。如果这两个字段的字符串包含NUL或非Latin-1字符，将导致Write方法返回错误。
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// Tar writer NewWriter创建一个写入w的*Writer。
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Check if it's a file or a directory
	f, err := os.Open(srcDirPath)
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	if fi.IsDir() {
		// handle source directory
		fmt.Println("Cerating tar.gz from directory...")
		tarGzDir(srcDirPath, path.Base(srcDirPath), tw)
	} else {
		// handle file directly
		fmt.Println("Cerating tar.gz from " + fi.Name() + "...")
		tarGzFile(srcDirPath, fi.Name(), tw, fi)
	}
	fmt.Println("Well done!")
}

// Deal with directories
// if find files, handle them with tarGzFile
// Every recurrence append the base path to the recPath
// recPath is the path inside of tar.gz
func tarGzDir(srcDirPath string, recPath string, tw *tar.Writer) {
	// Open source diretory
	dir, err := os.Open(srcDirPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// Get file info slice
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	for _, fi := range fis {
		// Append path
		curPath := srcDirPath + "/" + fi.Name()
		// Check it is directory or file
		if fi.IsDir() {
			// Directory
			// (Directory won't add unitl all subfiles are added)
			fmt.Printf("Adding path...%s\n", curPath)
			tarGzDir(curPath, recPath+"/"+fi.Name(), tw)
		} else {
			// File
			fmt.Printf("Adding file...%s\n", curPath)
		}

		tarGzFile(curPath, recPath+"/"+fi.Name(), tw, fi)
	}
}

// Deal with files
func tarGzFile(srcFile string, recPath string, tw *tar.Writer, fi os.FileInfo) {
	if fi.IsDir() {
		// Create tar header
		hdr := new(tar.Header)
		// if last character of header name is '/' it also can be directory
		// but if you don't set Typeflag, error will occur when you untargz
		hdr.Name = recPath + "/"
		hdr.Typeflag = tar.TypeDir
		hdr.Size = 0
		//hdr.Mode = 0755 | c_ISDIR
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// Write hander
		err := tw.WriteHeader(hdr)
		if err != nil {
			panic(err)
		}
	} else {
		// File reader
		fr, err := os.Open(srcFile)
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		// Create tar header
		hdr := new(tar.Header)
		hdr.Name = recPath
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// Write hander
		err = tw.WriteHeader(hdr)
		if err != nil {
			panic(err)
		}

		// Write file data
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
}

// Ungzip and untar from source file to destination directory
// you need check file exist before you call this function
func UnTarGz(srcFilePath string, destDirPath string) {
	fmt.Println("UnTarGzing " + srcFilePath + "...")
	// Create destination directory
	os.Mkdir(destDirPath, os.ModePerm)

	fr, err := os.Open(srcFilePath)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	// Gzip reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	// Tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		//handleError(err)
		fmt.Println("UnTarGzing file..." + hdr.Name)
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			// Write data to file
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("Well done!")
}

//Copyfile
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func TarGzPath(srcDirPath string, destFilePath string) {
	// file write
	fw, err := os.Create(destFilePath) //"tar/lin_golang_src.tar.gz"
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// 打开文件夹
	dir, err := os.Open(srcDirPath) //"file/"
	if err != nil {
		panic(nil)
	}
	defer dir.Close()
	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	// 遍历文件列表
	for _, fi := range fis {
		// 跳过文件夹, 我这里就不递归了
		if fi.IsDir() {
			continue
		}
		// 打印文件名称
		fmt.Println(fi.Name())
		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()
		// 信息头
		h := new(tar.Header)
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()
		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}
		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("tar.gz ok")
}
func zippath(srcDirPath string, destFilePath string) {
	// file write
	file, err := os.Create(destFilePath) //"tar/lin_golang_src.tar.gz"
	if err != nil {
		//		panic(err)
		log.Fatal(err)
	}
	defer file.Close()

	wzip := zip.NewWriter(file)
	defer wzip.Close()
	//	defer func() {
	//		if err := wzip.Close(); err != nil {
	//			log.Fatal(err)
	//		}
	//	}()
	// 打开文件夹
	dir, err := os.Open(srcDirPath) //"file/"
	if err != nil {
		panic(nil)
	}
	defer dir.Close()
	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	// 遍历文件列表
	for _, fi := range fis {
		// 跳过文件夹, 我这里就不递归了
		if fi.IsDir() {
			continue
		}
		// 打印文件名称
		fmt.Println(fi.Name())
		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		header := &zip.FileHeader{
			Name:   fi.Name(),
			Flags:  1 << 11, // 使用utf8编码
			Method: zip.Deflate,
		}
		// 写信息头
		//		err = tw.WriteHeader(h)
		//		if err != nil {
		//			panic(err)
		//		}
		f, err := wzip.CreateHeader(header)
		if err != nil {
			log.Fatal(err)
		}
		len := fi.Size() //[]byte("fr")
		data := make([]byte, len)
		_, err = fr.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("read %d bytes: %q\n", count, data[:count], len)
		_, err = f.Write(data)
		// 写文件
		//		_, err = io.Copy(wzip, fr)

		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println("tar.gz ok")
}
