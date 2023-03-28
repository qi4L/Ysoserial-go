package ysoserial

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	TC_NULL           = 0x70
	TC_REFERENCE      = 0x71
	TC_CLASSDESC      = 0x72
	TC_OBJECT         = 0x73
	TC_STRING         = 0x74
	TC_ARRAY          = 0x75
	TC_CLASS          = 0x76
	TC_BLOCKDATA      = 0x77
	TC_ENDBLOCKDATA   = 0x78
	TC_PROXYCLASSDESC = 0x7d

	// create for gadget jdk8u20, mark NO TC_ENDBLOCKDATA write
	NO_TC_ENDBLOCKDATA = "NO_TC_ENDBLOCKDATA"

	SC_WRITE_METHOD   = 0x01
	SC_SERIALIZABLE   = 0x02
	SC_EXTERNALIZABLE = 0x04
	SC_BLOCKDATA      = 0x08

	Byte    = 'B'
	Char    = 'C'
	Double  = 'D'
	Float   = 'F'
	Integer = 'I'
	Long    = 'J'
	Short   = 'S'
	Boolean = 'Z'

	Object = 'L'
	Array  = '['
)

var (
	STREAM_MAGIC   = []byte{0xac, 0xed}
	STREAM_VERSION = []byte{0x00, 0x05}
)

type Field struct {
	kind  interface{}
	name  string
	value interface{}
}

type blockData []byte

type reference int32

type ObjectAndHandle struct {
	v      interface{}
	handle int32
}

type Handles struct {
	handle int32
	m      map[uint32]ObjectAndHandle
}

func (hand Handles) hash(kind int, v interface{}) uint32 {
	h := fnv.New32a()
	hash := uint32(0)
	switch o := v.(type) {
	case string:
		h.Write([]byte(o))
	case ClassDesc:
		hash = o.hash(kind)
	case ClassObject:
		hash = o.hash(kind)
	case *ClassObject:
		hash = o.hash(kind)
	case ProxyClassObject:
		hash = o.hash(kind)
	case *ProxyClassObject:
		hash = o.hash(kind)
	case ArrayObject:
		hash = o.hash(kind)
	case *ArrayObject:
		hash = o.hash(kind)
	default:
		log.Fatalf("unknown type: 0x%x %p", kind, v)
	}

	if hash == 0 {
		hash = h.Sum32()
	}

	return hash
}

func (hand *Handles) Get(kind int, v interface{}) (int32, bool) {
	hash := hand.hash(kind, v)
	if v, ok := hand.m[hash]; ok {
		//log.Infof("hit %d, 0x%x, %v", kind, v.handle - 0x7e0000, v.v)
		return v.handle, true
	}
	return 0, false
}

func (hand *Handles) Add(kind int, v interface{}) {
	hash := hand.hash(kind, v)
	//log.Infof("add handle, %d, 0x%x, 0x%x, %v", hash, kind, hand.handle, v)
	if _, ok := hand.m[hash]; ok {
		log.Fatalf("duplicate hash, 0x%x, %v", kind, v)
	}
	hand.m[hash] = ObjectAndHandle{
		v:      v,
		handle: hand.handle,
	}
	hand.handle++
}

func newHandles() *Handles {
	return &Handles{
		handle: 0x007e0000,
		m:      make(map[uint32]ObjectAndHandle),
	}
}

// like java class define
type ClassDesc struct {
	name             string
	serialVersionUID int64
	classDescFlags   byte
}

func (class ClassDesc) Serialize(handles *Handles) []byte {
	buff := new(bytes.Buffer)

	// refer class desc and data part
	if handle, ok := handles.Get(TC_CLASS, class); ok {
		buff.WriteByte(TC_REFERENCE)
		binary.Write(buff, binary.BigEndian, handle)
		return buff.Bytes()
	}

	buff.WriteByte(byte(TC_CLASS))

	// only refer class desc part
	if handle, ok := handles.Get(TC_CLASSDESC, class); ok {
		buff.WriteByte(TC_REFERENCE)
		binary.Write(buff, binary.BigEndian, handle)
	} else {
		buff.WriteByte(TC_CLASSDESC)
		// class name
		binary.Write(buff, binary.BigEndian, uint16(len(class.name)))
		buff.WriteString(class.name)

		// serialVersionUID
		binary.Write(buff, binary.BigEndian, class.serialVersionUID)

		// classDescFlags
		buff.WriteByte(class.classDescFlags)

		// new handle for class desc
		handles.Add(TC_CLASSDESC, class)

		// field count, always 0
		binary.Write(buff, binary.BigEndian, uint16((0)))

		// field desc, always null

		// classAnnotations, always null
		buff.WriteByte(TC_ENDBLOCKDATA)

		// super class, always null
		buff.WriteByte(TC_NULL)
	}

	// newHandle
	handles.Add(TC_CLASS, class)
	return buff.Bytes()
}

