package asciiart

import (
	"fmt"
)

func Show() {
	strArt := `
 -------------------------------------------------------------------------- 
████████╗██╗███╗   ███╗███████╗██╗  ██╗███████╗███████╗██████╗ ███████╗██████╗ 
╚══██╔══╝██║████╗ ████║██╔════╝██║ ██╔╝██╔════╝██╔════╝██╔══██╗██╔════╝██╔══██╗
   ██║   ██║██╔████╔██║█████╗  █████╔╝ █████╗  █████╗  ██████╔╝█████╗  ██████╔╝
   ██║   ██║██║╚██╔╝██║██╔══╝  ██╔═██╗ ██╔══╝  ██╔══╝  ██╔═══╝ ██╔══╝  ██╔══██╗
   ██║   ██║██║ ╚═╝ ██║███████╗██║  ██╗███████╗███████╗██║     ███████╗██║  ██║
   ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝
 ---- the-witcher-knight --------------------------------------------------`

	// Display the ASCII art
	fmt.Println(strArt)
}
