## Keep an eye on the kids in the pool
Baywatch is *currently* an API exploration tool for BlueKai DMP instances.

Add your API keys and BK ID in a .env file as "BK_KEY" and
"BK_SECRET" and "BK_PARTNER_ID" respectively.

usage:

```
$ ./baywatch [ server ] | [ callType, method, endPoint ] <opt: data > < opt: format >
```
BK's ping example would translate to:

```
$ ./baywatch -callType=services -method=GET -endpoint=Ping
```
and a cookie count call can be written as:
```
$ ./baywatch -callType=taxonomy -method=GET -endPoint=categories -data="parentCategory.id=12345"
```

### note, example cookie count response formatting included, you should pepper to taste.


run as a server

```
$ ./baywatch -server
```

Here, urls follow this pattern:
http://server:port/callType/method/endPoint?data

e.g.:
```
localhost:8080/services/GET/Ping
localhost:8080/services/GET/sites
localhost:8080/taxonomy/GET/categories?parentCategory.id=399598
```

See [BlueKai's API Docs](https://kb.bluekai.com/display/PD/BlueKai+API+Docs) for ideas (login required)


@alvaromuir, Verizon Corporate Digital Media
