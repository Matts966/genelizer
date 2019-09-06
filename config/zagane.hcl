// see passes/unstopiter in https://github.com/gcpug/zagane
rule "zagane" {
    package = "cloud.google.com/go/spanner"
    doc = "check if iterator is stopped"
    type "*RowIterator" {
        should = ["Do", "Stop"]
    }
}
