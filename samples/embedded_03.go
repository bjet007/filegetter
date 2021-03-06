import (
	"io"
	"log"
)

// GetFile will reach out to s3 or use the local file system to retrieve an email file
func (g *Getter) GetFile(localPath, host, bucket, key string) (io.ReadCloser, Source, error) {
	// validation against host, bucket, and key elided for brevity
	if g.useRemoteFS {
		// we have everything we need to do remote fs stuff
		fh, err := g.remoteFetcher.GetRemoteFile(g.accessKey, g.accessSecret, host, bucket, key)
		if err == nil {
			// early return
			return fh, Remote, nil
		}

		log.Println("falling back to local source - %v", err)
	}
	fh, err := g.localFetcher.Open(localPath)
	if err != nil { /* return wrapped error */
	}

	return fh, Local, nil
}