func (class ClassDesc) hash(kind int) uint32 {
	hash := fnv.New32a()
	switch kind {
	case TC_CLASSDESC: // class desc
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	case TC_OBJECT:
		hash.Write([]byte{TC_OBJECT, TC_CLASSDESC})
		hash.Write([]byte(fmt.Sprintf("%v", class)))
	case TC_CLASS:
		hash.Write([]byte{TC_CLASS, TC_CLASSDESC})
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	}
	return hash.Sum32()
}

// like java object with keyword `new` instantiation
type ClassObject struct {
	name              string
	serialVersionUID  int64
	classDescFlags    byte
	fields            []Field
	classAnnotations  []interface{}
	super             *ClassObject
	objectAnnotations []interface{}
}

func (class ClassObject) Serialize(handles *Handles) []byte {
	return class.serializeObject(handles)
}

func (class ClassObject) serializeObject(handles *Handles) []byte {
	buff := new(bytes.Buffer)

	// refer class desc and data part
	if handle, ok := handles.Get(TC_OBJECT, class); ok {
		buff.WriteByte(TC_REFERENCE)
		binary.Write(buff, binary.BigEndian, handle)
		return buff.Bytes()
	}

	buff.WriteByte(byte(TC_OBJECT))

	buff.Write(class.serializeClassDesc(handles))

	handles.Add(TC_OBJECT, class)

	buff.Write(class.serializeClassData(handles))
	return buff.Bytes()
}

func (class ClassObject) serializeClassDesc(handles *Handles) []byte {
	buff := new(bytes.Buffer)

	// only refer class desc part
	if handle, ok := handles.Get(TC_CLASSDESC, class); ok {
		buff.WriteByte(TC_REFERENCE)
		binary.Write(buff, binary.BigEndian, handle)
	} else {
		buff.WriteByte(TC_CLASSDESC)
		// class name
		binary.Write(buff, binary.BigEndian, uint16(len(class.name)))
		buff.WriteString(class.name)

		// serialVersionUID
		binary.Write(buff, binary.BigEndian, class.serialVersionUID)

		// classDescFlags
		buff.WriteByte(class.classDescFlags)

		// new handle for class desc
		handles.Add(TC_CLASSDESC, class)

		// class desc
		writeFieldsDesc(buff, handles, class.fields)

		// class annotations
		for _, v := range class.classAnnotations {
			buff.Write(serialize(handles, v))
		}
		buff.WriteByte(TC_ENDBLOCKDATA)

		// super class
		if class.super != nil {
			buff.Write(class.super.serializeClassDesc(handles))
		} else {
			buff.WriteByte(TC_NULL)
		}
	}

	return buff.Bytes()
}

func (class ClassObject) serializeClassData(handles *Handles) []byte {
	buff := new(bytes.Buffer)

	// super class data
	if class.super != nil {
		buff.Write(class.super.serializeClassData(handles))
	}

	// class data
	for _, f := range class.fields {
		buff.Write(serialize(handles, f.value))
	}

	// object annotations
	if class.classDescFlags&(SC_WRITE_METHOD|SC_BLOCKDATA) > 0 {
		endFlag := true
		for _, v := range class.objectAnnotations {
			if v == NO_TC_ENDBLOCKDATA {
				endFlag = false
			} else if endFlag {
				buff.Write(serialize(handles, v))
			} else {
				log.Fatal("NO_TC_ENDBLOCKDATA should be the last one")
			}
		}

		if endFlag {
			buff.WriteByte(TC_ENDBLOCKDATA)
		}
	}

	return buff.Bytes()
}

