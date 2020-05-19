# Introduction

> 帮助前端代理请求，以此来访问无法跨域的 api



## API

**Http GET**

- `/shProxyApi/v1/get`

**输入参数**

| 字段名 | 类型 | 必填 |     描述      |            备注             |
| :----: | :--: | :--: | :-----------: | :-------------------------: |
|  url   | TEXT |  Y   | 需要代理的URL | 此字段需要使用Base64URL编码 |

**示例**

```javascript
  // 请求百度翻译API
  const appid = "yourid";
  const salt = yoursalt;
  const key = "yourkey";
  const sign = md5(`${appid}${text}${salt}${key}`);

  const url = new URL("https://fanyi-api.baidu.com/api/trans/vip/translate");
  const params = url.searchParams;
  params.append("appid", appid);
  params.append("salt", salt);
  params.append("sign", sign);
  params.append("from", "auto");
  params.append("to", "zh");
  params.append("q", "text");
  const encodeHref = Base64.btoa(url.href);

  httpService.get(`http(s)://host:port/shProxyApi/v1/get?url=${encodeHref}`)
```

