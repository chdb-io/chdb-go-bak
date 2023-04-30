package chdb

/*
#cgo pkg-config: python3
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lchdb
#include "chdb.h"
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char *MyResult(char *query, char *format) {

    char * argv[] = {(char *)"clickhouse", (char *)"--multiquery", (char *)"--output-format=CSV", (char *)"--query="};
    char dataFormat[100]; 
    char *localQuery;
    //Total 4 = 3 arguments + 1 programm name
    int argc = 4;
    struct local_result *myresult;    

    //Format
    snprintf(dataFormat, sizeof(dataFormat), "--format=%s", format);
    argv[2]=strdup(dataFormat);

    //Query - 10 characters + length of query
    localQuery = (char *) malloc(strlen(query)+10);
    if(localQuery == NULL) {
    
        printf("Out of memmory\n");
        return NULL;
    }
    
    sprintf(localQuery, "--query=%s", query);
    argv[3]=strdup(localQuery);
    free(localQuery);

    //Main query and result
    myresult = query_stable(argc, argv);

    //Free it
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
