JWT 构成

      1. header头部
      2. payload负载
      3. 签证(signature)

header：

       1. 申明类型
       2. 申明加密算法
       example 声明 最后最后会加密：
       `{
          'typ': 'JWT',
          'alg': 'HS256'
        }`

payload:

       存放有效信息的地方，有效信息包括：
            1. 标准中注册的声明：
                      1.1 iss: jwt签发者
                      1.2 sub: jwt面向用户
                      1.3 aud: 接受jwt的一方
                      1.4 exp: 过期时间，必须大于签发时间
                      1.5 nbf: 什么时候jwt不可用
                      1.6 jat: jwt签发时间
                      1.7 jti: jwt唯一身份标识，一次性token，避免重放攻击
            2. 公共的声明
                       可以存放任何信息，但是不建议存放私密信息，因为客户端可以做部分解密
            3. 私有声明
                        提供消费者和提供者所共同定义的声明，不建议存放敏感信息，采用base64对称解密
             定义payload 最后会加密：

             `{
                "sub": "1234567890",
                "name": "John Doe",
                "admin": true
              }
`

签发：

    签证信息，三部分组成：
            1. header（base64加密后的）
            2. payload （base64加密后的）
            3. secret
            example：
            `var encodedString = base64UrlEncode(header) + '.' + base64UrlEncode(payload);

             var signature = HMACSHA256(encodedString, 'secret'); `
     这个部分需要base64加密后的header和base64加密后的payload使用.连接组成的字符串，然后通过header中声明的加密方式进行加盐secret组合加密，然后就构成了jwt的第三部分。
     注意：secret是保存在服务器端的，jwt的签发生成也是在服务器端的，
     secret就是用来进行jwt的签发和jwt的验证，所以，它就是你服务端的私钥，
     在任何场景都不应该流露出去。一旦客户端得知这个secret, 那就意味着客户端是可以自我签发jwt了。

流程：

    img/jwt认证.png

优点

    1. 因为json的通用性，所以JWT是可以进行跨语言支持的，像JAVA,JavaScript,NodeJS,PHP等很多语言都可以使用。
    2. 因为有了payload部分，所以JWT可以在自身存储一些其他业务逻辑所必要的非敏感信息。
    3. 便于传输，jwt的构成非常简单，字节占用很小，所以它是非常便于传输的。
    4. 它不需要在服务端保存会话信息, 所以它易于应用的扩展

安全相关

    1. 不应该在jwt的payload部分存放敏感信息，因为该部分是客户端可解密的部分。
    2. 保护好secret私钥，该私钥非常重要。
    3. 请使用https协议

