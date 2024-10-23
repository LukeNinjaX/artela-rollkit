#!/bin/bash

# Base64 公钥
pub_key_base64="+SrfTnc7KUBAvp1s2LdutuDBc6kr6+nBCXtzPG7t4VM="

# Step 1: 将 Base64 公钥解码为二进制
pub_key_bin=$(echo "$pub_key_base64" | base64 -d)

# Step 2: 计算 SHA-256 哈希
hash=$(echo -n "$pub_key_bin" | openssl dgst -sha256 -binary)

# Step 3: 取前 20 个字节
address=$(echo -n "$hash" | head -c 20 | xxd -p -c 20)

# Step 4: 转换为大写并输出
echo "${address^^}"

