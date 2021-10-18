### 镜像构建和nsenter操作
构建本地镜像,第一次写的时候忘记不同平台的区别,直接将mac编译的二进制文件放到dockerfile中,不能正常运行

`docker build -t gohttp:v1.1 .`

将镜像推送到Docker官方镜像仓库

`docker tag gohttp:v1.1 kevin1111/gohttp:v1.1`

`docker push kevin1111/gohttp:v1.1`

通过Docker 命令本地启动httpserver

`docker run -d -it -p 8000:8080 kevin1111/gohttp:v1.1`

通过nsenter 进入容器查看IP配置
`docker ps | grep gohttp`

`docker inspect ae98047d9fe4 | grep -i pid`

`nsenter -t 99266 -n ip a`


### Dockerfile的最佳实践

构建Dockerfile实践的目标: 镜像需要**易于管理、少漏洞、镜像小、层级少、利用缓存**

* 合理化初始进程

* 通过多段构建减少镜像层级

- 把多行参数按字母排序，可以减少可能出现的重复参数，并且提高可读性
- 编写 dockerfile 的时候，应该把变更频率低的编译指令优先构建以便放在镜像底层以有效利用 build cache
- 复制文件时，每个文件应独立复制，这确保某个文件变更时，只影响改文件对应的缓存

