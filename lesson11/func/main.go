package main

func main () {
	// 目录
	// home 目录
	// dir, err := os.UserHomeDir()
	// fmt.Println(dir, err)

	// dir, err := os.UserCacheDir()
	// fmt.Println(dir, err)

	// dir, err := os.UserConfigDir()
	// fmt.Println(dir, err)

	// fmt.Println(os.TempDir())

	// Mkdir 创建目录
	// err := os.Mkdir("dir1", os.ModePerm)
	// fmt.Println(err)

	// Rename 重命名或移动目录
	// err := os.Rename("dir1", "dir2")
	// fmt.Println(err)

	// MkdirAll 创建目录
	// err := os.MkdirAll("dir2/subdir", os.ModePerm)
	// fmt.Println(err)

	// Remove 删除文件或目录
	// err := os.Remove("dir1")
	// fmt.Println(err)

	// RemoveAll 删除文件或目录
	// err := os.RemoveAll("dir2")
	// fmt.Println(err)

	// Getwd 获取当前工作目录
	// dir, err := os.Getwd()
	// fmt.Println(dir, err)

	// Chdir 修改当前工作目录
	// err := os.Chdir("/Users/weirubo/Desktop/go-package/lesson11/file")
	// fmt.Println(err)
	// dir, err := os.Getwd()
	// fmt.Println(dir, err)

	// ======================

	// 文件
	// Create 创建文件，如果文件已存在，将被截断
	// f1, err := os.Create("file1.txt")
	// fmt.Println(f1, err)

	// Open
	// f1, err := os.Open("file1.txt")
	// fmt.Println(f1, err)

	// OpenFile
	// f1, err := os.OpenFile("file1.txt", os.O_RDWR, os.ModePerm)
	// fmt.Println(f1, err)


	// =======================
	// FileInfo 文件信息
	// Lstat
	// fi, err := os.Lstat("file1.txt")
	// fmt.Println(fi, err)

	// Stat
	// fi, err := os.Stat("file1.txt")
	// fmt.Println(fi, err)


	// ========================
	// FileMode 文件权限
	// fi, err := os.Stat("file1.txt")
	// fmt.Println(fi, err)
	// fm := fi.Mode()
	// fmt.Println(fm)

}
