package vfs

import (
        "embed"
)

//go:embed vfs/bnetserver.*
//go:embed vfs/Aurora*
//go:embed vfs/*_bundle.binpb
var VFS embed.FS
