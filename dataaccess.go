package dataaccess

func AppendParamString(key string, value string) (keyvalue string){
  keyvalue := "&K="+key+"|"+value
    return keyvalue
}

