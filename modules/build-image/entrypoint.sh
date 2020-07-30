#!/bin/sh

npm set registry https://npm.ponglehub.co.uk

export USERNAME=npm-build-image
export PASSWORD=jofdfg4ry9423u9f4f
export EMAIL=unknown@example.com

/usr/bin/expect <<EOD
spawn npm adduser --scope=@minion-ci
expect {
  "Username:" {send "$USERNAME\r"; exp_continue}
  "Password:" {send "$PASSWORD\r"; exp_continue}
  "Email: (this IS public)" {send "$EMAIL\r"; exp_continue}
}
EOD

echo "$@"

eval $@