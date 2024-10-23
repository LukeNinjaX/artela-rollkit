import base64
import hashlib
from Crypto.Hash import RIPEMD160
from bech32 import bech32_encode, convertbits

# Base64解码
pubkey_base64 = "+SrfTnc7KUBAvp1s2LdutuDBc6kr6+nBCXtzPG7t4VM="
pubkey_bytes = base64.b64decode(pubkey_base64)

# SHA256哈希
sha256_hash = hashlib.sha256(pubkey_bytes).digest()

# 使用 pycryptodome 进行 RIPEMD160 哈希
ripemd160_hash = RIPEMD160.new()
ripemd160_hash.update(sha256_hash)
ripemd160_digest = ripemd160_hash.digest()

# Bech32编码，添加前缀
bech32_prefix = 'cosmos'
address_bytes = convertbits(ripemd160_digest, 8, 5)  # 将字节转换为5位字
bech32_address = bech32_encode(bech32_prefix, address_bytes)

# 输出结果
print(f"Bech32地址: {bech32_address}")
print(f"Hex地址: {ripemd160_digest.hex().upper()}")

