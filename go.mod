module github.com/nbit99/go-openw-sdk/v2

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/google/uuid v1.1.1
	github.com/nbit99/go-owcrypt v1.0.5
	github.com/nbit99/openwallet/v2 v2.0.10
	github.com/tidwall/gjson v1.3.5
)

//replace github.com/blocktree/go-owcdrivers => ../go-owcdrivers
//replace github.com/blocktree/openwallet => ../openwallet
