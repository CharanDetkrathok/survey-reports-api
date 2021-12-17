package timeZone

import "time"

func Init() {

	ict := time.Now().Local().Location()
	time.Local = ict

}