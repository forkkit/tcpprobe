package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"reflect"
	"strings"
	"time"
)

func (c *client) printer(counter int) {
	if c.req.quiet {
		return
	}

	switch {
	case c.req.json:
		c.printJSON(counter, false)
	case c.req.jsonPretty:
		c.printJSON(counter, true)
	default:
		c.printText(counter)
	}
}

func (c *client) printText(counter int) {
	v := reflect.ValueOf(c.stats)
	filterLen := len(c.req.filter)

	ip, _, _ := net.SplitHostPort(c.addr)
	datetime := time.Unix(c.timestamp, 0).Format(time.RFC3339)
	fmt.Printf("%s target: %s (%s) seq: %d\n", datetime, c.target, ip, counter)
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if f.Tag.Get("unexported") == "true" {
			continue
		}
		if _, ok := c.req.filter[strings.ToLower(f.Name)]; ok || filterLen == 0 {
			fmt.Printf("%s:%v ", f.Name, v.Field(i).Interface())
		}
	}
	fmt.Println("")
}

func (c *client) printJSON(counter int, pretty bool) {
	var (
		b   []byte
		err error
	)

	ip, _, _ := net.SplitHostPort(c.addr)
	d := struct {
		Target    string
		IP        string
		Timestamp int64
		Seq       int
		stats
	}{
		c.target,
		ip,
		c.timestamp,
		counter,
		c.stats,
	}

	if len(c.req.filter) > 0 {
		b, err = jsonMarshalFilter(d, c.req.filter, pretty)
	} else if pretty {
		b, err = json.MarshalIndent(d, "", "  ")
	} else {
		b, err = json.Marshal(d)
	}

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}

func jsonMarshalFilter(s interface{}, filter map[string]struct{}, pretty bool) ([]byte, error) {
	var m map[string]interface{}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, &m)

	for k := range m {
		if _, ok := filter[strings.ToLower(k)]; !ok {
			delete(m, k)
		}
	}

	if pretty {
		return json.MarshalIndent(m, "", " ")
	}

	return json.Marshal(m)
}
