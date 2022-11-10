package psql

type Config struct {
	DriverName     string `mapstructure:"driverName"`
	DataSourceName string `mapstructure:"dataSourceName"`
}
