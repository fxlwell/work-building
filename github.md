# github 工具使用记录
## 1.工作git和个人git配置
```
  vim ~/.ssh/config 
  
  Host workgit
    User LevinLFX
    HostName github.com
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/id_rsa.pub

  Host mygit
    User levinfx
    HostName github.com
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/levin_rsa.pub
   
  ```
 ## 2.修改远程分支
 ```
 git remote rm {remte_name}
 gir remote add origin git@workgit:levinfx/{git_name}.git
 ```
