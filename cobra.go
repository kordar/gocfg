package gocfg

//
///**
// * 管理多种不同驱动的配置实例
// */
//
//var cobras = make(map[string]*Cobra)
//
//func AddCobra(cobra *Cobra) {
//	cobras[cobra.Name()] = cobra
//}
//
//func GetCobra(group string, ext string, driverName string) *Cobra {
//	key := fmt.Sprintf("%s-%s-%s", group, ext, driverName)
//	return cobras[key]
//}
//
//func GetCobraWithKey(group string, key string) *Cobra {
//	for _, cobra := range cobras {
//		if cobra.GetGroupName() != group {
//			continue
//		}
//		cfg := cobra.GetCfg()
//		if cfg.IsSet(key) {
//			return cobra
//		}
//	}
//	return nil
//}
//
//func AddCobraByGroup(group GroupItem) {
//	cobra := NewCobra(group.GetKey())
//	cobra.LoadGroupItem(group)
//	AddCobra(&cobra)
//}
//
//type Cobra struct {
//	driver Driver
//	item   GroupItem
//	name   string
//}
//
//func NewCobra(name string) Cobra {
//	return Cobra{name: name}
//}
//
//func (c *Cobra) Name() string {
//	return c.name
//}
//
//func (c *Cobra) LoadGroupItem(groupItem GroupItem) {
//	c.item = groupItem
//	c.driver = CreateDriver(groupItem.DriverName)
//	c.driver.LoadFiles(groupItem)
//}
//
//func (c *Cobra) GetCfg() Driver {
//	return c.driver
//}
//
//func (c *Cobra) GetGroupName() string {
//	return c.item.Group
//}
