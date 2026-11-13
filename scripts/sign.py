from cryptography.hazmat.primitives.asymmetric import rsa, ec

key = rsa.generate_private_key(public_exponent=65537, key_size=2048)
ec_key = ec.generate_private_key(ec.SECP256R1())