func (class ClassObject) SerializeWithMagicHeader() []byte {
	buff := new(bytes.Buffer)
	buff.Write(STREAM_MAGIC)
	buff.Write(STREAM_VERSION)
	buff.Write(class.Serialize(newHandles()))
	return buff.Bytes()
}

func (class ClassObject) hash(kind int) uint32 {
	hash := fnv.New32a()
	switch kind {
	case TC_CLASSDESC: // class desc
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	case TC_OBJECT: // class desc include super class desc
		hash.Write([]byte{TC_OBJECT, TC_CLASSDESC})
		hash.Write(class.serializeClassDesc(newHandles()))
		hash.Write(class.serializeClassData(newHandles()))
	case TC_CLASS:
		hash.Write([]byte{TC_CLASS, TC_CLASSDESC})
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	}
	return hash.Sum32()
}

type ProxyClassObject struct {
	interfaces       []string
	name             string
	serialVersionUID int64
	classDescFlags   byte
	fields           []Field
}

func (class ProxyClassObject) Serialize(handles *Handles) []byte {
	buff := new(bytes.Buffer)
	buff.WriteByte(TC_OBJECT)
	buff.WriteByte(TC_PROXYCLASSDESC)
	handles.Add(TC_PROXYCLASSDESC, class)

	binary.Write(buff, binary.BigEndian, int32(len(class.interfaces)))

	for _, name := range class.interfaces {
		binary.Write(buff, binary.BigEndian, int16(len(name))) // string length
		buff.WriteString(name)
	}

	// classAnnotations, classAnnotations always is null
	buff.WriteByte(TC_ENDBLOCKDATA)

	// superClassDesc
	{
		buff.WriteByte(TC_CLASSDESC)

		// class name
		binary.Write(buff, binary.BigEndian, int16(len(class.name)))
		buff.WriteString(class.name)

		// serialVersionUID
		binary.Write(buff, binary.BigEndian, class.serialVersionUID)

		handles.Add(TC_CLASSDESC, class)

		// classDescFlags
		buff.WriteByte(class.classDescFlags)

		// field desc
		writeFieldsDesc(buff, handles, class.fields)

		// classAnnotations, always null
		buff.WriteByte(TC_ENDBLOCKDATA)

		// superClassDesc, always null
		buff.WriteByte(TC_NULL)
	}

	handles.Add(TC_OBJECT, class)
	// class data
	// super fields data
	for _, f := range class.fields {
		buff.Write(serialize(handles, f.value))
	}

	return buff.Bytes()
}

func (class ProxyClassObject) hash(kind int) uint32 {
	hash := fnv.New32a()
	switch kind {
	case TC_PROXYCLASSDESC:
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	case TC_OBJECT:
		hash.Write([]byte{TC_PROXYCLASSDESC})
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	}
	return hash.Sum32()
}

type ArrayObject struct {
	name             string
	serialVersionUID int64
	classDescFlags   byte
	values           []interface{}
	annotations      []interface{}
}

func (class ArrayObject) Serialize(handles *Handles) []byte {
	buff := new(bytes.Buffer)
	if handle, ok := handles.Get(TC_ARRAY, class); ok {
		buff.WriteByte(TC_REFERENCE)
		binary.Write(buff, binary.BigEndian, handle)
		return buff.Bytes()
	}
	buff.WriteByte(TC_ARRAY)

	// only class desc part
	if handle, ok := handles.Get(TC_CLASSDESC, class); ok {
		buff.WriteByte(TC_REFERENCE)
		binary.Write(buff, binary.BigEndian, handle)
	} else {
		buff.WriteByte(TC_CLASSDESC)
		// class name
		binary.Write(buff, binary.BigEndian, uint16(len(class.name)))
		buff.WriteString(class.name)

		// serialVersionUID
		binary.Write(buff, binary.BigEndian, class.serialVersionUID)

		// classDescFlags
		buff.WriteByte(class.classDescFlags)

		// newHandle
		handles.Add(TC_CLASSDESC, class)

		// fields count, always 0
		binary.Write(buff, binary.BigEndian, int16(0))

		// classAnnotations, always null
		buff.WriteByte(TC_ENDBLOCKDATA)

		// superClassDes, always null
		buff.WriteByte(TC_NULL)
	}

	// newHandle
	handles.Add(TC_ARRAY, class)

	binary.Write(buff, binary.BigEndian, int32(len(class.values)))

	// array values
	for _, f := range class.values {
		buff.Write(serialize(handles, f))
	}

	return buff.Bytes()
}

