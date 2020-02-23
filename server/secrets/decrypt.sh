#!/bin/bash

BaseDir="./server/secrets"
#BaseDir="."
PASSPHRASE=${LARGE_SECRET_PASSPHRASE} # from github env
Categories=(
"cloudsql"
"cloudstorage"
"firebase"
)

pwd

for CategoryDir in "${Categories[@]}" ;do
  PlaneDir="${BaseDir}/${CategoryDir}/plane"
  EncDir="${BaseDir}/${CategoryDir}/encrypted"
#  echo "${EncDir}"
  for EncKeyPath in "${EncDir}"/*; do
#    echo "${EncKeyPath}"
    KeyName=$(basename "${EncKeyPath}")
#    echo "${KeyName}"
    DecryptedKeyName=${KeyName%.*}
#    echo "${DecryptedKeyName}"
    DecryptedKeyPath=${PlaneDir}/${DecryptedKeyName}
    echo "${DecryptedKeyPath}"
    gpg --output "${DecryptedKeyPath}" --quiet --batch --yes --passphrase "${PASSPHRASE}" --decrypt "${EncKeyPath}"
  done
done
