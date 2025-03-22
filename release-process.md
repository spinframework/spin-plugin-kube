# Cutting a new release of the Spin Kube plugin

To cut a new release of the Spin Kube plugin, you will need to do the following:

1. Confirm that [CI is green](https://github.com/spinframework/spin-trigger-sqs/actions) for the commit selected to be tagged and released.

2. Update [`go.mod`](https://github.com/spinframework/spin-plugin-kube/blob/main/go.mod#L13) to ensure that the Spin Operator module is at the desired/latest version. Create a pull request with these changes and merge once approved.

3. Change the version number in [spin-pluginify.toml](./spin-pluginify.toml). Create a pull request with these changes and merge once approved.

4. Checkout the commit with the version bump from above.

5. Create and push a new tag with a `v` and then the version number.

    As an example, via the `git` CLI:

    ```console
    # Create a GPG-signed and annotated tag
    git tag -s -m "Spin Kube Plugin v0.4.0" v0.4.0

    # Push the tag to the remote corresponding to spinframework/spin-kube-plugin (here 'origin')
    git push origin v0.4.0
    ```

6. Pushing the tag upstream will trigger the [release action](https://github.com/spinframework/spin-plugin-kube/blob/main/.github/workflows/release.yml).
    - The release build will create the packaged versions of the plugin, the updated plugin manifest and a checksums file
    - These assets are uploaded to a new GitHub release for the pushed tag
    - Release notes are auto-generated but edit as needed especially around breaking changes or other notable items
  
7. Validate that CI created a PR in the [fermyon/spin-plugins](https://github.com/fermyon/spin-plugins) repository with the [updated manifest](https://github.com/fermyon/spin-plugins/tree/main/manifests/kube).

8. If applicable, create PR(s) or coordinate [documentation](https://github.com/spinframework/spinkube-docs) needs, e.g. for new features or updated functionality.
