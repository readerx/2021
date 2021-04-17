#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

#define CHUNK (128 * 1024)

void usage(const char *prog, FILE *out) {
    fprintf(out, "usage: %s allocsize\n", prog);
    fprintf(out, " allocsize is kbytes, or number[KMGP] (P = pages)\n");
    exit(out == stderr);
}

int main(int argc, char *argv[]) {
    unsigned long long kbtotal = 0, kballoc, numchunk, alloc;
    unsigned long long i, j, k;
    char **mem, *tmp;
    int sum;

    if (argc == 2) {
        char *end = NULL;
        kbtotal = strtoull(argv[1], &end, 0);

        switch (*end) {
            case 'g':
            case 'G':
                kbtotal *= (1024 * 1024);
            case 'm':
            case 'M':
                kbtotal *= 1024;
            case '\0':
            case 'k':
            case 'K':
                break;
            case 'p':
            case 'P':
                kbtotal *= 4;
                break;
            default:
                usage(argv[0], stderr);
                break;
        }
    }

    if (argc != 2 || kbtotal == 0) {
        usage(argv[0], stderr);
    }

    numchunk = (kbtotal + CHUNK - 1) / CHUNK;
    mem = calloc(numchunk, sizeof(*mem));
    if (mem == NULL) {
        fprintf(stderr, "error allocating initial chunk array\n");
        exit(-1);
    }

    alloc = CHUNK;
    printf("[%d] allocating %lld kbytes in %llu kbyte chunks\n", getpid(), kbtotal, alloc);
    for (i = kballoc = 0; i < numchunk && alloc > 0; i++, kballoc += alloc) {
        if (kbtotal - kballoc < alloc) {
            alloc = kbtotal - kballoc;
        }

        while (alloc > 0 && (mem[i] = malloc(alloc * 1024)) == NULL) {
            fprintf(stderr, "malloc(%llu) failed (%lld/%lld)\n", alloc * 1024, kballoc, kbtotal);
            alloc /= 2;
        }
        if (alloc == 0) {
            break;
        }

        printf("touching %p ([%lld-%lld]/%lld)\n", mem[i], kballoc, kballoc + alloc - 1, kbtotal);
        for (j = 0, tmp = mem[i]; j < alloc; j += 4) {
            for (k = 0, sum = 0; k < 4095; k++, tmp++) {
                sum += *tmp;
            }
            *tmp = sum;
        }
    }

    if (kballoc == 0) {
        exit(-2);
    }

    printf("touched %lld kbytes\n", kballoc);

    printf("sleeping...\n");
    sleep(60);

    printf("free memory\n");
    for (i = 0; i < numchunk; i++) {
        free(mem[i]);
    }

    printf("sleeping...\n");
    sleep(30);

    return 0;
}
