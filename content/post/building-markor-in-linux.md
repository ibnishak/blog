---
title: "Building markor from source code in linux"
date: 2019-12-12
categories: ["digital"]
tags: ["android"]
author: Riz
---


Markor can be installed from Google playstore, FDroid as well as releases page of this repo. Nevertheless, if you want the cutting edge version, here are the general instructions for Linux platform. Open terminal run the following set of commands one by one.


- If you do not have android-sdk installed, you need to install it first. Most package managers has it. Given below is a portable method.

```bash
export ANDROID_SDK_ROOT="$HOME/android-sdk"

if [[ ! -d ${ANDROID_SDK_ROOT}]]; then
    curl https://dl.google.com/android/repository/sdk-tools-linux-3859397.zip > sdk-tools.zip 
    unzip -qq -n sdk-tools.zip -d $HOME/android-sdk
fi
echo y | ${ANDROID_SDK_ROOT}/tools/bin/sdkmanager 'tools'
echo y | ${ANDROID_SDK_ROOT}/tools/bin/sdkmanager 'platform-tools' 
echo y | ${ANDROID_SDK_ROOT}/tools/bin/sdkmanager 'build-tools;29.0.3' 
echo y | ${ANDROID_SDK_ROOT}/tools/bin/sdkmanager 'platforms;android-29' 
echo y | ${ANDROID_SDK_ROOT}/tools/bin/sdkmanager 'extras;google;m2repository' 
${ANDROID_SDK_ROOT}/tools/bin/sdkmanager --licenses
```

- Clone markor to $HOME/markor

```git
git clone --depth=1 https://github.com/gsantner/markor.git && cd markor
```


* Then run the gradlew scripts
```
./gradlew --no-daemon clean
./gradlew --no-daemon build

```
* Change directory to build folder. Depending on the state of upstream repository or your git clone, this could be `flavorDefault` or `flavorAtest`

```bash
cd app/build/outputs/apk/flavorDefault/release 
ls
```
- You will see an unsigned apk created with name the reads like `net.gsantner.markor-v***********unsigned.apk`. We are going to align the apk first before signing it.
```bash
${ANDROID_SDK_ROOT}/build-tools/28.0.3/zipalign -v -p 4 net.gsantner.markor-v99-2.1.6-flavorDefault-release-unsigned.apk markor-unsign-aligned-apk
```
- At this point you need to create a signing key witha a unique alias. In the following command, change the `my-unique-name` with some unique combination of words
```bash
keytool -genkey -v -keystore markor-release-key.jks -keyalg RSA -keysize 2048 -validity 10000 -alias my-unique-name 
```
- Now we use the key just created to sign it.

```bash
${ANDROID_SDK_ROOT}/build-tools/28.0.3/apksigner sign --ks markor-release-key.jks --out markor-signed-aligned.apk markor-unsign-aligned.apk
${ANDROID_SDK_ROOT}/build-tools/28.0.3/apksigner verify markor-signed-aligned.apk
```

* Copy the apk to phone and install it.