package environment

import (
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	// ไฟล์ที่จะจัดเก็บตัว Connection string Database
	viper.SetConfigName("environment")
	// ภาษาที่จะใช้ในการ Config
	viper.SetConfigType("yaml")
	// ที่อยู่ของ file config เริ่มค้นหาจาก root ด้านนอกสุด
	viper.AddConfigPath("./environment")
	// แล้วเข้ามาที่ environment folder 
	viper.AddConfigPath("environment")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// เรียก file config.yaml มาใช้
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}