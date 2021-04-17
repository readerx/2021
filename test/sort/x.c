#include <pthread.h>
#include <stdio.h>

void print_arr(int *arr, int size) {
    for (int i = 0; i < size; i++) {
        printf("%d,", arr[i]);
    }
    printf("\n");
}

int s_sort(int *arr, int size) {
    for (int i = 1; i < size; i++) {
        int data = arr[i];
        for (int j = i - 1; j >= 0; j--) {
            if (arr[j] > data) {
                arr[j + 1] = arr[j];
                arr[j] = data;
            }
        }
    }
    return 0;
}

int main(int argc, char *argv[]) {
    int arr[] = {3, 4, 9, 7, 8, 5, 2};
    print_arr(arr, 7);

    s_sort(arr, 7);
    print_arr(arr, 7);

    return 0;
}
