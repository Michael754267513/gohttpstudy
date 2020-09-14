个人理解：

    只是对权限的控制，不会对用户账号密码的保存控制，只是针对用户或者组去做权限控制
    通过用户登入用户set-cookie里面，保存用户的sessionID，来做用户下次登入的身份验证，通过中间件gin.context c.get获取用户是否存在。
    然后根据中间件传送的 资源名，动作和策略验证来判断是否有权限去操作
Casbin 可以：

    支持自定义请求的格式，默认的请求格式为{subject, object, action}。
    具有访问控制模型model和策略policy两个核心概念。
    支持RBAC中的多层角色继承，不止主体可以有角色，资源也可以具有角色。
    support built-in superuser like root or administrator. A superuser can do anything without explicit permissions.
    支持多种内置的操作符，如 keyMatch，方便对路径式的资源进行管理，如 /foo/bar 可以映射到 /foo*

Casbin 不能：

    身份认证 authentication（即验证用户的用户名、密码），casbin只负责访问控制。应该有其他专门的组件负责身份认证，然后由casbin进行访问控制，二者是相互配合的关系。
    管理用户列表或角色列表。 Casbin 认为由项目自身来管理用户、角色列表更为合适， 用户通常有他们的密码，但是 Casbin 的设计思想并不是把它作为一个存储密码的容器。 而是存储RBAC方案中用户和角色之间的映射关系。

Casbin 工作原理

    在 Casbin 中, 访问控制模型被抽象为基于 PERM (Policy, Effect, Request, Matcher)
    的一个文件。 因此，切换或升级项目的授权机制与修改配置一样简单。
    您可以通过组合可用的模型来定制您自己的访问控制模型。
    例如，您可以在一个model中获得RBAC角色和ABAC属性，并共享一组policy规则。

models：

    https://casbin.org/docs/zh-CN/supported-models


        