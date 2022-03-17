# Orange-server V1

## Depolyment

### 1. config.json

The following configuration is for testnet

```buildoutcfg
{
  "chain": "eth",
  "chain_rpc": "http://172.168.3.22:7545",
  "sys_data_service": "did:ont:AS1QrBpgiPtPoggSU4YRyYNFBtCRnBMaDU",
  "file_path": "http://172.168.3.38:8080/files/",
  "avatar_file_path": "images/",
  "wasm_executor": {
    "did": "did:ont:ARNzB1pTkG61NDwxwzJfNJF8BqcZjpfNev",
    "address": "ARNzB1pTkG61NDwxwzJfNJF8BqcZjpfNev",
    "wallet": "./wallet.dat",
    "password": "<password>"
  },
  "db": {
    "UserName": "<name>",
    "Password": "<password>",
    "DBAddr": "<ip:port>",
    "DBName": "<db name>"
  },
  "did_config": [{
    "chain": "ont",
    "wallet": "./wallet.dat",
    "password": "<password>",
    "url": "http://polaris2.ont.io:20336",
    "DID": "did:ont:AXdmdzbyf3WZKQzRtrNQwAR91ZxMUfhXkt",
    "DIDContract": "4f7f159ac4b9913bb185fdf1895705f61b7d0cc6",
    "CredentialExpirationDays": 10,
    "gasprice": 2500,
    "gaslimit": 20000,
    "commit": false
  }],
  "ontlogin_config": {
    "chain": ["ONT","ETH"],
    "alg": ["ES256"],
    "serverInfo": {
      "name": "orange_server",
      "icon": "http://orangeicon.jpg",
      "url": "http://orange.io",
      "did": "did:ont:AXdmdzbyf3WZKQzRtrNQwAR91ZxMUfhXkt",
      "VerificationMethod": ""
    }
  },
  "outer_task_caller": ["<caller did>","<caller did2>"],
  "sig_auth": false,
  "mail_config": {
    "mail_address": "support@orangeprotocol.io",
    "host": "<host>",
    "smtp_port": 587,
    "password": "password",
    "subject": "Verification Code",
    "content":"<html>\n<head>\n   <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" />\n   <title>You verification code</title>\n</head>\n<body>\n   <p>{{VERIFICATION_CODE}}</p><br><p>Please do not reveal this code to others.</p>\n</body>"
  },
  "snapshot_assets_config": {
    "ap_did": "did:ont:testap",
    "ap_method": "snapshotStrageScore",
    "dp_did": "did:ont:AS1QrBpgiPtPoggSU4YRyYNFBtCRnBMaDU",
    "dp_method": "queryUserSnapShotAssets"
  },
  "eth_wallet": {
    "key_store_path": "./keystore",
    "password": "<password>"
  },
  "nft_config": {
    "nft_infos": {
      "eth": {
        "contract_address": "0x5f3c3ea1de47a2930ba8dbe436cf2ec5382b2584",
        "rpc": "https://speedy-nodes-nyc.moralis.io/6eb43157cbc67a17e7644196/eth/kovan"
      },
      "bsc": {
        "contract_address": "0xc19282f3d1cf8d70597283eb05f70c5bae198ce8",
        "rpc": "https://speedy-nodes-nyc.moralis.io/6eb43157cbc67a17e7644196/bsc/testnet"
      },
      "polygon": {
        "contract_address": "0xf3C7Ea39AC417cDa1867c290A348759545B9eC74",
        "rpc": "https://speedy-nodes-nyc.moralis.io/6eb43157cbc67a17e7644196/polygon/mumbai"
      }
    }
  },
  "graph_config": {
    "eth": "https://api.thegraph.com/subgraphs/name/orangeprotocol/orange-eth",
    "bsc": "https://api.thegraph.com/subgraphs/name/orangeprotocol/orange-bsc",
    "polygon": "https://api.thegraph.com/subgraphs/name/orangeprotocol/orange-mumbai"
  },

  "batch_task_count": 5,
  "task_timeout_seconds": 3600,
  "nft_timeout_minutes": 30
}

```

### 2. Start

```buildoutcfg
./orange-server-v1 
```
default port 8080 
--server-port to change for other port