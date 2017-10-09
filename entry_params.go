package goWallabag

import (
	"net/url"
	"strconv"
	"strings"
)

type ParamsSetter func(*url.Values)
