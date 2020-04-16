[<img src="https://github.com/insolar/doc-pics/raw/master/st/github-readme-banner.png">](http://insolar.io/?utm_source=Github)

# Insolar MainNet

## Introduction 
Insolar MainNet is the application that implements smart contracts logic for Insolar MainNet. 

This application works ontop of [Insolar platform](https://github.com/insolar/insolar) and allows you to:

* Create Insolar wallet

* Migrate INS onto Insolar exchanging INS to XNS 1:10 where INS is Ethereum ERC-20 token and XNS is Insolar native token.

* Deposit migrated tokens into the wallet

* Transfer XNS to Insolar MainNet users

* Receive XNS from Insolar MainNet users


## Quick start

You can test Insolar Mainnet locally:

1. Install everything from the **Prerequisites** section.
2. Install this application.
3. Deploy this application locally.
4. Test this application locally.

### Prerequisites

Install Golang programming tools **v1.12**:

#### Fresh Go installation

* Linux: install the tools v1.12 from [golang.org](https://golang.org/doc/install#install). It is recommended to use the default settings. Set the [`$GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH).

* macOS: use Homebrew to install the tools: `brew install go@1.12`. Set the [`$GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH).

#### Multiple versions of Go on Linux

If you already have another version installed and want to keep it, you can [install Go v1.12 via go get](https://golang.org/doc/install#extra_versions) or [use GVM](https://github.com/moovweb/gvm).

#### Multiple versions of Go on macOS

If you're a macOS user and installed a Go version from golang.org, then address **Multiple versions of Go on Linux**.

If you used `brew` to install another version of the go package, then install go@1.12 and switch to it:

```
brew install go@1.12
brew unlink go
brew link go@1.12 --force
```

This is if you have a generic go package installed. You can unlink go@1.12 later and link back your go package.

```
brew install go@1.12
brew switch go@1.12
```

This is if you're already using a specific version of the tools via a go@ package.

### Install 

1. Download the Mainnet package:

   ```
   go get github.com/insolar/mainnet
   ```

2. Go to the package directory:

   ```
   cd $GOPATH/src/github.com/insolar/mainnet
   ```

3. Install dependencies and build binaries using [the makefile](https://github.com/insolar/mainnet/blob/master/Makefile) that automates this process:

   ```
   make
   ```

### Deploy locally
 
1. In the directory where you downloaded the Mainnet package to, run the launcher:

   ```
   insolar-scripts/insolard/launchnet.sh -g
   ```

   The launcher generates the necessary bootstrap data, starts a pulse watcher, and launches a number of nodes. <br>
   In a local setup, the "nodes" are simply services listening on different ports.<br>
   
   The default number of nodes is 5. You can vary this number by commenting/uncommenting nodes in the `discovery_nodes` section in `scripts/insolard/bootstrap_template.yaml`.
   
### Test locally

#### Benchmark test
When the pulse watcher says `INSOLAR STATE: READY`, you can run a benchmark in another terminal tab/window:
     ```
     bin/benchmark -c=4 -r=25 -k=.artifacts/launchnet/configs/
     ```

     Options:
     * `-k`: Path to the root user's key pair.
     * `-c`: Number of concurrent threads in which requests are sent.
     * `-r`: Number of transfer requests to be sent in each thread.
     
#### Functional tests

The tests include creating user accounts, migrating INS to Mainnet using mockups, exchanging INS for XNS using mockups, transferring XNS mockups between user accounts.

Lifehack: you can also automatically install, deploy and run functional tests via `make` with the key `functest`:

   ```
   make functest
   ```

## Contribute!

Feel free to submit issues, fork the repository and send pull requests! 

To make the process smooth for both reviewers and contributors, familiarize yourself with the following guidelines:

1. [Open source contributor guide](https://github.com/freeCodeCamp/how-to-contribute-to-open-source).
2. [Style guide: Effective Go](https://golang.org/doc/effective_go.html).
3. [List of shorthands for Go code review comments](https://github.com/golang/go/wiki/CodeReviewComments).

When submitting an issue, **include a complete test function** that reproduces it.

Thank you for your intention to contribute to the Insolar Mainnet project. As a company developing open-source code, we highly appreciate external contributions to our project.

## Contacts

If you have any additional questions, join our [developers chat on Telegram](https://t.me/InsolarTech).

Our social media:

[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-facebook.png" width="36" height="36">](https://facebook.com/insolario)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-twitter.png" width="36" height="36">](https://twitter.com/insolario)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-medium.png" width="36" height="36">](https://medium.com/insolar)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-youtube.png" width="36" height="36">](https://youtube.com/insolar)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-reddit.png" width="36" height="36">](https://www.reddit.com/r/insolar/)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-linkedin.png" width="36" height="36">](https://www.linkedin.com/company/insolario/)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-instagram.png" width="36" height="36">](https://instagram.com/insolario)
[<img src="https://github.com/insolar/doc-pics/raw/master/st/ico-social-telegram.png" width="36" height="36">](https://t.me/InsolarAnnouncements) 

## License

This project is licensed under the terms of the [Insolar License 1.0](LICENSE.md).