func (class ArrayObject) hash(kind int) uint32 {
	hash := fnv.New32a()
	switch kind {
	case TC_CLASSDESC:
		hash.Write([]byte(class.name))
		binary.Write(hash, binary.BigEndian, class.serialVersionUID)
	case TC_ARRAY:
		hash.Write([]byte{TC_ARRAY})
		hash.Write([]byte(fmt.Sprintf("%v", class)))
	}
	return hash.Sum32()
}

func serialize(handles *Handles, v interface{}) []byte {
	buff := new(bytes.Buffer)
	switch o := v.(type) {
	case nil:
		buff.WriteByte(TC_NULL)
	case blockData:
		// blockdata:
		//  blockdatashort
		//  blockdatalong
		//
		// blockdatashort:
		//   TC_BLOCKDATA (unsigned byte)<size> (byte)[size]
		//
		// blockdatalong:
		//   TC_BLOCKDATALONG (int)<size> (byte)[size]
		buff.WriteByte(TC_BLOCKDATA)
		buff.WriteByte(byte(len(o)))
		buff.Write(o)
	case reference:
		buff.WriteByte(TC_REFERENCE)
		if o >= 0x007e0000 {
			binary.Write(buff, binary.BigEndian, o)
		} else {
			binary.Write(buff, binary.BigEndian, handles.handle+int32(o))
		}
	case ClassDesc:
		buff.Write(o.Serialize(handles))
	case ClassObject:
		buff.Write(o.Serialize(handles))
	case *ClassObject:
		buff.Write(o.Serialize(handles))
	case ProxyClassObject:
		buff.Write(o.Serialize(handles))
	case *ProxyClassObject:
		buff.Write(o.Serialize(handles))
	case ArrayObject:
		buff.Write(o.Serialize(handles))
	case string:
		// TC_STRING newHandle (utf)
		if handle, ok := handles.Get(TC_STRING, o); ok {
			buff.WriteByte(TC_REFERENCE)
			binary.Write(buff, binary.BigEndian, handle)
		} else {
			buff.WriteByte(TC_STRING)
			handles.Add(TC_STRING, o)
			binary.Write(buff, binary.BigEndian, uint16(len(o)))
			buff.WriteString(o)
		}
	case int:
		binary.Write(buff, binary.BigEndian, int32(o))
	case float64:
		binary.Write(buff, binary.BigEndian, float32(o))
	case bool:
		if o {
			buff.WriteByte(0x01)
		} else {
			buff.WriteByte(0x00)
		}
	case byte:
		buff.WriteByte(o)
	default:
		log.Fatal("serialize", v)
	}

	return buff.Bytes()
}

func interfaces(vs ...interface{}) []interface{} {
	arr := make([]interface{}, len(vs))
	for i, v := range vs {
		arr[i] = v
	}
	return arr
}

func bytesToInterfaces(bytes []byte) []interface{} {
	ifs := make([]interface{}, len(bytes))
	for i := range bytes {
		ifs[i] = bytes[i]
	}
	return ifs
}

func writeFieldsDesc(buff *bytes.Buffer, handles *Handles, fields []Field) {
	binary.Write(buff, binary.BigEndian, uint16(len(fields)))

	// field desc
	for _, f := range fields {
		switch v := f.kind.(type) {
		case int32:
			buff.WriteByte(byte(v))
		case string:
			if strings.HasPrefix(v, string(Object)) {
				buff.WriteByte(Object)
			} else if strings.HasPrefix(v, string(Array)) {
				buff.WriteByte(Array)
			} else {
				log.Fatal("unknown type:", f)
			}
		default:
			log.Fatal("unknown type:", f)
		}

		// field name
		binary.Write(buff, binary.BigEndian, int16(len(f.name)))
		buff.WriteString(f.name)

		// class name
		if v, ok := f.kind.(string); ok {
			buff.Write(serialize(handles, v))
		}
	}
}
