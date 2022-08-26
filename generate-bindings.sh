#!/bin/sh
# Copyright 2022 Stichting ThingsIX Foundation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0


set -e

THINGS_TOKEN_SMART_CONTRACT_REPO=git@github.com:ThingsIXFoundation/smart-contracts-evm.git

generate_bindings() {
  local work_dir=$1
  local pkg=$2
  local artifact_file=$(find $work_dir/artifacts -name "$3.json")
  local abi_file="$work_dir/$3.abi"

  # extract abi from artifacts file
  jq .abi "$artifact_file" > "$abi_file"

  echo "Generate bindings for package $pkg"
  
  abigen --abi "$abi_file" --pkg $pkg --out "bindings.go"
}

# clone smart contract repo and install dependencies
work_dir=$(mktemp -d /tmp/thingsix-tokens-XXXXX)
echo "Clone smart contract repo in work directory $work_dir"
git clone $THINGS_TOKEN_SMART_CONTRACT_REPO $work_dir --quiet

# generate go-bindings for gateway registry contract
generate_bindings $work_dir gateway_registry GatewayRegistry

# cleanup work directory
echo "delete work directory"
rm -rf $work_dir

# install dependencies
echo "tidy go dependencies"
go mod tidy