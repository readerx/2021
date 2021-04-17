#include <dirent.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>

int main(int argc, char *argv[]) {
    struct dirent *entry;
    DIR *dp;
    int max = 7;
    char buff[1024];

    fprintf(stderr, "time_t = %db\n", sizeof(time_t) * 8);

    dp = opendir("E:\\workspace\\temp");
    if (dp == NULL) {
        perror("opendir");
        return -1;
    }
    while (--max && (entry = readdir(dp))) {
        fprintf(stderr, "dir: %s\n", entry->d_name);

        struct stat st;
        memset(buff, 0, 1024);
        sprintf(buff, "%s\\%s", "E:\\workspace\\temp", entry->d_name);
        int ret = stat(buff, &st);
        if (ret < 0) {
            printf("stat error\n");
            return 0;
        }
        if (S_ISREG(st.st_mode)) {
            printf("print is file\n");
            FILE *f = fopen(buff, "r");
            int size = fread(buff, 1, 1024, f);
            if (size < 0) {
                printf("read error\n");
                return 0;
            }
            printf("read: %s\n", buff);
        }
    }

    closedir(dp);
    return 0;
}
