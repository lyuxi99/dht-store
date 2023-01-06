#!/bin/bash
if [[ -z $1 || -z $2 || -z $3 ]]; then
  echo Please specify the directory of the ca certificate, a directory for output, and the IP.
  echo e.g. bash ./scripts/generate_cert.sh ./config/ca-cert/ ./config/node1/ 127.0.0.1
  exit 1
fi

#$1 = $(realpath $1)
#$2 = $(realpath $2)
mkdir -p $2/

openssl req -newkey rsa:4096 -nodes -keyout ${2}/hostkey.pem -out ${2}/host-req.pem -subj "/emailAddress=test@test.com"
openssl x509 -req -in ${2}/host-req.pem -days 60 -CA ${1}/ca-cert.pem -CAkey ${1}/ca-key.pem -CAcreateserial -out ${2}/hostcert.pem \
  -extfile <(printf "subjectAltName=IP:${3}")
#openssl x509 -in ${2}/server-cert.pem -noout -text
