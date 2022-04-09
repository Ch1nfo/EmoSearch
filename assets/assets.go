package assets

import (
	"embed"
)

//go:embed  user-agents.txt wp.txt ASP.txt ASPX.txt DIR.txt JSP.txt MDB.txt PHP.txt
var Info embed.FS
