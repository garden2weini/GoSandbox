# 京东商城沙箱及DataV API

## RUN JD VOP Sandbox
```
go run sandbox.go
curl -l http://localhost:9008/jd/product_pool_no?id=merlin
```

## Demo
### Restful
```
go run demo/rest/hello-rest.go
curl -l http://localhost:4200/api/tuts
curl -l http://localhost:4200/api/tuts/1
```

## Prepare
```
# HTTP request multiplexer
go get -u github.com/gorilla/mux
go get -u github.com/satori/go.uuid
```

## Try it
参考[MQTT规范][MQTT]

## 利用openssl生成sanbox.crt和sanbox.key文件
```
>openssl genrsa -out sandbox.key 2048
# 根据私钥生成公钥
>openssl rsa -in sandbox.key -out sandbox.key.public
# 根据私钥直接生成自签发的数字证书
>openssl req -new -x509 -key sandbox.key -out sandbox.crt -days 3650
>openssl pkcs12 -export -in sandbox.crt -inkey sandbox.key -out sandbox.p12
...Enter Export Password: helloworld
# Note: 从Oracle下载对应jdk的安全扩展包先(如 bcprov-ext-jdk15on-157.jar)
# 生成bks证书
>keytool -importkeystore -srckeystore sandbox.p12 -srcstoretype pkcs12 -destkeystore sandbox.bks -deststoretype bks -provider org.bouncycastle.jce.provider.BouncyCastleProvider -providerpath bcprov-ext-jdk15on-157.jar
...输入目标密钥库口令: helloworld
...输入源密钥库口令: helloworld
```

## REF
- [The Twelve-Factor App](https://12factor.net/zh_cn/)
- [Go指南](http://tour.studygolang.com)
- [Go语言从新手到大神：每个人都会踩的五十个坑](https://segmentfault.com/p/1210000009466285/read)

[MQTT]: https://help.aliyun.com/document_detail/30540.html?spm=a2c56.pc_iot_community_sale.landing.2.314952065S91sf "MQTT协议规范"
[ting]: https://dict.eudic.net/ting "每日英语听力 - 欧路词典"

### 坑1
1.生成证书的时候，要指定ip
keytool -genkey -v -alias client -keyalg RSA -storetype PKCS12 -keystore client.p12 -validity 36500
注意：first last name需要指定如下
    -ext SAN=ip:xxx.xxx.xxx.xxx
2.P12 BKS 转换的时候密码过长，报错
    java.security.KeyStoreException: java.io.IOException: Error initialising store of key store: java.security.InvalidKeyException: Illegal key size
    是国外对技术出口的限制，限定了密钥的长度，需要替换jdk里面的两个文件（jre\lib\security）：local_policy.jar，US_export_policy.jar，要下载jdk版本对应的，不然会出错！
JDK6的下载地址：http://www.oracle.com/technetwork/java/javase/downloads/jce-6-download-429243.html 
JDK7的下载地址：http://www.oracle.com/technetwork/java/javase/downloads/jce-7-download-432124.html 
JDK8的下载地址：http://www.oracle.com/technetwork/java/javase/downloads/jce8-download-2133166.html 
下载后解压，可以看到local_policy.jar和US_export_policy.jar以及readme.txt。 
如果安装了JRE，将两个jar文件放到%JRE_HOME%\lib\security下覆盖原来文件，记得先备份。 
如果安装了JDK，将两个jar文件也放到%JDK_HOME%\jre\lib\security下。 
原文链接：https://blog.csdn.net/itheimach/article/details/70188646