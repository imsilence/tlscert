0. cfssl下载
    https://github.com/cloudflare/cfssl
    
1. 生成配置
    cfssl print-defaults config
    cfssl print-defaults csr

2. 生成根证书
    cfssl gencert -initca=true config/ca-csr.json  | cfssljson -bare cert/ca/ca

3. 生成服务端证书
    cfssl gencert --ca=cert/ca/ca.pem --ca-key=cert/ca/ca-key.pem --config=config/config.json --profile=server config/server-csr.json  | cfssljson -bare cert/server/server

4. 生成客户端证书
    cfssl gencert --ca=cert/ca/ca.pem --ca-key=cert/ca/ca-key.pem --config=config/config.json --profile=client config/client-csr.json  | cfssljson -bare cert/client/client

5. 生成peer证书(可使用于服务端或客户端)
    cfssl gencert --ca=cert/ca/ca.pem --ca-key=cert/ca/ca-key.pem --config=config/config.json --profile=peer config/peer-csr.json  | cfssljson -bare cert/peer/peer