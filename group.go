package gocfg

//
//type ConfigItem struct {
//	Group      string `json:"group"`
//	Filename   string `json:"filename"`
//	ExtType    string `json:"ext_type"`
//	DriverName string `json:"driver_name"`
//}
//
//func (c ConfigItem) GetKey() string {
//	return fmt.Sprintf("%s-%s-%s", c.Group, c.ExtType, c.DriverName)
//}
//
//type GroupItem struct {
//	Group      string   `json:"group"`
//	Files      []string `json:"files"`
//	ExtType    string   `json:"ext_type"`
//	DriverName string   `json:"driver_name"`
//}
//
//func (c GroupItem) GetKey() string {
//	// TODO 内容寻址暂只支持固定group、driverName、ext扩展实现的固定配置查询，暂不能通过key自动实现配置寻址
//	return fmt.Sprintf("%s-%s-%s", c.Group, c.ExtType, c.DriverName)
//}
//
//func ToGroupItemList(items []ConfigItem) []GroupItem {
//	mm := map[string]*GroupItem{}
//	for _, item := range items {
//		// group和driver唯一确定配置
//		key := item.GetKey()
//		if mm[key] == nil {
//			mm[key] = &GroupItem{
//				Group:      item.Group,
//				Files:      make([]string, 0),
//				DriverName: item.DriverName,
//				ExtType:    item.ExtType,
//			}
//		}
//		mm[key].Files = append(mm[key].Files, item.Filename)
//	}
//
//	data := make([]GroupItem, 0)
//	for _, m := range mm {
//		data = append(data, *m)
//	}
//	return data
//}
