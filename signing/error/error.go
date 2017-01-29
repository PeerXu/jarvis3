package signing

import (
	jerrors "github.com/PeerXu/jarvis3/errors"
)

func NewSignError(typ string, dsc string, err error) jerrors.JarvisError {
	return jerrors.NewJarvisError("signing", typ, dsc, err)
}
