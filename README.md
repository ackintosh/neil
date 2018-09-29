# Neil

## Installing

```bash
$ cd neil
$ go install
```

## Running Neil

### Let's run 2 nodes locally

* Node1

By default, REST API server runs on `:3001` and WebSocket server runs on `:6001`.

```bash
$ $GOPATH/bin/neil

Mining blocks...
Neil is running...
Press Ctrl+C to stop. :)
WebSocket server for P2P communication is listening on ws://localhost:6001
REST API server is listening on http://localhost:3001
Added new block: 000000de8edd4697444158e16d949ff3c170a33a3b6e88dd730da86411cb2114
Added new block: 000005f9194626a90fb25df2a9cc4b9a91621693d606a020651d9800024ac880
...
...
...
```

* Node2

The port numbers daemon runs on can be specified via CLI arguments.

```bash
$ $GOPATH/bin/neil -apiPort=3003 -p2pPort=6006

Mining blocks...
Neil is running...
Press Ctrl+C to stop. :)
REST API server is listening on http://localhost:3003
WebSocket server for P2P communication is listening on ws://localhost:6006
Added new block: 00000071ea0f2c6b5f7cde2aebc14e9adfd5cd5cb38fbee87bce26cfab9f5ae6
...
...
...
```

### Obtain the chain which exists in the node1

```bash
$ curl -s http://localhost:3001/chain | jq

[
  {
    "Index": 0,
    "Timestamp": 1537720981,
    "PrevBlockHash": "",
    "Hash": "",
    "Transactions": [],
    "Nonce": 0
  },
  {
    "Index": 1,
    "Timestamp": 1537720981,
    "PrevBlockHash": "",
    "Hash": "000000de8edd4697444158e16d949ff3c170a33a3b6e88dd730da86411cb2114",
    "Transactions": [
      {
        "Sender": "Qm9i",
        "Recipient": "SXZhbg==",
        "Amount": 1,
        "Hash": "939b8cf5b8495b66cfce91eb261bec2c8f86fcee9d33e1af6bd21ebe43acb1c8"
      },
      {
        "Sender": "Qm9i",
        "Recipient": "SXZhbg==",
        "Amount": 2,
        "Hash": "2dcff73a5bac8dd3c14d7d125640bf805a6731f74e045f5a81373cdd91d4e80e"
      }
    ],
    "Nonce": 1033341
  },
  {
    "Index": 2,
    "Timestamp": 1537720983,
    "PrevBlockHash": "000000de8edd4697444158e16d949ff3c170a33a3b6e88dd730da86411cb2114",
    "Hash": "000005f9194626a90fb25df2a9cc4b9a91621693d606a020651d9800024ac880",
    "Transactions": [],
    "Nonce": 279745
  }
]
```

### Add peer to notify to another one that a new block is mined

```bash
$ curl -v http://localhost:3001/peers -d '{"Address": "ws://127.0.01:6006"}'

*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 3001 (#0)
> POST /peers HTTP/1.1
> Host: localhost:3001
> User-Agent: curl/7.60.0
> Accept: */*
> Content-Length: 33
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 33 out of 33 bytes
< HTTP/1.1 201 Created
< Date: Thu, 27 Sep 2018 23:36:52 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

### Then the nodes start communication

- Node1

```bash
[From: 6006] Hi, I've mined a new block! Its hash is: 00000ea815ce4e3ac511e4058f9af1b0d98f8806a58637c18a69eebf8adc8392
[From: 6006] Hi, I've mined a new block! Its hash is: 00000b4217a25261ed871f9f784dc008f8dea1590a11e1823943ef641c479ca4
Added new block: 000006a4b84e827c42a1b9286af7dccf84eaaedc4fcf0941d4adc681594811cb
Added new block: 00000988e261bfd264b6697efab046338a5d205ae75ae2a2a1af1eb99300cc78
Added new block: 000002ef56d438f4fde9b9215c59099107d0cb17cb02829ef81c47478fca8268
[From: 6006] Hi, I've mined a new block! Its hash is: 00000f330495165c416a932fd2f1a3b7413ec6bf2785ec2511032864ddb20e6d
Added new block: 00000facff4d62f7080456c2a3a712295dca835c8fceaa828d0a35135f5b22b0
```

- Node2

```bash
Added new block: 00000ea815ce4e3ac511e4058f9af1b0d98f8806a58637c18a69eebf8adc8392
Added new block: 00000b4217a25261ed871f9f784dc008f8dea1590a11e1823943ef641c479ca4
[From: 6001] Hi, I've mined a new block! Its hash is: 000006a4b84e827c42a1b9286af7dccf84eaaedc4fcf0941d4adc681594811cb
[From: 6001] Hi, I've mined a new block! Its hash is: 00000988e261bfd264b6697efab046338a5d205ae75ae2a2a1af1eb99300cc78
[From: 6001] Hi, I've mined a new block! Its hash is: 000002ef56d438f4fde9b9215c59099107d0cb17cb02829ef81c47478fca8268
Added new block: 00000f330495165c416a932fd2f1a3b7413ec6bf2785ec2511032864ddb20e6d
```

## REST API specification

See [openapi.yml](https://github.com/ackintosh/neil/blob/master/openapi.yml) for further detail of the REST API.
