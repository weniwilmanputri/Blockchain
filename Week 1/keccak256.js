from base64 import encode
import sha3
import os

print("Keccak 256 di python\n")
NamaMakanan = input("Masukan Nama Makanan: ")
os.system('CLS')
print("Nama Makanan: \n", NamaMakanan)
encoded = NamaMakanan.encode()
obj_encoded = sha3.keccak_256(encoded)
print("Nama Makanan sesudah hash Keccak 256: \n", obj_encoded.hexdigest()
