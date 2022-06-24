package live

import (
	"reflect"
	"strings"
)

var (
	liveDataType = reflect.TypeOf(Nil)
)

type blacklist []string

func (bl blacklist) cover(pkgPath string) bool {
	for _, x := range bl {
		if strings.HasPrefix(pkgPath, x) {
			if n := len(x); n == len(pkgPath) || pkgPath[n] == '/' {
				return true
			}
		}
	}
	return false
}

type Config struct {
	Blacklist     blacklist
	SkipTypeCheck bool
}

func NewConfig(pkgBlacklist []string) Config {
	var cfg Config
	for _, x := range pkgBlacklist {
		if v := strings.TrimRight(x, `/`); v != "" {
			cfg.Blacklist = append(cfg.Blacklist, v)
		}
	}
	return cfg
}

func (cfg Config) WrapValueDirect(v any) Data {
	if v == nil {
		return Nil
	}
	if cfg.SkipTypeCheck {
		return Data{v}
	}

	cfg.checkType(reflect.TypeOf(v), 1)
	return Data{v}
}

//gocyclo:ignore
func (cfg Config) checkType(t reflect.Type, depth int) {
	if depth > 20 {
		return
	}
	if t == liveDataType {
		return
	}

	if pkgPath := t.PkgPath(); pkgPath != "" {
		if len(cfg.Blacklist) > 0 {
			if cfg.Blacklist.cover(pkgPath) {
				panic(pkgPath + " is in the blacklist of live.Config")
			}
		}
	}

	switch t.Kind() {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
		panic("live data does not support uintptr")
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Array:
		cfg.checkType(t.Elem(), depth+1)
	case reflect.Chan:
		cfg.checkType(t.Elem(), depth+1)
	case reflect.Func:
		panic("live data does not support func")
	case reflect.Interface:
		panic("live data does not support interface")
	case reflect.Map:
		keyT := t.Key()
		cfg.checkType(keyT, depth+1)
		cfg.checkType(t.Elem(), depth+1)
	case reflect.Ptr:
		cfg.checkType(t.Elem(), depth+1)
	case reflect.Slice:
		cfg.checkType(t.Elem(), depth+1)
	case reflect.String:
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if live, ok := f.Tag.Lookup("live"); ok {
				if live == "true" || live == "1" {
					continue
				}
			}
			cfg.checkType(f.Type, depth+1)
		}
	case reflect.UnsafePointer:
		panic("live data does not support unsafe pointer")
	}
}
