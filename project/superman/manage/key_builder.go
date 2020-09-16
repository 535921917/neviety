package manage

import (
	"fmt"
	"shuai/superman/consts"
)


func KeyForServiceName(psm string) string {
	return fmt.Sprintf(consts.ServiceName, psm)
}

func KeyForServiceZSet(psm string) string {
	return fmt.Sprintf(consts.ServiceZSet, psm)
}

func KeyForServiceList(psm string) string {
	return fmt.Sprintf(consts.ServiceList, psm)
}

func KeyForServiceVersion(psm string) string {
	return fmt.Sprintf(consts.ServiceVersion, psm)
}

/*func KeyFor(psm string) string {
	return fmt.Sprintf(consts.ServiceName, psm)
}
*/
