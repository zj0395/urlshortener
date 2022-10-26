# UrlShorten 短链接服务
将长链接缩短为由[0-9a-zA-Z]组成的6位短链  

## 快速开始
1. 更改`config/*.tpl`文件中的mysql配置，并删除文件名中的`.tpl`
2. 在自己的mysql中建表，可用`go run meta/sql_generate.go`生成建表sql。如果想修改分表数量，可以修改`models.UrlShortenTableCnt`值和`models.AccessHistoryTableCnt`，默认都是2
3. 运行 `go run main.go`

### Curl
#### Create
```
curl --location --request GET 'http://{YOUR_HOST}/surl/create?url=https://baidu.com/'
```
#### Access
```
curl --location --request GET 'http://{YOUR_HOST}/{YOUR_CODE}'
```

### 自定义
#### 生成自己的唯一递增序列
1. 可以使用`go run meta/random_seq.go`生成自己的唯一序列，见`utils/shorten/defines.go`

#### 分表
1. 采用的mysql分表方式中，单表最大2千万个短链数据量，最大支持2590个分表
2. 默认分表数量为2，可在`model/url_shorten.go`中更改`UrlShortenTableCnt`变量

## 随机性测试
数字增加1时，生成的6位短链变化很大。测试结果显示，数量加1运行1,000,000次时，生成的短链中变化了5-6个字符;  
可自行测试`bash unit_test.sh`
```
        	changeCharNum	cnt
        	1	0
        	2	0
        	3	0
        	4	0
        	5	977621
        	6	22379
```
