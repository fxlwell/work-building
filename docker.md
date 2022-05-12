```
docker pull image  拉去镜像
```

```
vim /etc/docker/daemon.json 修改国内镜像
{
 "registry-mirrors": ["https://registry.docker-cn.com"]
}
```

```
register 镜像，可搭建局域网，镜像地址
```

```
docker inpect container_name 查看容器详细信息
```
```
--volume-from 继承模式
docker valume inspect 目录名字
```
```
docker commmit -a="auth" -m="new image" mycontainer_name newimage 创建镜像
```

### DockerFile 常用指令
```
FROM scratch 空镜像
ADD 拷贝并解压
COPY 仅仅拷贝
CMD 会被docker run 的命令覆盖掉，Dockerfile中只能一条生效
ENTRYPOINT 不会被覆盖
WORKDIR 设置工作目录
VOLUME 一般不常用，常在docker run -v 指定

```
