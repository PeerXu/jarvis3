package computing

import (
	jerrors "github.com/PeerXu/jarvis3/errors"
)

func NewComputeError(typ string, dsc string, err error) jerrors.JarvisError {
	return jerrors.NewJarvisError("computing", typ, dsc, err)
}
