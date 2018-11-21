SHELL := /usr/bin/env bash
CERT_TTL?=365

.PHONY: gen-ca
gen-ca:
	mkdir -p certs/ca
	mkdir -p certs/newcerts
	touch certs/ca/serial
	touch certs/index.txt

	openssl genrsa -out certs/ca/ca.key 2048
	openssl req -new -x509 -days 3650 -key certs/ca/ca.key -out certs/ca/ca.crt -config ssl-ca.cnf -passout pass:password -subj "/C=RU/ST=Denial/L=Test/O=Test"

gen-server-key:
	openssl genrsa -out certs/server.key 2048

gen-server-cert-sign-req:
	openssl req -out certs/server.csr -key certs/server.key -new -config ./ssl-ca.cnf -subj "/C=RU/ST=Denial/L=Test/O=Test/CN=server.application"

gen-server-signed-cert:
	openssl ca -batch -config ssl-ca.cnf -name CA_signing -out certs/server.crt -infiles certs/server.csr

gen-server-cert:gen-server-key gen-server-cert-sign-req gen-server-signed-cert

gen-client-key:
	openssl genrsa -out certs/client.key 2048

gen-client-cert-sign-req:
	openssl req -out certs/client.csr -key certs/client.key -new -config ./ssl-ca.cnf -subj "/C=RU/ST=Denial/L=Test/O=Test/CN=client.application"

gen-client-signed-cert:
	openssl ca -batch -config ssl-ca.cnf -name CA_signing -out certs/client.crt -infiles certs/client.csr

gen-client-cert:gen-client-key gen-client-cert-sign-req gen-client-signed-cert