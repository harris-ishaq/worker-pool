package config

import "os"

var (
	HOSTNAME, _ = os.Hostname()
	// DBURL database url
	DBURL string = os.Getenv("DBLNK")
	// DBLogURL database log url
	DBLogURL string = os.Getenv("DBLNKLOG")
	// DBUSER database user
	DBUSER string = os.Getenv("DBUSER")
	// DBPASS database password
	DBPASS string = os.Getenv("DBPASS")
	// DBNAME database name
	DBNAME string = os.Getenv("DBUSE")
	// DBLogName database log name
	DBLogNAME string = os.Getenv("DBLOG")
)

var (
	// NATSURL nats streaming server url
	NATSURL string = os.Getenv("NATSURL")
	// CLUSTERID cluster id for nats server
	CLUSTERID string = os.Getenv("NATSCLUSTER")
	// Channel Name
	CH_POSTINGDATE  string = os.Getenv("POSTING_DATE")
	CH_REPORTUPDATE string = os.Getenv("CHANNEL_REPORT_UPDATE")
)

var (
	TIMEZONE   string = os.Getenv("TZ")
	ACCT_NO    string = os.Getenv("SOURCE_ACCT_NO")
	GRANT_TYPE string = os.Getenv("GRANT_TYPE")
	STARTDATE  string = os.Getenv("START_DATE")
	ENDDATE    string = os.Getenv("END_DATE")
)
