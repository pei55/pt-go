# go的模板库

一个go的模板库，用来使用gogen生成新的项目


### 代码获取

git clone https://github.com/pei55/pt-go.git

### 编译
该编译脚本只支持编译成linux下运行程序，由于未对windows系统进行兼容设计，不保证
windows可用性
``` shell
cd bdmdbaike
sh ./build.sh
```
### 部署
    1.把编译好的程序以及配置目录下的配置文件拷贝到目标机器
    2.修改配置文件对应配置即可
    3.拷贝bdbaike.service文件到/etc/systemd/system/bdbaike.service
    4.确保替换ExecStart中的/path/to/myapp为你的应用程序的实际路径。

### 管理及启动
```text
# 重新加载systemd，以便它知道新的服务文件
sudo systemctl daemon-reload

# 启用服务，使其在启动时自动启动
sudo systemctl enable bdbaike.service

# 立即启动服务
sudo systemctl start bdbaike.service
你可以使用以下命令来检查服务的状态：

sudo systemctl status bdbaike.service
如果你需要停止或禁用服务，可以使用以下命令：

# 停止服务
sudo systemctl stop bdbaike.service

# 禁用服务，使其不在启动时自动启动
sudo systemctl disable bdbaike.service
```


### 作者
ag

### 许可证

### 致谢
感谢任何对项目有所贡献的人
灵感来源
等等

