#!/bin/bash
if [[ -z $1 ]]; then
  echo Please specify the output directory for the ca certificate.
  echo e.g. bash ./scripts/generate_ca.sh ./config/ca-cert
  exit 1
fi
mkdir -p ${1}/
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ${1}/ca-key.pem -out ${1}/ca-cert.pem -subj "/emailAddress=test@test.com"
#openssl x509 -in ${1}/ca-cert.pem -noout -text
