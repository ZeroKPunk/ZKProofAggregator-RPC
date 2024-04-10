#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct VecWrapper_u8 {
  const uint8_t *data;
  uintptr_t len;
  uintptr_t cap;
} VecWrapper_u8;

typedef struct VecWrapper_u8 Bytes;

typedef struct VecWrapper_Bytes {
  const Bytes *data;
  uintptr_t len;
  uintptr_t cap;
} VecWrapper_Bytes;

typedef struct VecWrapper_Bytes Siblings;

typedef struct Proof {
  uintptr_t index;
  Bytes value;
  Siblings siblings;
  Bytes root;
} Proof;

struct Proof generate_sparse_merkle_proof(const uint8_t *const *keys,
                                          uintptr_t keys_len,
                                          const uint8_t *const *values,
                                          uintptr_t values_len,
                                          uintptr_t index);
