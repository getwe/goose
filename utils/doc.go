//
// 
// 

// Package utils 是goose中的一个公共包,存放的是公共类型的定义和一些基础工具.
//
// BigFile是逻辑大文件操作类,支持追加写和随机读,利用多个小文件组成的逻辑大文件提供服务
// 其最大的作用是避免磁盘文件过大超出系统限制
//
// jsonfile提供结构体和磁盘文件的序列化/反序列化函数
//
// type.go定义了goose使用到的类型
package utils
