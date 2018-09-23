# Neil

* Terminal 1

```bash
$ cd neil
$ go install
$ $GOPATH/bin/neil

Neil is running...
Press Ctrl+C to stop. :)
API server is running on port 3001.
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
