package gocfg

//
//// GetAllFile 递归获取指定目录下的所有文件名
//func GetAllFile(pathname string, ext ...string) ([]string, error) {
//	result := make([]string, 0)
//
//	fis, err := ioutil.ReadDir(pathname)
//	if err != nil {
//		logger.Errorf("[gocfg] 读取文件目录失败，pathname=%v, err=%v", pathname, err)
//		return result, err
//	}
//
//	// 所有文件/文件夹
//	for _, fi := range fis {
//		fullname := path.Join(pathname, fi.Name())
//		// 是文件夹则递归进入获取;是文件，则压入数组
//		if fi.IsDir() {
//			temp, err2 := GetAllFile(fullname, ext...)
//			if err2 != nil {
//				logger.Errorf("[gocfg] 读取文件目录失败,fullname=%v, err=%v", fullname, err)
//				return result, err2
//			}
//			result = append(result, temp...)
//		} else {
//			suffix := path.Ext(fullname)[1:]
//			flag := false
//			for _, s := range ext {
//				if suffix == s {
//					flag = true
//					break
//				}
//			}
//			if flag {
//				result = append(result, fullname)
//			}
//		}
//	}
//
//	return result, nil
//}
