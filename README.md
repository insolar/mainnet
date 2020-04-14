<h1 align="center"> Insolar MainNet </h1> <br>

[<img src="https://github.com/insolar/doc-pics/raw/master/st/github-readme-banner.png">](http://insolar.io/?utm_source=Github)

# Introduction 
Insolar MainNet is the application that implements smart contracts logic for Insolar MainNet. 

This application works ontop of [Insolar platform](https://github.com/insolar/insolar) and allows you to:

* Create Insolar wallet

* Migrate INS onto Insolar exchanging INS to XNS 1:1 where INS is Ethereum ERC-20 token and XNS is Insolar native token.

* Deposit migrated tokens into the wallet

* Transfer XNS to Insolar MainNet users

* Receive XNS from Insolar MainNet users


# Quick start

You can test Insolar Mainnet locally. 
To do that, you need to install and deploy it as described below.

## Install

1. Install the latest 1.12 version of the [Golang programming tools](https://golang.org/doc/install#install). 
   Make sure the `$GOPATH` environment variable is set.

2. Download the Mainnet package:

   ```
   go get github.com/insolar/mainnet
   ```

3. Go to the package directory:

   ```
   cd $GOPATH/src/github.com/insolar/mainnet
   ```

4. Install dependencies and build binaries using [the makefile](https://github.com/insolar/mainnet/blob/master/Makefile) that automates this process:

   ```
   make
   ```

## Deploy locally

1. Run the launcher:

   ```
   insolar-scripts/insolard/launchnet.sh -g
   ```

   The launcher generates te necessary bootstrap data, starts a pulse watcher, and launches a number of nodes. 
   In a local setup, the "nodes" are simply services listening on different ports.
   The default number of nodes is 5. You can vary this number by commenting/uncommenting nodes in the `discovery_nodes` section in `scripts/insolard/bootstrap_template.yaml`.

2. When the pulse watcher says `INSOLAR STATE: READY`, you can run a benchmark:
     ```
     bin/benchmark -c=4 -r=25 -k=.artifacts/launchnet/configs/
     ```

     Options:
     * `-k`: Path to the root user's key pair.
     * `-c`: Number of concurrent threads in which requests are sent.
     * `-r`: Number of transfer requests to be sent in each thread.

# Contribute!

Feel free to submit issues, fork the repository and send pull requests! 

To make the process smooth for both reviewers and contributors, familiarize yourself with the following guidelines:

1. [Open source contributor guide](https://github.com/freeCodeCamp/how-to-contribute-to-open-source).
2. [Style guide: Effective Go](https://golang.org/doc/effective_go.html).
3. [List of shorthands for Go code review comments](https://github.com/golang/go/wiki/CodeReviewComments).

When submitting an issue, **include a complete test function** that reproduces it.

Thank you for your intention to contribute to the Insolar Mainnet project. As a company developing open-source code, we highly appreciate external contributions to our project.

# Contacts

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

# License

This project is licensed under the terms of the [Insolar License 1.0](LICENSE.md).
