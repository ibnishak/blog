---
title: "Building markor from source code in linux"
date: 2019-12-12
tags: ["android"]
author: Riz
---


Markor can be installed from Google playstore, FDroid as well as releases page of this repo. Nevertheless, if you want the cutting edge version, here are the general instructions for Linux platform. Open terminal run the following set of commands one by one.


- If you do not have android-sdk installed, you need to install it first. 

```bash
mkdir $HOME/android-sdk-dl
if test ! -e $HOME/android-sdk-dl/sdk-tools.zip ; then curl https://dl.google.com/android/repository/sdk-tools-linux-3859397.zip > $HOME/android-sdk-dl/sdk-tools.zip ; fi
unzip -qq -n $HOME/android-sdk-dl/sdk-tools.zip -d $HOME/android-sdk
echo y | $HOME/android-sdk/tools/bin/sdkmanager 'tools'
echo y | $HOME/android-sdk/tools/bin/sdkmanager 'platform-tools' 
echo y | $HOME/android-sdk/tools/bin/sdkmanager 'build-tools;26.0.2' 
echo y | $HOME/android-sdk/tools/bin/sdkmanager 'platforms;android-27' 
echo y | $HOME/android-sdk/tools/bin/sdkmanager 'extras;google;m2repository' 
$HOME/android-sdk/tools/bin/sdkmanager --licenses
export ANDROID_HOME=$HOME/android-sdk

* Clone markor to $HOME/markor
```git
git clone --depth=1 https://github.com/gsantner/markor.git && cd markor
```


* Then run the gradlew scripts
```
./gradlew --no-daemon clean
./gradlew --no-daemon build

```
* Change directory to releases folder
```
cd app/build/outputs/apk/flavorDefault/release 
ls
```
* You will see an unsigned apk created with name the reads like `net.gsantner.markor-v***********unsigned.apk`. We are going to align the apk first before signing it.
```
$HOME/android-sdk/build-tools/28.0.3/zipalign -v -p 4 net.gsantner.markor-v99-2.1.6-flavorDefault-release-unsigned.apk markor-unsign-aligned-apk
```
* At this point you need to create a signing key witha a unique alias. In the following command, change the `my-unique-name` with some unique combination of words
```
keytool -genkey -v -keystore markor-release-key.jks -keyalg RSA -keysize 2048 -validity 10000 -alias my-unique-name 
```
* Now we use the key just created to sign it.
```
$HOME/android-sdk/build-tools/28.0.3/apksigner sign --ks markor-release-key.jks --out markor-signed-aligned.apk markor-unsign-aligned.apk
$HOME/android-sdk/build-tools/28.0.3/apksigner verify markor-signed-aligned.apk
```

* Copy the apk to phone and install it.