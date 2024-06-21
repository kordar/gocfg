package gocfg

//type ViperDriver struct {
//	ins *viper.Viper
//}
//
//func (v *ViperDriver) GetName() string {
//	return "viper"
//}
//
//func (v *ViperDriver) init(files []string, ext string) {
//	if len(files) == 0 {
//		return
//	}
//	if v.ins == nil {
//		v.ins = viper.New()
//		v.ins.SetConfigName(files[0])
//		v.ins.SetConfigType(ext)
//		v.ins.AddConfigPath(".")
//		files = files[1:]
//	}
//	err := v.ins.ReadInConfig()
//	if err != nil {
//		logger.Panic(err)
//	}
//	for _, s := range files {
//		if file, err2 := os.ReadFile(s); err2 == nil {
//			_ = v.ins.MergeConfig(bytes.NewBuffer(file))
//		}
//	}
//}
//
//func (v *ViperDriver) LoadFiles(groupItem GroupItem) {
//
//	// 剔除dev,test,pro
//	newFiles := make([]string, 0)
//	devFiles := make([]string, 0)
//	proFiles := make([]string, 0)
//	testFiles := make([]string, 0)
//	for _, file := range groupItem.Files {
//		if strings.Contains(file, "-dev.") {
//			devFiles = append(devFiles, file)
//			continue
//		}
//		if strings.Contains(file, "-pro.") {
//			proFiles = append(proFiles, file)
//			continue
//		}
//		if strings.Contains(file, "-test.") {
//			testFiles = append(testFiles, file)
//			continue
//		}
//		newFiles = append(newFiles, file)
//	}
//
//	v.init(newFiles, groupItem.ExtType)
//	// -----------
//	profile := viper.GetString("PROFILE")
//	if profile == DEV {
//		v.init(devFiles, groupItem.ExtType)
//	}
//	if profile == PRO {
//		v.init(proFiles, groupItem.ExtType)
//	}
//	if profile == TEST {
//		v.init(testFiles, groupItem.ExtType)
//	}
//
//}
//
//func (v *ViperDriver) GetInstance() interface{} {
//	return v.ins
//}
//
//func (v *ViperDriver) Get(key string) interface{} {
//	return v.ins.Get(key)
//}
//
//func (v *ViperDriver) GetString(key string) string {
//	return v.ins.GetString(key)
//}
//
//func (v *ViperDriver) GetBool(key string) bool {
//	return v.ins.GetBool(key)
//}
//
//func (v *ViperDriver) GetInt(key string) int {
//	return v.ins.GetInt(key)
//}
//
//func (v *ViperDriver) GetInt32(key string) int32 {
//	return v.ins.GetInt32(key)
//}
//
//func (v *ViperDriver) GetInt64(key string) int64 {
//	return v.GetInt64(key)
//}
//
//func (v *ViperDriver) GetUint(key string) uint {
//	return v.ins.GetUint(key)
//}
//
//func (v *ViperDriver) GetUint32(key string) uint32 {
//	return v.ins.GetUint32(key)
//}
//
//func (v *ViperDriver) GetUint64(key string) uint64 {
//	return v.ins.GetUint64(key)
//}
//
//func (v *ViperDriver) GetFloat64(key string) float64 {
//	return v.ins.GetFloat64(key)
//}
//
//func (v *ViperDriver) GetTime(key string) time.Time {
//	return v.ins.GetTime(key)
//}
//
//func (v *ViperDriver) GetDuration(key string) time.Duration {
//	return v.GetDuration(key)
//}
//
//func (v *ViperDriver) GetIntSlice(key string) []int {
//	return v.ins.GetIntSlice(key)
//}
//
//func (v *ViperDriver) GetStringSlice(key string) []string {
//	return v.ins.GetStringSlice(key)
//}
//
//func (v *ViperDriver) GetStringMap(key string) map[string]interface{} {
//	return v.ins.GetStringMap(key)
//}
//
//func (v *ViperDriver) GetStringMapString(key string) map[string]string {
//	return v.ins.GetStringMapString(key)
//}
//
//func (v *ViperDriver) GetStringMapStringSlice(key string) map[string][]string {
//	return v.ins.GetStringMapStringSlice(key)
//}
//
//func (v *ViperDriver) GetSizeInBytes(key string) uint {
//	return v.ins.GetSizeInBytes(key)
//}
//
//func (v *ViperDriver) IsSet(key string) bool {
//	return v.ins.IsSet(key)
//}
//
//func (v *ViperDriver) AllSettings() map[string]interface{} {
//	return v.ins.AllSettings()
//}
//
//func (v *ViperDriver) Unmarshal(rawVal interface{}) error {
//	return v.ins.Unmarshal(rawVal)
//}
//
//func (v *ViperDriver) UnmarshalExact(rawVal interface{}) error {
//	return v.ins.UnmarshalExact(rawVal)
//}
//
//func (v *ViperDriver) UnmarshalKey(key string, rawVal interface{}) error {
//	return v.ins.UnmarshalKey(key, rawVal)
//}
