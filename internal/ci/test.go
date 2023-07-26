package ci

import "dagger.io/dagger"

func Test(client *dagger.Client) *dagger.Container {
	src := LoadSource(client, "./")

	container := NewBaseContainer(client.Pipeline("test"), src)

	container = container.WithExec([]string{"go", "test", "./..."})

	return container
}
