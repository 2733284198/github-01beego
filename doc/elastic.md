
#GET _search

```
{
  "query": {
    "match_all": {}
  }
}
```

GET /goods


put /goods

# 映射
GET /goods/_mapping

# 定义映射
PUT /goods/_mapping
{
"properites":{
"content":{
"type":"text",
"analyzer": "ik",
"search_analyzer": "ik"
}
}
}

GET /goods/

# 查询所有数据
GET /goods/_doc/_search

# 查询数据
POST /goods/_doc/_search
{
"query": {
"match": {
"content":"中国"
}
}
}

# 查询数据id - error
POST /goods/_doc/_search
{
"query": {
"id":111
}
}

# 增加数据

## 增加数据
POST /goods/_doc
{
"content":"我是中国人"
}

## 指定id
POST /goods/_doc/111
{
"content":"我是中国人冰水111"
}

### 修改
PUT /goods/_doc/111
{
"content":"我是中国人冰水111--222"
}

DELETE /goods/_doc/111
{

}



## 




