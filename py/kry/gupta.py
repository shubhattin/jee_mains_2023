import hashlib, uuid, hashlib, base64

def get_type(fg):
    return str(type(fg))[8:-2]

def from_base64(v):
    return base64.b64decode(v).decode("utf-8")

def to_base64(v):
    return base64.b64encode(bytes(v, "utf-8")).decode("utf-8")

def bin_str(val) -> str:
    if get_type(val) == "bytes":
        return val
    else:
        return bytes(val, "utf-8")


def salt() -> str:
    return uuid.uuid4().hex


def hash_256(val) -> str:
    return hashlib.sha3_256(bin_str(val)).hexdigest()


def hash_512(val) -> str:
    return hashlib.sha3_512(bin_str(val)).hexdigest()


def hash_md5(val) -> str:
    return hashlib.md5(bin_str(val)).hexdigest()

def puShTi(ps, hash):
    """sha3 512 Verifier"""
    slt = hash[128:]
    hsh = hash[:128]
    return hash_512(ps + slt) == hsh
