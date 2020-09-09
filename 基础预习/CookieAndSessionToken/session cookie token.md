Cookie session  token

Cookie：

        用于客户端保存登入信息，4096字节限制，
Cookie 带来的问题：

        1. cookie 长度限制
        2. cookie 安全问题
session：

        根据cookie传送过来的cookie，一般是sessionID，来判断是谁登入。
        session用户每次登入需要把session 存放在服务器内存中，下次登入必须访问这个服务器才能够得到登入信息

session带来的问题：

       1. session保存的服务器，如果访问到另外提供相同服务的服务器，没有保存session，就需要重新登入
       2. 用户每次登入需要保存session到服务器内存，当用户过多的话，不宜与扩展
       3. 因为用户信息存在cookie里面，当我获取了这段cookie的话，就可以已该用户身份操作(没有超时的情况下)

Token：

    用户使用用户名密码来请求服务器
    服务器进行验证用户的信息
    服务器通过验证发送给用户一个token
    客户端存储token，并在每次请求时附送上这个token值
    服务端验证token值，并返回数据

Token 带来的问题：

    安全问题