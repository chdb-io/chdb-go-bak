package chdb

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lchdb
#include "chdb.h"
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char *Execute(char *query, char *format) {

    char * argv[] = {(char *)"clickhouse", (char *)"--multiquery", (char *)"--output-format=CSV", (char *)"--query="};
    char dataFormat[100];
    char *localQuery;
    // Total 4 = 3 arguments + 1 programm name
    int argc = 4;
    struct local_result *result;

    // Format
    snprintf(dataFormat, sizeof(dataFormat), "--format=%s", format);
    argv[2]=strdup(dataFormat);

    // Query - 10 characters + length of query
    localQuery = (char *) malloc(strlen(query)+10);
    if(localQuery == NULL) {

        printf("Out of memmory\n");
        return NULL;
    }

    sprintf(localQuery, "--query=%s", query);
    argv[3]=strdup(localQuery);
    free(localQuery);

    // Main query and result
    result = query_stable(argc, argv);

    //Free it
    free(argv[2]);
    free(argv[3]);

    return result->buf;
}

char *Session(char *query, char *format, char* path) {

    char * argv[] = {(char *)"clickhouse", (char *)"--multiquery", (char *)"--output-format=CSV", (char *)"--query=", (char *)"--path=/tmp/"};
    char dataFormat[100];
    char dataPath[100];
    char *localQuery;
    // Total 4 = 4 arguments + 1 programm name
    int argc = 5;
    struct local_result *result;

    // Format
    snprintf(dataFormat, sizeof(dataFormat), "--format=%s", format);
    argv[2]=strdup(dataFormat);

    // Query - 10 characters + length of query
    localQuery = (char *) malloc(strlen(query)+10);
    if(localQuery == NULL) {

        printf("Out of memmory\n");
        return NULL;
    }

    sprintf(localQuery, "--query=%s", query);
    argv[3]=strdup(localQuery);
    free(localQuery);

    // Path
    snprintf(dataPath, sizeof(dataPath), "--path=%s", path);
    argv[4]=strdup(dataPath);

    // Main query and result
    result = query_stable(argc, argv);

    //Free it
    free(argv[2]);
    free(argv[3]);
    free(argv[4]);

    return result->buf;
}
*/
import "C"
import "unsafe"

func Query(str1 string, str2 string) string {
	if str2 == "" {
		str2 = "CSV"
	}
	query := C.CString(str1)
	defer C.free(unsafe.Pointer(query))
	format := C.CString(str2)
	defer C.free(unsafe.Pointer(format))
	resultData := C.Execute(query, format)
	return C.GoString(resultData)
}

func Session(str1 string, str2 string, str3 string) string {
	if str3 == "" {
		str3 = "/tmp/"
	}
	if str2 == "" {
		str2 = "CSV"
	}
	query := C.CString(str1)
	defer C.free(unsafe.Pointer(query))
	format := C.CString(str2)
	defer C.free(unsafe.Pointer(format))
	path := C.CString(str3)
	defer C.free(unsafe.Pointer(path))
	resultData := C.Session(query, format, path)
	return C.GoString(resultData)
}
