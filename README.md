## 层级说明

```base
├── interfaces 接入层 【http/grpc适配, 这一层调用application层; 比如interfaces层定义了输入层的相关方法，以使用gin提供http接口为例，这里的handler等为使用gin提供的一些http接口，这一层调用application层】
│   ├── grpc
│   └── http
│   └── facade  引用其他微服务（接口防腐层）
│   ├── event 事件
│   │   └── subscribe mq消费入口
│   ├── job 定时任务
├── application 应用层 【主要是调用domain层与infrastructure层来实现功能】
│   ├── assembler   负责将内部领域模型转化为可对外的DTO
│   └── dto Application层的所有接口返回值为DTO -- 入参/出参
│   └── adapter 负责业务流程的编排，但本身不负责任何业务逻辑
├── domain 领域层 【主要是定义了entity，service，以及service接口】
│   ├── entity 实体 业务逻辑。也可以参数校验，扩展一些简单方法，减轻service的压力
│   ├── event 事件
│   │   ├── publish 所有发送mq在此处理
│   │   └── subscribe 所有接受到mq处理逻辑在此处理
│   ├── service 领域服务 也是业务逻辑的核心
└── infrastructure 基础设施层 【这里提供了db层的实现，还有其他一些基础的组件，提供给application层或者interfaces层使用】
│   ├── config 配置文件
│   ├── consts 系统常量
│   ├── pkg 常用工具类封装（DB,log,util等）
│   └── repository 针对domain层的repository接口的实现
│   └── converter domain内对象转化 po {互转}
│   └── po 数据库映射对象
```


## DDD小结

DDD一般分为interfaces、application、domain、infrastructure这几层；其中domain层不依赖其他层，它定义repository接口，infrastructure层会实现；application层会调用domain、infrastructure层；interfaces层一般调用application层或者infrastructure层。

## 相关概念

- DDD等相关概念：https://domain-driven-design.org/zh/ddd-concept-reference.html

- VO（View Object）：视图对象，用于展示层，它的作用是把某个指定页面（或组件）的所有数据封装起来。

- DTO（Data Transfer Object）：数据传输对象，这个概念来源于J2EE的设计模式，原来的目的是为了EJB的分布式应用提供粗粒度的数据实体，以减少分布式调用的次数，从而提高分布式调用的性能和降低网络负载，但在这里，我泛指用于展示层与服务层之间的数据传输对象。

- DO（Domain Object）：领域对象(entity)，就是从现实世界中抽象出来的有形或无形的业务实体。

- PO（Persistent Object）：持久化对象，它跟持久层（通常是关系型数据库）的数据结构形成一一对应的映射关系，如果持久层是关系型数据库，那么，数据表中的每个字段（或若干个）就对应PO的一个（或若干个）属性。

> ```
> 用户发出请求（可能是填写表单），表单的数据在展示层被匹配为VO。
> 展示层把VO转换为服务层对应方法所要求的DTO，传送给服务层。
> 服务层首先根据DTO的数据构造（或重建）一个DO，调用DO的业务方法完成具体业务。
> 服务层把DO转换为持久层对应的PO（可以使用ORM工具，也可以不用），调用持久层的持久化方法，把PO传递给它，完成持久化操作。
> 对于一个逆向操作，如读取数据，也是用类似的方式转换和传递
> ```
