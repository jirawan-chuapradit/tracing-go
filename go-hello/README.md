Create certification self-signed

1. create private key
2. create Certificate Signing Request (CSR)
3. self-signed with CSR and private key

--------------------------------

1. สร้าง ca.key and ca.cert
openssl req -new -newkey rsa:2048 -keyout ca.key -x509 -sha256 -days 365 -out ca.crt

private key with encrypted with passphrase
self-signed certificate is saved in ca.crt

2. configuration file for the server certificate
this file will create cert that can connect to https://localhost without set InsecureSkipVerify  = true

3.  Generate server certificate using the self-signed CA
- create server key
openssl genrsa -out server.key 2048

- create CSR (Certificate Signing Request)
openssl req -new -key server.key -out server.csr -config server.cnf

- create server certificate with self signed CA
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key \
-CAcreateserial -out server.crt -days 365 -sha256 -extfile server.cnf -extensions v3_ext










