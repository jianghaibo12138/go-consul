package consul_client

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func packageQueryStrParam(dc, tag, near, nodeMeta, filter, ns, acquire, release, format string) string {
	var paraStrArr []string
	if dc != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("dc=%s", dc))
	}
	if tag != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("tag=%s", tag))
	}
	if near != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("near=%s", near))
	}
	if nodeMeta != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("node-meta=%s", nodeMeta))
	}
	if filter != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("filter=%s", filter))
	}
	if ns != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("ns=%s", ns))
	}
	if acquire != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("acquire=%s", acquire))
	}
	if release != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("release=%s", release))
	}
	if format != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("format=%s", format))
	}
	var paraStr string
	if len(paraStrArr) != 0 {
		paraStr = strings.Join(paraStrArr, "&")
	}
	return paraStr
}

func packageQueryBoolParam(recurse, raw, keys, separator, passing, replaceExistingChecks bool) string {
	var paraStrArr []string
	if recurse {
		paraStrArr = append(paraStrArr, fmt.Sprintf("recurse=%v", recurse))
	}
	if raw {
		paraStrArr = append(paraStrArr, fmt.Sprintf("raw=%v", raw))
	}
	if keys {
		paraStrArr = append(paraStrArr, fmt.Sprintf("keys=%v", keys))
	}
	if separator {
		paraStrArr = append(paraStrArr, fmt.Sprintf("separator=%v", separator))
	}
	if passing {
		paraStrArr = append(paraStrArr, fmt.Sprintf("passing=%v", passing))
	}
	if replaceExistingChecks {
		paraStrArr = append(paraStrArr, fmt.Sprintf("replace-existing-checks=%v", passing))
	}
	var paraStr string
	if len(paraStrArr) != 0 {
		paraStr = strings.Join(paraStrArr, "&")
	}
	return paraStr
}

func packageQueryIntParam(flags, cas int) string {
	var paraStrArr []string
	if flags != 0 {
		paraStrArr = append(paraStrArr, fmt.Sprintf("flags=%d", flags))
	}
	if cas != 0 {
		paraStrArr = append(paraStrArr, fmt.Sprintf("cas=%d", cas))
	}
	var paraStr string
	if len(paraStrArr) != 0 {
		paraStr = strings.Join(paraStrArr, "&")
	}
	return paraStr
}

func ReadJsonConf(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bytes
}
