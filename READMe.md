# 偶遇问题

> go get 无反应、访问github.com速度慢、没反应问题的解决方案

    vim /etc/hosts
    192.30.253.112 github.com
    151.101.185.194 github.global.ssl.fastly.net

> go下载被墙掉的第三方库

    1. 使用gopm代替go下载
        
    //使用gopm(Go Package Manager)代替go下载,是go上的包管理工具，十分好用
    //1. 下载安装gopm
    go get -u github.com/gpmgo/gopm
    //2. 使用gopm安装被墙的包
    gopm get github.com/Shopify/sarama
        
    2. 使用镜像仓库
        
    golang 在 github 上建立了一个镜像库，如 https://github.com/golang/net 即是 https://golang.org/x/net 的镜像库.
    获取 golang.org/x/net 包（其他包类似），
    其实只需要以下步骤： 
       
    mkdir -p $GOPATH/src/golang.org/x
    cd $GOPATH/src/golang.org/x
    git　clone https://github.com/golang/net.git

    
    3. 使用国内网站打包
    http://www.golangtc.com

# go func practice

## 数据Array

- 定义数组的格式 var <varName> [n]<type>, n > 0
- 数组的长度也是类型的一部分，因此具有不同长度的数组为不同类型
- 注意区分指向数组的指针和指针数组
- 数组在GO中为值类型
- 数组之间可以使用 == 或 != 进行比较，但不可以使用 <或>
- 可以使用new来创建数组，此方法返回一个指向数组的指针
- GO支持多维数组


## 切片Slice

- 其本身不是数组，它指向[底层的数组]
- 作为边长数组的替代方案，可以关联底层的数组的局部或全部
- 为引用类型
- 可以直接创建或从底层数组获取生成
- 使用len()获取元素个数，cap()获取容量
- 一般使用make()创建
- 如果多个slice指向相同底层数组，其中一个的值改变会影响全部

- make([]T, len, cap)
- 其中cap可以省略，则和len的值相同
- len表示粗拉在的元素个数，cap表示容量

### Reslice

- Reslice 时索引以被slice的切片为准
- 索引不可以超过被slice的切片的容量cap()的值
- 索引越界不会导致底层数组的重新分配而是引发错误

### Append

- 可以在slice尾部追加元素
- 可以将一个slice追加在另一个slice的尾部
- 如果最终长度未超过追加到slice的容量则返回原始slice
- 如果超过追加到slice的容量则将重新分配数组并拷贝原始数据

### Copy

    复制
    copy(s1, s2)


### map

- 类似其他语言中的哈希表或者字典，以key-value形式存储数据
- Key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
- Map查找比现行搜索快很多，但比使用索引访问数据的类型慢100倍
- Map使用make()创建，支持:=这种简写方式

- make([keyType]valueType, cap)，cap表示容量，可省略
- 超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用len()获取元素的个数

- 键值对不存时自动添加，使用delete()删除某键值对
- 使用 for range 对map或slice进行迭代操作


### 函数function

- Go 函数不支持 嵌套、重载和默认参数
- 但支持一下特性：

    无需声明原型、不定长度变参、多返回值、命名返回值参数匿名函数、闭包

- 定义函数使用关键字func，且左大括号不能另起一行
- 函数也可以作为一种类型使用

### 结构 struct

- Go 中的struct与C中的struct非常相似，并且Go没有class
- 使用 type <Name> struct{} 定义结构，名称遵循可践行规则
- 支持指向自身的指针类型成员
- 支持匿名结构，可用作成员或定义成员变量
- 匿名结构也可以用于map的值
- 可以使用字面值对结构进行初始化
- 允许直接通过指针来读写结构成员
- 相同类型的成员可进行直接拷贝复制
- 支持 == 与 != 比较运算符，但不支持 > 或 <
- 支持匿名字段，本质上是定义了以某个类型名为名称的字段
- 嵌入结构作为匿名字段看起来像继承，但不是继承
- 可以使用匿名字段指针

### 方法 method

- Go 中虽没有class，但依旧有method
- 通过显示说明receiver来实现与某个类型的组合
- 只能为同一个包中的类型定义方法
- Receiver 可以是类型的值或者指针
- 不存在方法重载
- 从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第一个参数（Method Value vs. Method Expression）
- 如果外部结构或嵌入结构存在同名方法，则优先调用外部结构的方法
- 类型别名不会拥有底层类所附带的方法
- 方法可以调用结构中的非公开字段

### 接口 interface



### 反射 reflection


### 并发 concurrency


### 项目与坑




### 并发 concurency

- 从源码的解析来看，goroutine只是由官方实现的超级“线程池”而已
- 并发不是并行
- 并发主要有切换时间片来实现“同时”运行，在并行则是直接利用多核实现多线程的运行，但Go可以设置使用核数，以发挥多核计算机的能力。
- Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。


### Channel

- Channel 是 goroutine 沟通的桥梁，大都是阻塞同步的
- 通过make创建，close关闭
- Channel是引用类型
- 可以使用 for range 来迭代不断通过 channel
- 可以设置单项或双向通道
- 可以设置缓存大小，在未被填满前不会发生阻塞

### Select

- 可处理一个或多个channel的发送与接收
- 同时有多个可用的channel时按随机顺序处理
- 可用空的select来阻塞main函数
- 可设置超时

