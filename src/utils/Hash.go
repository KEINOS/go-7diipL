package utils

import (
	"crypto/md5" //nolint:gosec // not for cryptographic usage
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash/fnv"

	"golang.org/x/crypto/sha3"
	"golang.org/x/xerrors"
	"lukechampine.com/blake3"
)

// Hash returns the hashed value using the algorithm of algo.
//
// Supported hash
//
//   fnv1_32    (checksum usage, 4 byte, 8 char length)
//   fnv1_64    (checksum usage, 8 byte, 16 char length)
//   md5        (checksum usage, 16 byte, 32 char length)
//   sha2_256   (casual usage, 32 byte, 64 char length)
//   blake3_256 (experimental, 32 byte, 64 char length)
//   sha3_256   (recommended, 32 byte, 64 char length)
//   sha2_512   (recommended, 64 byte, 128 char length)
//   sha3_512   (recommended, 64 byte, 128 char length)
//   blake3_512 (experimental, 64 byte, 64 char length)
func Hash(algo string, value string) (string, []byte, error) {
	switch algo {
	case "blake3_256":
		valByte := blake3.Sum256([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	case "blake3_512":
		valByte := blake3.Sum512([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	case "fnv1_32":
		h := fnv.New32()
		_, err := h.Write([]byte(value))
		valByte := h.Sum(nil)

		return fmt.Sprintf("%x", valByte), valByte, err
	case "fnv1_64":
		h := fnv.New64()
		_, err := h.Write([]byte(value))
		valByte := h.Sum(nil)

		return fmt.Sprintf("%x", valByte), valByte, err
	case "md5":
		//nolint:gosec // not used for cryptographical purpose
		valByte := md5.Sum([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	case "sha2_256":
		valByte := sha256.Sum256([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	case "sha2_512":
		valByte := sha512.Sum512([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	case "sha3_256":
		valByte := sha3.Sum256([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	case "sha3_512":
		valByte := sha3.Sum512([]byte(value))

		return fmt.Sprintf("%x", valByte), valByte[:], nil
	}

	return "", nil, xerrors.New("Unsupported algorithm: " + algo)
}
