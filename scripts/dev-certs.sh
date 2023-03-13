#!/usr/bin/env bash

CERTS_DIR=${PWD}/certs
CERTS_KEY_FILE=${CERTS_DIR}/server.key
CERTS_CRT_FILE=${CERTS_DIR}/server.crt
CERTS_CNF_FILE=${CERTS_DIR}/server.cnf
CERTS_PFX_FILE=${CERTS_DIR}/server.pfx

if [ -f $CERTS_KEY_FILE ]; then
    exit 0;
fi

echo "Generating local SSL certificate"
mkdir -p ${CERTS_DIR}

cat <<EOF > ${CERTS_CNF_FILE}
[ req ]

default_bits       = 2048
default_md = sha256

prompt             = no
string_mask        = default
distinguished_name = req_dn

x509_extensions = x509_ext

[ req_dn ]

countryName            = FR
stateOrProvinceName    = Nancy
organizationName       = Labyrinth
commonName             = Labyrinth

[ x509_ext ]

subjectKeyIdentifier    = hash
authorityKeyIdentifier  = keyid:always

keyUsage = critical, digitalSignature, keyEncipherment

extendedKeyUsage = serverAuth

subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
EOF

openssl req -x509 \
    -new \
    -nodes \
    -days 720 \
    -keyout ${CERTS_KEY_FILE} \
    -out ${CERTS_CRT_FILE} \
    -config ${CERTS_CNF_FILE}
    
openssl pkcs12 -export -out ${CERTS_PFX_FILE} -inkey ${CERTS_KEY_FILE} -in ${CERTS_CRT_FILE}