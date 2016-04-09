## Keep an eye on the kids in the pool
Baywatch is a automated cookie pool threshold monitor for BlueKai DMP instances.

Written in Go.

Add your API keys and BK ID in a .env file as "BK_KEY" and
"BK_SECRET" and "BK_PARTNER_ID" respectively.

usage:

```
$ ./baywatch [base] [resource] [endpoint] <opt:data>
```
BK's ping example would translate to:

```
$ ./baywatch services GET Ping
```


### Updates 04.09.15
added server version

```
$ ./baywatch server
```
@alvaromuir, verizon national digital media
