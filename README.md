## 项目说明
* 该项目为看车玩车数据中心
* 支持三种链接方式: http(restful), RPC, TCP(目前支持RPC)
* 数据驱动支持: ES, mysql, mongodb, redis(rabbit mq)


* 使用golang版本 >12.5
* go mod vendor

## 开始

* 参数说明

    |  **名称**   | 说明  |类型  | 备注  |
    |  ----  | ----  | ---  |---  |
    | driver  | 驱动名称,支持 mysql/mongodb/redis/es|string | 默认mysql  |
    | table  | 数据表名 |string| **必填**,不需要表前缀  |
    | where  | 条件 |array| 支持多层嵌套  |
    | field  | 查询字段 |string| 比如:"id,name" |
    | limit  | 查询多少条 |int| 默认1  |
    | offset  | 查询起始 |int|默认为0，用于分页  |
    | order  | 排序 |array| 默认 "id asc"  |
    | group  | 分组 |array| 默认  |
    | join  | 链接 |array| 默认  |
    | sum  | 统计 |array| 默认  |
* 操作符
    > =、in、not in、>、>=、<、<=、like、field_in_set、not between、between
* 连接符
    > and、or    
* 状态码说明

    |  **code**   | 说明 | 备注  |
    |  ----  | ----  |---  |
    | 0  | 成功 |  |
    | 400  | 参数错误  |
    | 500  | 执行错误  |
    | 10000  | 表名不存在 |
    | 10002  | 操作符错误 | |
    | 10003  | 排序错误| |
    | 20001  | 数据存在| |
    | 20002  | 数据不存在|   |
    | 20004  | 创建数据失败| |
###mysql    
* 查询数据(一条或多条)  
```php
$client = new client()
$req = [
                 "driver"=>"mysql",
                 "table"=>"admin_user",
                 "where"=>["id", ">", 1],//id为field, ">" 操作符, 1 数值
                 "field"=>"id,username,sex",
                 "limit"=>10,
                 "offset"=>1,
                 "order"=>[["id", "desc"], ["create_time", "desc"]],
                 "group"=>["id", "create_time"],
                 "sum"=>["id", "create_time"],//统计
             ];
$ret = $client->Fetch(json_encode($req)); //注意全部是json格式请求
  /**
    where条件简单举例:
    A = ["id", ">", 1]
    B = ["username", "=", "hercule"]
    C = ["token", "in", ["dasdas","dasdsad"]]
    D = ["nickname", "field_in_set", "ddd"]
    E = ["realname", "like", "%刘德华"]

  "where"=>A ===> sql: where id > 1
  "where"=>[A, B, C] ===> A and B and C,**_注意第一层,and操作符可以省略,嵌套的and不能省略_**
  "where"=>["or", A, B, C] ==>  A or B or C
  "where"=>[ A, B, ["or", C,D]] ==>  A and B and ( C Or D)
  "where"=>[ A, ["or", B,C, ["and", D, E]]] ===>A and ( B or C or (D and E) ) //**嵌套里面and不能省略**
  
  
  */
  结果:
  {
      "code": 0,
      "msg": "ok",
      "data": [
          {
              "id": 29529,
              "is_digest": 0,
              "is_insert": 3,
              "is_pass": 0,
              "is_top": 0,
              "is_twitter": 0,
              "visible": 1
          }
      ]
  }
  
  
```
* 增加一条或多条数据
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "body"=>[
                ["sex"=>1, "username"=>"test"],
                ["sex"=>2, "username"=>"test2"],
             ],
       ];
$ret = $client->Create(json_encode($req));
*/
  结果:
  {
      "code": 0,
      "msg": "ok",
      "data": [
          29529,29528
      ]
  }
  
```
* 删除数据
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "where"=>["id", ">", 1],
       ];
$ret = $client->Delete(json_encode($req));
*/
  结果:
 {
      "code": 0,
      "msg": "ok",
      "data": {}
  }
```
* 修改数据
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "where"=>["id", ">", 1],
             "body"=>["username"=>"sss", "sex"=>1]
       ];
$ret = $client->Update(json_encode($req));
 */
   结果:
 {
      "code": 0,
      "msg": "ok",
      "data": {}
  }
```
* 查询数据是否存在
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "where"=>["id", ">", 1],
       ];
$ret = $client->Exist(json_encode($req));
 {
      "code": 0,
      "msg": "ok",
      "data": true/false
  }
```
* 查询数据条数
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "where"=>["id", ">", 1],
       ];
$ret = $client->Count(json_encode($req));
 {
      "code": 0,
      "msg": "ok",
      "data": 100
  }
```
###通过id进行删改查
* 查询数据
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "id"=>1,
             "field"=>"id,username,sex",
       ];
$ret = $client->Get(json_encode($req));
{
      "code": 0,
      "msg": "ok",
      "data": {
                "id": 29529,
                "is_digest": 0,
                "is_insert": 3,
                "is_pass": 0,
                "is_top": 0,
                "is_twitter": 0,
                "visible": 1
        }
  }
  
```

* 修改数据
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "id"=>1,
             "body"=>[
                "username"=>"ssss",
                "sex"=>1,
             ],
       ];
$ret = $client->Set(json_encode($req));
{
      "code": 0,
      "msg": "ok",
      "data": {}
  }
  
```
* 删除数据
```php
$client = new client()
$req = [
             "driver"=>"mysql",
             "table"=>"admin_user",
             "id"=>1,
       ];
$ret = $client->Del(json_encode($req));
{
      "code": 0,
      "msg": "ok",
      "data": {}
  }
  
```
###连接操作
```php
$client = new client()
$req = [
                 "driver"=>"mysql",
                 "table"=>"admin_user as user", //当取别名的时候 as 不能省略
                 "where"=>["id", ">", 1],
                 "field"=>"id,username,sex",
                 "join"=>[
                 ["INNER", "group", "group.id = user.group_id"],
                 ["INNER", "group as g", "g.id = u.group_id"]
                 ]
                 //可以多个链接,,如果是别名 "group as g" 同样 as 不能省略
                 "limit"=>10,
                 "offset"=>1,
                 "order"=>[["id", "desc"], ["create_time", "desc"]],
                 "group"=>["id", "create_time"],
             ];
$ret = $client->Fetch(json_encode($req)); //注意全部是json格式请求
  
  */