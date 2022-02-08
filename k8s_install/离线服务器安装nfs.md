# 离线服务器安装nfs

## 安装方法和顺序

dpkg -i nfs-common_1.2.0-4ubuntu4.2_i386.deb

dpkg -i nfs-kernel-server_1.2.0-4ubuntu4.2_i386.deb

## 打开/etc/exports文件

sudo vim /etc/exports

在末尾加入/home *(rw,sync,no_root_squash)

其中：/home表示要共享的目录，*表示所有的网段，()里面表示可读写，资料同步写入内部磁盘，nfs客户端共享目录使用者权限

## 重启服务

sudo /etc/init.d/nfs-kernel-server restart

 

 **NFS简介**

NFS（Network File System）即网络文件系统，是FreeBSD支持的文件系统中的一种，它允许网络中的计算机之间通过TCP/IP网络共享资源。在NFS的应用中，本地NFS的客户端应用可以透明地读写位于远端NFS服务器上的文件，就像访问本地文件一样。

nfs服务是实现Linux和Linux之间的文件共享，nfs服务的搭建比较简单。

------

# **Ubuntu系统配置NFS服务器**

**1、安装nfs服务**

```
sudo apt install nfs-common nfs-kernel-server
```

**2、修改配置文件**

vim /etc/exports

```
/home/nfs/ *(rw,sync,no_root_squash)
```

各段表达的意思如下，根据实际进行修改

```
/home/nfs/ ：共享的目录
* ：指定哪些用户可以访问
 * 所有可以ping同该主机的用户
 192.168.1.* 指定网段，在该网段中的用户可以挂载
 192.168.1.12 只有该用户能挂载
(ro,sync,no_root_squash)： 权限
 ro : 只读
 rw : 读写
 sync : 同步
 no_root_squash: 不降低root用户的权限
```

**3、重启nfs服务**

sudo /etc/init.d/nfs-kernel-server restart

到这里，nfs的服务器就搭建好了。

------

# 客户端挂载NFS服务器

**1、检查客户端和服务端的网络是否连通（ping命令）**

ping + 主机IP

**2、查看服务端的共享目录**

showmount -e + 主机IP

```
showmount -e 172.18.186.160
```

**3、将该目录挂载到本地**

新建挂载点：mkdir -p /opt/nfs_test

挂载：mount 172.18.186.160:/home/nfs /opt/nfs_test

到这里就挂载成功了。

 