package utils

import (
	"log"
)

// 打印icon样式
func PrintLogo(b bool, docs string) {
	logo := `
__   ___   _ ____        __  __ ___ ____ ____   ___  
\ \ / / \ | |  _ \      |  \/  |_ _/ ___|  _ \ / _ \ 
 \ V /|  \| | | | |_____| |\/| || | |   | |_) | | | |
  | | | |\  | |_| |_____| |  | || | |___|  _ <| |_| |
  |_| |_| \_|____/      |_|  |_|___\____|_| \_\\___/ 
-----------------------------------------------------
%s
======================================================
`
	if b {
		log.Printf(logo, docs)
	}
}
