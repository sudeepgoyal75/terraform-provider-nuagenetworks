#!/bin/bash

mkdir build

while read tag; do
  echo "processing $tag"
  rm -rf ./${tag:1}

  API_VERSION=${tag:1} monogen -L terraform -c monolithe.ini -f ~/git/vsd-api-specifications -b $tag
  if [ $? -eq 0 ]; then
    mkdir ${tag:1}
    mv terraform/nuagenetworks ${tag:1}/
    sed -i "s/terraform-providers\\/terraform-provider-nuagenetworks\\/nuagenetworks/tpretz\\/terraform-provider-nuagenetworks\\/${tag:1}\\/nuagenetworks/" terraform/main.go
    mv terraform/main.go ${tag:1}/
    cd ${tag:1}
    go build -o ../build/terraform-provider-nuagenetworks-${tag:1}
    cd ..
  fi
  rm -r terraform
done < <(git -C ~/git/vsd-api-specifications tag | grep '^r')

