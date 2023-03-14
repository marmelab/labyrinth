
CA_KEY_FILE=${CERTS_DIR}/ca.key
CA_CSR_FILE=${CERTS_DIR}/ca.csr
CA_CNT_FILE=${CERTS_DIR}/ca.cnt
CA_CRT_FILE=${CERTS_DIR}/ca.crt
CERTS_KEY_FILE=${CERTS_DIR}/server.key
CERTS_CSR_FILE=${CERTS_DIR}/server.csr
CERTS_CRT_FILE=${CERTS_DIR}/server.crt
CERTS_CNF_FILE=${CERTS_DIR}/server.cnf
CERTS_PFX_FILE=${CERTS_DIR}/server.pfx

if [ -f $CERTS_KEY_FILE ]; then
    exit 0;
fi

mkdir -p ${CERTS_DIR}

echo "Generating CA certificate"
openssl genrsa -out ${CA_KEY_FILE} 4096
openssl req -new -key ${CA_KEY_FILE} -out ${CA_CSR_FILE} -sha256 -subj '/CN=Labyrinth CA'

cat <<EOF > ${CA_CNT_FILE}
[ca]
basicConstraints = critical,CA:TRUE,pathlen:1
keyUsage = critical, nonRepudiation, cRLSign, keyCertSign
subjectKeyIdentifier=hash
EOF

openssl x509 -req -days 3650 -in ${CA_CSR_FILE} -signkey ${CA_KEY_FILE} -sha256 -out ${CA_CRT_FILE} -extfile ${CA_CNT_FILE} -extensions ca

echo "Generating server SSL certificate"

openssl genrsa -out ${CERTS_KEY_FILE} 4096
openssl req -new -key ${CERTS_KEY_FILE} -out ${CERTS_CSR_FILE} -sha256 -subj /CN=$SERVER_NAME

cat <<EOF > ${CERTS_CNF_FILE}
[ server ]

default_bits       = 2048
default_md = sha256

prompt             = no
string_mask        = default
distinguished_name = req_dn

x509_extensions = x509_ext

[ req_dn ]

countryName            = FR
stateOrProvinceName    = Nancy
organizationName       = Marmelab Labyrinth
commonName             = $SERVER_NAME

[ x509_ext ]

subjectKeyIdentifier    = hash
authorityKeyIdentifier  = keyid:always

keyUsage = critical, digitalSignature, keyEncipherment

extendedKeyUsage = serverAuth

subjectAltName = @alt_names

[alt_names]
DNS.1 = $SERVER_NAME
EOF

openssl x509 -req \
    -days 750 \
    -in ${CERTS_CSR_FILE} \
    -sha256 \
    -CA ${CA_CRT_FILE} \
    -CAkey ${CA_KEY_FILE} \
    -CAcreateserial \
    -out ${CERTS_CRT_FILE} \
    -extfile ${CERTS_CNF_FILE} \
    -extensions x509_ext

