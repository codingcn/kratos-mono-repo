#!/bin/bash

cd app

kratos new kratos-mono-repo/app/user

rm -rf go.mod
rm -rf go.sum
rm -rf third_party