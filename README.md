# ImgOps

This is a CLI tool for reverse searching images through ImgOps website.
It supports files and URLs.

## How It Works

1. You upload a file or pass a URL to the tool
2. Depending on your choice, it will open search results from sites you want.  
  These are Google, TinEye, Yandex, Bing, Reddit and few others  
  If you don't provide a “target”, it will open ImgOps page with the provided image.

## Download

[Click here](https://github.com/dogancelik/imgops/releases/latest) to download latest version.

## How To Use

Open up your command prompt:

### Upload a file

Reverse search a local file and open Google & TinEye results:

```batch
imgops "C:\my_file.jpg" google,tineye
```

### Pass a URL

Reverse search a URL and open Google results:

```batch
imgops "http://example.com/foobar.jpg" google
```
