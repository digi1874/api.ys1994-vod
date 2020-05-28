# 视频接口
* [前端网站](https://www.ys1994.nl/)
* [前端项目](https://github.com/digi1874/www.ys1994)

## 构建
> 依赖
> 1. go 1.13+
> 2. mysql (4.1+，本项目开发时使用5.7；库需要设置为utf8mb4)
> 3. 项目根目录下创建文件db.json连接mysql数据库
>> ./db.json
>> ```
>> {
>>   "user": "用户名",
>>   "password": "密码",
>>   "localhost": "地址",
>>   "databaseName": "库名"
>> }
>> ```
>> #
> ```
> # 开发，开启http://localhost:8051/
> $ go run main.go -env=dev
>
> # 生产程序
> $ go build
> ```
> #

## 接口说明
> * response status code:
>> 1. 200: 确定
>> 2. 400: 错误
>> 3. 401: 无权限，token无效
>> 4. 404: 不存在
> * response data msg: 回应说明
> #
> 1. 视频列表
>> * url：/vod/list?page=1&size=20&filter={}&sync=1
>> 1. page: 页数，非必需
>> 2. size: 每页数量，非必需
>> 3. filter: 过滤，非必需；filter={name:片名, area:地区, lang:语言, ...}
>> 4. sync: 同步源，非必需，filter.name不为空才有效
>> * method: GET
>> * response data: {"data":{"count": 0,"page": 1,"size": 20,"data": []}}
>> 1.
>> ```
>> // get: /vod/list
>> // response data:
>> {
>> "data": {
>>    "count": 60,
>>    "page": 1,
>>    "size": 20,
>>    "data": [
>>        {
>>            "id": 46,
>>            "actor": "金喜爱,朴海俊,朴善英,金永敏,李璟荣,金宣敬,蔡国熙,韩素希,李学周,沈恩宇",
>>            "area": "韩国",
>>            "serial": "更新至16集",
>>            "director": "毛完日",
>>            "name": "夫妻的世界",
>>            "pic": "https://images.cnblogsc.com/pic/upload/vod/2020-03/1585336717.jpg",
>>            "year": 2020,
>>            "typeId": 15
>>        },
>>        ...
>>    ]
>> }
>> ```
>> #
>
> 2. 视频详情
>> * url：/vod/detail/:id
>> * method: GET
>> * response data: {"data":{...}}
>> 1.
>> ```
>> // get: /vod/list
>> // response data:
>> {
>> "data": {
>>    "id": 46,
>>    "actor": "金喜爱,朴海俊,朴善英,金永敏,李璟荣,金宣敬,蔡国熙,韩素希,李学周,沈恩宇",
>>    "area": "韩国",
>>    "serial": "更新至16集",
>>    "director": "毛完日",
>>    "name": "夫妻的世界",
>>    "pic": "https://images.cnblogsc.com/pic/upload/vod/2020-03/1585336717.jpg",
>>    "year": 2020,
>>    "typeId": 15,
>>    "subName": "The Married Life",
>>    "lang": "韩语",
>>    "content": "该剧...",
>>    "typePId": 2,
>>    "updatedTime": 1590548043,
>>    "urls": [
>>        {
>>            "id": 578,
>>            "name": "第01集",
>>            "url": "youku.cdn11-okzy.com/share/03e03424a898e574153a10db9a4db79a"
>>        },
>>        ...
>>    ],
>>    "m3u8s": [
>>        {
>>            "id": 565,
>>            "name": "第01集",
>>            "url": "youku.cdn11-okzy.com/20200327/14550_8cb43c22/index.m3u8"
>>        },
>>        ...
>>    ],
>>    "downURLs": [
>>        {
>>            "id": 186,
>>            "name": "第01集",
>>            "url": "http://okxzy.xzokzyzy.com/20200327/14550_8cb43c22/夫妻的世界01.mp4"
>>        },
>>        ...
>>    ]
>> }
>> ```
>> #
