package config

const (
	DBReadDSN                      = "postgresql://postgres:password@postgres-db-read:5432/praisindo"
	DBWriteDSN                     = "postgresql://postgres:password@postgres-db-write:5433/praisindo"
	TopicInsertUser                = "insert-user"
	KafkaBrokerAddress             = "broker:19092"
	ReadDBInserterConsumerGroupID  = "read-db-cg-1"
	WriteDBInserterConsumerGroupID = "write-db-cg-1"
)
