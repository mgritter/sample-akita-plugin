module github.com/mgritter/sample-akita-plugin

go 1.16

require (
	github.com/akitasoftware/akita-cli v0.0.0-20210821012226-9dc94c13580c
	github.com/akitasoftware/akita-ir v0.0.0-20210818150446-55531f1ef499
	github.com/akitasoftware/akita-libs v0.0.0-20210819233753-8617521ce7fe
)

replace github.com/google/martian/v3 v3.0.1 => github.com/akitasoftware/martian/v3 v3.0.1-0.20210608174341-829c1134e9de

replace github.com/akitasoftware/akita-cli => ../akita-cli
