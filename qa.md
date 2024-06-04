
## trpc http

无论是基于http 还是基于trpc，都需要用ip+port定位，因为协议底层是tcp，trpc可以一个ip+port实现多个service，不同server不同port；http可以根据请求路径不同，实现不同的服务

| 协议   | 请求服务方式     | 实现不同服务方式 |
| ---- | ---------- | -------- |
| http | ip+port+路径 | 不同路径     |
| trpc | ip+port+方法 | 不同方法     |

trpc相比于http的优势：跨语言，server端与client可以不同语言；配置更灵活（yaml）；内置多种插件；无需知道调用细节；一个服务既可以作为server，也可以作为client


trpc也可以实现泛http rpc


trpc插件的使用

- 修改yaml文件，添加对应配置项
- 导入对应的包



## 新闻文章删除链路升级

- 删文服务迁移
- 字段统一


多个平台使用同一个删除接口？？

平台ABC发布同一篇文章，那么不同平台是不是存储数据库不同？

A平台调用删文服务，BC平台受影响吗？如果BC平台不受影响，是不是只是屏蔽了A平台文章对外的外链访问？

文章删除完成之后，




