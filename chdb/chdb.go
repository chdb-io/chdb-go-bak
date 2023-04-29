package chdb

/*
#cgo CFLAGS: -I./ -I/use/include/python3.7
#cgo LDFLAGS: -L. -lpython3.7m -lchdb
#include "chdb.h"
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char *MyResult(char *query, char *format) {

    char dataQuery[200];
    struct local_result *myresult;
    char * argv[] = {(char *)"clickhouse", (char *)"--multiquery", (char *)"--output-format=CSV", (char *)"--query="};

    //Format
    snprintf(dataQuery, sizeof(dataQuery), "--format=%s", format);
    argv[2]=strdup(dataQuery);

    //Query
    snprintf(dataQuery, sizeof(dataQuery), "--query=%s", query);
    argv[3]=strdup(dataQuery);

    //Total 4 = 3 arguments + 1 programm name
    int argc = 4;

    myresult = query_stable(argc, argv);

    free(argv[2]);
    free(argv[3]);

    return myresult->buf;
}
*/
import "C"
import "unsafe"

func Query(str1, str2 string) string {
    query := C.CString(str1)
    defer C.free(unsafe.Pointer(query))
    format := C.CString(str2)
    defer C.free(unsafe.Pointer(format))
    resultData := C.MyResult(query, format)
    return C.GoString(resultData)
}
