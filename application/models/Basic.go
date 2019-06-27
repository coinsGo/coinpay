package models

var Db *DB

var CfgHeights map[uint]uint64

var CfgConfirms map[uint]uint64

func Setup(env string) {

	CfgHeights = GetHeights()
	CfgConfirms = GetConfirms()

}
