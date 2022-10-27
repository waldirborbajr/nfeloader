#!/usr/bin/env bash
#
# https://www.golinuxcloud.com/golang-http/
#

openssl genrsa -out ca.key 4096

openssl req -new -x509 -days 365 -key ca.key -out cacert.pem -subj "/C=IN/ST=NSW/L=Bengaluru/O=GoLinuxCloud/OU=Org/CN=RootCA"

cat << EOF > server_cert_ext.cnf
basicConstraints = CA:FALSE
nsCertType = server
nsComment = "OpenSSL Generated Server Certificate"
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer:always
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName = @alt_names
[alt_names]
IP.1 = 172.20.10.2
DNS.1 = ubuntu
DNS.2 = localhost
EOF

# --- Step 2

openssl genrsa -out server.key 4096

openssl req -new -key server.key -out server.csr -subj "/C=IN/ST=NSW/L=Bengaluru/O=GoLinuxCloud/OU=Org/CN=ubuntu"

openssl x509 -req -in server.csr  -CA cacert.pem -CAkey ca.key -out server.crt -CAcreateserial -days 365 -sha256 -extfile server_cert_ext.cnf

# --- step 3

cp server.crt certbundle.pem

cat cacert.pem >> certbundle.pem

cat certbundle.pem

