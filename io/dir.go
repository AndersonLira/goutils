package io

import (
    "os"
    "strings"
    "path/filepath"
)

//ListFiles returns file names for giving root path. Extensions parameter filter files.
//All file names will returned if it is ommited
func ListFiles(root string, extensions []string) (files []string, err error){

    err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if len(extensions) == 0 {
            files = append(files, path)
            return nil
        }

        for _,ext := range extensions {
            if strings.HasSuffix(path, ext){
                files = append(files, path)
                return nil
            }
        }
        return nil
    })

    if err != nil {
        return
    }

    return
}
