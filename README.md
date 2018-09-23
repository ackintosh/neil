# Neil

* Terminal 1

```bash
$ cd neil
$ go install
$ $GOPATH/bin/neil

Mining blocks...
Neil is running...
Press Ctrl+C to stop. :)
API server is running on port 3001.

Added new block: 3030306b8f15384d8ab65c50969b8ad89406a3867896710fa50a8c12f3297e44
Added new block: 30303088d55da740ef25367cc76997f619a1619d0817b78875861ad93a73e4db
...
...
...
```

* Terminal 2

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
    "Hash": "00000b9c434cc4e006a76974d6db550da2127fd5b5f2b100a772abe98965376b",
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
    "PrevBlockHash": "00000b9c434cc4e006a76974d6db550da2127fd5b5f2b100a772abe98965376b",
    "Hash": "000009e9b4e5d014df8f3ec641cdf8bebac7f832c89897df13101efeea5bed7f",
    "Transactions": [],
    "Nonce": 279745
  }
]
```

See [openapi.yml](https://github.com/ackintosh/neil/blob/master/openapi.yml) for further detail of the REST API.
