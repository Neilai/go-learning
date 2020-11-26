## go module 使用

- golang 升级到 1.11后可以使用 新的模块管理方式，之前是使用类似Vendor的机制，所有的包往Gopath里面找当modules 功能启用时，依赖包的存放位置变更为$GOPATH/pkg。
- goland若是识别不到go modules的包，需要在设置里面开启go module integration。
- go.mod 类似 package.json  go.sum 类似package-lock.json 
- 常用命令 go mod init(类似npm init) ,go mod tidy(类似npm install)
- 使用go mod时 , 导入本地包不要使用相对路径，要根据顶层Mod的package name往下写（和目录名无关）

## init函数及其调用顺序

init是导入的时候执行的，先调用依赖的依赖init，同层依赖init按字典序调用。

