package core

const (
	LIST_COUNTRY                         = "10201170"
	LIST_NATION                          = "10201280"
	LIST_ETHNICITY                       = "10201220"
	LIST_EDUCATION                       = "10201310"
	LIST_ADDRESS                         = "10201150"
	LIST_ADDRESS_TYPE                    = "10201140"
	LIST_RESPECT                         = "10201340" // body = []
	LIST_SEGMENT_OF_CUSTOMER             = "10201180"
	LIST_RELATION_TYPE                   = "10201230"
	LIST_CUSTOMER_CATEGORY               = "10201610"
	LIST_TYPE_CUSTOMER_RELATION_CUSTOMER = "10201330"

	EDUCATION_DEGREE_LIST   = "10201200"
	EDUCATION_DEGREE_CREATE = "10201201"
	EDUCATION_DEGREE_DELETE = "10201203"
	EDUCATION_DEGREE_GET    = "10201204"

	LIST_EDUCATION_LEVEL = "10201190"

	NATION_LIST   = "10201280" // [[],0,25]
	NATION_INSERT = "10201281" // [{},{"name": "test", "name2": "test name", "orderNo": 0}]
	NATION_UPDATE = "10201282" // [{},{"nationalityId":23,"name":"test","name2":"TEST EDITED","orderNo":0}]
	NATION_DELETE = "10201283" // [{}, 22]
	NATION_SELECT = "10201284" // [22]

	ETHNICITY_LIST   = "10201220"
	ETHNICITY_INSERT = "10201221" // [{name: "Test", name2: "Test", nationalityId: "1", orderNo: 0}]
	ETHNICITY_UPDATE = "10201222" //
	ETHNICITY_DELETE = "10201223" // [31]
	ETHNICITY_SELECT = "10201224" // [31]
)
