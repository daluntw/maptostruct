# maptostruct

Convert ```map[string]interface{}``` to ```type struct```

## Install
```shell script
go get -t github.com/daluntw/maptostruct
```

## Usage

```go
package main

import "github.com/daluntw/maptostruct"

func main() {

    data := make(map[string]interface{})
    data["foo"] = "bar"

    type Struct struct {
        Foo string `json:"foo"`
    }

    var output Struct

    if err := maptostruct.MapToStruct(data, &output); err != nil {
        panic(err)
    }   

    print(output.Foo)
}

```