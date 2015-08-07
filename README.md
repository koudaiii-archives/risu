# risu

## Overview

* risu is rapid image supply unit

## Description

DockerHubやQuay.ioのSaaS、CI as a Serviceでは、Queueの状態やStackの状態で待たされたり、bundle install や assets:precompileを初期状態から行うため、build時間も30m以上かかる。

risuは自前で用意したServer上でコンテナをbuildを行う。
buildをする際に、ボトルネックとなったbundle install や assets:precompileをcacheとして利用し、高速化を図る。
build後、Quay.io等のregistoryへpushをすることでbuild時間にフォーカスして高速化を図ることを目的としたツールである。

## Requirement

* Vagrant
* VirtualBox
* Go 1.4 or later

## Install

```
$ git clone https://github.com/koudaiii/risu.git
```

## Usage

### required

* set quay.io token

need set up XXX.conf? .env? .yml?

### Run

```
go command
```

* Ominiauth in GitHub

image shot 1

* set up webhook in repository

image shot 2

* git push repository

image shot 3

## FEATURE

 * GitHubにあるprivate repositoryをweb hook経由で取得する
 * cacheの更新する
  * Host上で $  bundle package で gem instlal とともに、vendor/cache/ に保存 download
  * Host上で $  bundle exec rake assets:precompile
 * 更新したcacheをADDさせて docker build する
  * bundle install --path vendor/bundle --local
 * quay.ioにimageをPush

 - [ ] JSON Scheme設計
 - [ ] GitHub連携でhookからrepositoryを取得
 - [ ] docker build
 - [ ] cache設計
 - [ ] quayへimageをPushする
 - [ ] quayからSlackなどへNotificationする

## Contribution

Please fork.

```
 $ git clone https://github.com/your_user_name/risu.git

  ```
