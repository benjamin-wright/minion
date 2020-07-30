#!/bin/bash

export USERNAME=npm-build-agent
export PASSWORD=jofdfg4ry9423u9f4f
export EMAIL=unknown@example.com

/usr/bin/expect <<EOD
spawn npm adduser --registry=https://npm.ponglehub.co.uk --scope=@minion-ci --strict-ssl=false
expect {
  "Username:" {send "$USERNAME\r"; exp_continue}
  "Password:" {send "$PASSWORD\r"; exp_continue}
  "Email: (this IS public)" {send "$EMAIL\r"; exp_continue}
}
EOD