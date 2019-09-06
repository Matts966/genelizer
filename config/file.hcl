rule "file" {
    package = "os"
    doc = "check if file is closed"
    type "File" {
        should = [ "Close" ]
    }
}
