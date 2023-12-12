#include <stdio.h>
#include <stdlib.h>

#include "vec.h"

void vec_new(vec *v) {
    v->size = 0;
    v->capacity = VEC_INIT_CAPACITY;
    v->items = malloc(sizeof(void *) * v->capacity);
}

int vec_len(vec *v) {
    return v->size;
}

static void vec_rsz(vec *v, int capacity) {
    void **items = realloc(v->items, sizeof(void *) * capacity);
    if (items) {
        v->items = items;
        v->capacity = capacity;
    }
}

void vec_add(vec *v, void *item) {
    if (v->capacity == v->size) {
        vec_rsz(v, v->capacity * 2);
    }

    v->items[v->size++] = item;
}

void vec_set(vec *v, int index, void *item) {
    if (index >= 0 && index < v->size) {
        v->items[index] = item;
    }
}

void *vec_get(vec *v, int index) {
    if (index >= 0 && index < v->size) {
        return v->items[index];
    }

    return NULL;
}

void vec_del(vec *v, int index) {
    if (index < 0 || index >= v->size) {
        return;
    }

    v->items[index] = NULL;

    for (int i = index; i < v->size - 1; i++) {
        v->items[i] = v->items[i + 1];
        v->items[i + 1] = NULL;
    }

    v->size--;

    if (v->size > 0 && v->size == v->capacity / 4) {
        vec_rsz(v, v->capacity / 2);
    }
}

void vec_clr(vec *v) {
    v->size = 0;
}

void vec_free(vec *v) {
    free(v->items);
}

char* vec_str(vec *v) {
    char *str = malloc(sizeof(char) * 100);
    sprintf(str, "vec { size: %d, capacity: %d }", v->size, v->capacity);
    return str;
}