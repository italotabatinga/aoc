#pragma once

#define VEC_INIT_CAPACITY 4

typedef struct
{
    void **items;
    int size;
    int capacity;
} vec;

#define VEC_NEW(v) \
    vec v;       \
    vec_new(&v)
#define VEC_ADD(vec, item) vec_add(&vec, (void *)item)
#define VEC_SET(vec, id, item) vec_set(&vec, id, (void *)item)
#define VEC_GET(vec, type, id) (type) vec_get(&vec, id)
#define VEC_DEL(vec, id) vec_del(&vec, id)
#define VEC_LEN(vec) vec_len(&vec)
#define VEC_CLR(vec) vec_clr(&vec)
#define VEC_FREE(vec) vec_free(&vec)
#define VEC_STR(vec) vec_str(&vec)

void vec_new(vec *v);
int vec_len(vec *v);
void vec_add(vec *v, void *item);
void vec_set(vec *v, int index, void *item);
void *vec_get(vec *v, int index);
void vec_del(vec *v, int index);
void vec_clr(vec *v);
void vec_free(vec *v);
char *vec_str(vec *v);