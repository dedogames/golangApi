package configuration

type ConfigStruct struct {
	DynamoDBConfig struct {
		Region            string `yaml:"region_name"`
		ProductTable      string `yaml:"product_table"`
		ConnectionTimeout int    `yaml:"connection_timeout"`
		ReadTimeout       int    `yaml:"read_timeout"`
		MaxAttempts       int    `yaml:"max_attempts"`
	} `yaml:"dynamoDb"`
}
