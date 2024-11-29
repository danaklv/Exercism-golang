package main

import (
	"ex/tasks2"
	"fmt"
)

func main() {
//   res := tasks2.SplitLogLine("section 1<*>section 2<~~~>section 3")
//   fmt.Println(res[1])


  result := tasks2.TagWithUserName([]string{
    "[WRN] User James123 has exceeded storage space.",
	"[WRN] Host down. User   Michelle4 lost connection.",
	"[INF] Users can login again after 23:00.",
	"[DBG] We need to check that user names are at least 6 chars long.",
     
}) 
fmt.Println(result)
}
