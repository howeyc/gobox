package main

// Applet imports
import (
	"github.com/howeyc/gobox/applets/cat"
	"github.com/howeyc/gobox/applets/checksum"
	"github.com/howeyc/gobox/applets/chroot"
	"github.com/howeyc/gobox/applets/date"
	"github.com/howeyc/gobox/applets/echo"
	"github.com/howeyc/gobox/applets/grep"
	"github.com/howeyc/gobox/applets/gzip"
	"github.com/howeyc/gobox/applets/head"
	"github.com/howeyc/gobox/applets/httpd"
	"github.com/howeyc/gobox/applets/kill"
	"github.com/howeyc/gobox/applets/ls"
	"github.com/howeyc/gobox/applets/mkdir"
	"github.com/howeyc/gobox/applets/mknod"
	"github.com/howeyc/gobox/applets/mount"
	"github.com/howeyc/gobox/applets/ps"
	"github.com/howeyc/gobox/applets/rm"
	"github.com/howeyc/gobox/applets/shell"
	"github.com/howeyc/gobox/applets/spipe"
	"github.com/howeyc/gobox/applets/strings"
	"github.com/howeyc/gobox/applets/telnetd"
	"github.com/howeyc/gobox/applets/touch"
	"github.com/howeyc/gobox/applets/umount"
	"github.com/howeyc/gobox/applets/wget"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":      echo.Echo,
	"shell":     shell.Shell,
	"spipe":     spipe.Spipe,
	"spiped":    spipe.Spiped,
	"strings":   strings.Strings,
	"telnetd":   telnetd.Telnetd,
	"touch":     touch.Touch,
	"md5sum":    checksum.Hash,
	"sha1sum":   checksum.Hash,
	"sha256sum": checksum.Hash,
	"sha512sum": checksum.Hash,
	"crc32":     checksum.Hash,
	"ls":        ls.Ls,
	"rm":        rm.Rm,
	"httpd":     httpd.Httpd,
	"wget":      wget.Wget,
	"kill":      kill.Kill,
	"cat":       cat.Cat,
	"mknod":     mknod.Mknod,
	"mount":     mount.Mount,
	"umount":    umount.Umount,
	"chroot":    chroot.Chroot,
	"ps":        ps.Ps,
	"mkdir":     mkdir.Mkdir,
	"head":      head.Head,
	"grep":      grep.Grep,
	"gzip":      gzip.Gzip,
	"gunzip":    gzip.Gunzip,
	"zcat":      gzip.Zcat,
	"date":      date.Date,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
