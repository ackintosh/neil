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
    "Timestamp": 1537684382,
    "Data": "R2VuZXNpcyBCbG9jaw==",
    "PrevBlockHash": "",
    "Hash": "C5uEBPd/yvOqrH1qTSrJbJAIlj7q0Sl9iYidz5cKD6o=",
    "Transactions": []
  },
  {
    "Index": 1,
    "Timestamp": 1537684382,
    "Data": "U2FtcGxlIGJsb2Nr",
    "PrevBlockHash": "C5uEBPd/yvOqrH1qTSrJbJAIlj7q0Sl9iYidz5cKD6o=",
    "Hash": "/JV/BzOHipWOjMhlzU8aP6ip3ZCefxum0zjXdXQlENE=",
    "Transactions": [
      {
        "Sender": "Qm9i",
        "Recipient": "SXZhbg==",
        "Amount": 1
      },
      {
        "Sender": "Qm9i",
        "Recipient": "SXZhbg==",
        "Amount": 2
      }
    ]
  }
]
```

See [openapi.yml](https://github.com/ackintosh/neil/blob/master/openapi.yml) for further detail of the REST API.
