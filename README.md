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
$ curl -s http://localhost:3001/blocks | jq

[
  {
    "Timestamp": 1537674920,
    "Data": "R2VuZXNpcyBCbG9jaw==",
    "PrevBlockHash": "",
    "Hash": "bpN93jZa5YYoINUMa+KA67ATy5FhRMK8hckeOZTgWd4=",
    "Transactions": []
  },
  {
    "Timestamp": 1537674920,
    "Data": "U2FtcGxlIGJsb2Nr",
    "PrevBlockHash": "bpN93jZa5YYoINUMa+KA67ATy5FhRMK8hckeOZTgWd4=",
    "Hash": "xH+GiFs3mHuonyNnMAHT7YXFsYy9pzT+u2kxcURRxDY=",
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
