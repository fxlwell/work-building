# YUM
## mysql 8
```
yum update -y

wget -i -c https://dev.mysql.com/get/mysql80-community-release-el7-3.noarch.rpm

yum -y install mysql80-community-release-el7-3.noarch.rpm

yum module disable mysql

yum -y install mysql-community-server


systemctl start mysqld

查看默认密码
grep "password" /var/log/mysqld.log

mysql -uroot -p 

创建权限记录：
CREATE user ‘root’@’%’ IDENTIFIED BY ‘您的密码’;

授权：
GRANT ALL PRIVILEGES ON . TO ‘root’ @’%’ WITH GRANT OPTION;

修改密码过期策略：
ALTER USER ‘root’ @‘localhost’ IDENTIFIED BY ‘您的密码’ PASSWORD EXPIRE NEVER;

重新修改密码：
ALTER USER ‘root’ @’%’ IDENTIFIED WITH mysql_native_password BY ‘您的密码’;

刷新权限：
FLUSH PRIVILEGES;


mysql服务相关命令
启动服务：service mysqld start
停止服务：service mysqld stop
重启服务：service mysqld restart
查看服务状态：service mysqld status

设置mysql开机自启：
systemctl enable mysqld
systemctl daemon-reload
查看自启服务
systemctl list-units --type=service

```