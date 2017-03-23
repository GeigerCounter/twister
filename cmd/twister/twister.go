package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/go-errors/errors"
	"github.com/mattn/go-isatty"
	"github.com/mitchellh/go-homedir"
	"github.com/zyedidia/tcell"
	"github.com/zyedidia/tcell/encoding"
)

const (
	
)

var (
	screen tcell.Screen
	configDir string
)

