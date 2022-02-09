package oci_sdk

type Driver struct {
}

func (d Driver) Login() (err error) {
	return
}

func (d Driver) List(ref string) (err error) {
	return
}

func (d Driver) Upload(ref, fileName string, content []byte) (err error) {
	return
}

func (d Driver) Download(ref, fileName string) (err error) {
	//remotes.FetchHandler(store, fetcher)
	return
}
