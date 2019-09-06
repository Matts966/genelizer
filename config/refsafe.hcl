rule "reflectcan" {
    package = "reflect"
    doc = "check can funtion is called"
    func "Interface" {
        receiver = "Value"
        before "CanInterface" {}
    }
    func "Addr" {
        receiver = "Value"
        before "CanAddr" {}
    }
    func "Set" {
        receiver = "Value"
        before "CanSet" {}
    }
}

rule "reflectkind" {
    package = "reflect"
    doc = "check kind"
    func "SetPointer" {
        receiver = "Value"
        before "Kind" { 
            return = [ "UnsafePointer" ]
        }
    }
    func "SetBool" {
        receiver = "Value"
        before "Kind" { 
            return = [ "Bool" ]
        }
    }
    func "SetString" {
        receiver = "Value"
        before "Kind" { 
            return = [ "String" ]
        }
    }
}
