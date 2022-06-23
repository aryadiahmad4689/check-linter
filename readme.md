# Check Linter

## _Mencari Code Di Repo Anda Dengan Mudah_


Check Linter Dibangun Untuk Mempermudah Anda Mencari Kode Yang Ingin Ditemukan.

## Tecknologi
 - [Golang]

## Instalation

```sh 
go get github.com/aryadiahmad4689/check-linter@v1.0.7
```
## How to use

Pertama Anda Harus Membuat Sebuah File Json Dengan Nama `check_linter.json`


_Contoh Isian_


 ```
{
    "SelectorExpression":[
        {
            "Sel" : "Println",
            "X" : "fmt"
        },
        {
            "Sel" : "Print",
            "X" : "fmt"
        }
    ],

    "ExprStmt":[
        {
            "Sel" :":="
        }
    ],

    "DecStmt":[
        {
            "Sel" :"var"
        }
    ]
}
 ```

 _Config Yang Tersedia_

| Declaration | Tersedia |
| ------ | ------ |
| SelectorExpression | [Sel:Println/Print/Sprintf/Printf, X:fmt] |
| ExprStmt | [Sel: :=] |
| DecStmt | [Sel: var] |

Perintah
``` 
check-linter -root_dir=example 
```