# Simple Diff
1. A simple diffing utility that will compare two file 
2. The output of the data are any data that is present within the primary file but not in the secondary file

## Installation 
```
go install github.com/CTWO2000/simple-diff/cmd/simple-diff@v1.0.1
```

## To publish
1. Commit and push the lastest update
2. Add tag to repository 
    ```bash
    git tag <version | v1.0.0>
    ```
3. Push tag
    ```bash 
    git push origin v1.0.0
    ```
4. Publish pacakage with `go list`
    ```bash
    GOPROXY=proxy.golang.org go list -m github.com/CTWO2000/simple-diff@v1.0.0
    ```


