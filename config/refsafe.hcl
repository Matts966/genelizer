rule "reflectcan" {
    package = "reflect"
    doc = "check can funtion is called"
    func "Interface" {
        before "CanInterface" {}
    }
    func "Addr" {
        before "CanAddr" {}
    }
    func "Set" {
        before "CanSet" {}
    }
}

rule "reflectkind" {
    package = "reflect"
    doc = "check kind"
    func "SetPointer" {
       before "Kind" { 
            return = [ "UnsafePointer" ]
        }
    }
    func "SetBool" {
       before "Kind" { 
            return = [ "Bool" ]
        }
    }
    func "SetString" {
       before "Kind" { 
            return = [ "String" ]
        }
    }
}
