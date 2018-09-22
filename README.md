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
    "Timestamp": 1537608825,
    "Data": "R2VuZXNpcyBCbG9jaw==",
    "PrevBlockHash": "",
    "Hash": "xKt+P2UbcN1ouK9Mv03K3QNOBP0PJSCVcEr000rk/pc="
  },
  {
    "Timestamp": 1537608825,
    "Data": "U2VuZCAxIEJUQyB0byBJdmFu",
    "PrevBlockHash": "xKt+P2UbcN1ouK9Mv03K3QNOBP0PJSCVcEr000rk/pc=",
    "Hash": "8fKKuciRH34TQueQn31JINNPKscdVFN1GNKERA5J3IU="
  },
  {
    "Timestamp": 1537608825,
    "Data": "U2VuZCAyIG1vcmUgQlRDIHRvIEl2YW4=",
    "PrevBlockHash": "8fKKuciRH34TQueQn31JINNPKscdVFN1GNKERA5J3IU=",
    "Hash": "rgbCEArxzlypgBO9h/X966iociYyX9FPaoBFDktrmMo="
  }
]
```
