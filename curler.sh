#!/bin/bash

internalErrorURL=(
  "http://localhost:8080/internal"
  "http://localhost:8080/internal?currentUserID=1"
  "http://localhost:8080/internal?currentUserID=4"
  "http://localhost:8080/internal?currentUserID=4&name=Maria"
  "http://localhost:8080/internal?currentUserID=4&name=Nushi"
  "http://localhost:8080/internal?currentUserID=4&name=Mohammed"
  "http://localhost:8080/internal?currentUserID=4&name=Jose"
  "http://localhost:8080/internal?currentUserID=4&name=Wei"
)
publicErrorURL=(
  "http://localhost:8080/public"
  "http://localhost:8080/public?currentUserID=1"
  "http://localhost:8080/public?currentUserID=4"
  "http://localhost:8080/public?currentUserID=4&name=Maria"
  "http://localhost:8080/public?currentUserID=4&name=Nushi"
  "http://localhost:8080/public?currentUserID=4&name=Mohammed"
  "http://localhost:8080/public?currentUserID=4&name=Jose"
  "http://localhost:8080/public?currentUserID=4&name=Wei"
)
publicWithMidErrorURL=(
  "http://localhost:8080/mid"
  "http://localhost:8080/mid?currentUserID=1"
  "http://localhost:8080/mid?currentUserID=4"
  "http://localhost:8080/mid?currentUserID=4&name=Maria"
  "http://localhost:8080/mid?currentUserID=4&name=Nushi"
  "http://localhost:8080/mid?currentUserID=4&name=Mohammed"
  "http://localhost:8080/mid?currentUserID=4&name=Jose"
  "http://localhost:8080/mid?currentUserID=4&name=Wei"
)

function curlBinData {
  
    local res=$(curl -s -w "%{http_code}" $1)
    local body=${res::-3}
    local status=$(printf "%s" "$res" | tail -c 3)
    
    echo -e "$status\t\t\t|  $body"
}

echo "Internal Errors"
echo
echo -e "HTTP Status Code\t|\tRespoonse"
for i in ${!internalErrorURL[@]};
do
 curlBinData ${internalErrorURL[$i]}
done

echo
echo "Public Errors"
echo
echo -e "HTTP Status Code\t|\tRespoonse"
for i in ${!publicErrorURL[@]};
do
 curlBinData ${publicErrorURL[$i]}
done

echo
echo "Public With Middleware Errors"
echo
echo -e "HTTP Status Code\t|\tRespoonse"
for i in ${!publicWithMidErrorURL[@]};
do
 curlBinData ${publicWithMidErrorURL[$i]}
done

