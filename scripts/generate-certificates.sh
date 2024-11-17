#!/bin/bash

rm -rf certs
mkdir certs

cfssl selfsign -config scripts/cfssl.json --profile rootca "Katalog Root CA" scripts/csr.json | cfssljson -bare certs/root

cfssl genkey scripts/csr.json | cfssljson -bare certs/server
cfssl genkey scripts/csr.json | cfssljson -bare certs/client

cfssl sign -ca certs/root.pem -ca-key certs/root-key.pem -config scripts/cfssl.json -profile server certs/server.csr | cfssljson -bare certs/server
cfssl sign -ca certs/root.pem -ca-key certs/root-key.pem -config scripts/cfssl.json -profile client certs/client.csr | cfssljson -bare certs/client