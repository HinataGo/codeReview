package io

// 源码  os/types.go
// // A FileInfo describes a file and is returned by Stat and Lstat.
// type FileInfo interface {
//    Name() string       // base name of the file 文件名.扩展名 aa.txt
//    Size() int64        // 文件大小，字节数 12540
//    Mode() FileMode     // 文件权限 -rw-rw-rw-
//    ModTime() time.Time // 修改时间 2018-04-13 16:30:53 +0800 CST
//    IsDir() bool        // 是否文件夹
//    Sys() interface{}   // 基础数据源接口(can return nil)
// }
