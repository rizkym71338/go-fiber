package libs

import "gorm.io/gorm"

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
