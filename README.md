# Netflix File Chunker Service

## Preparation

Before using this service you need to know what ffmpeg_go is wrapper on linux command ffmpeg.  
So you need to install on your machine ffmpeg.

*Ubuntu 22.04*
```
sudo apt intall ffmpeg
```

## Configure the access for ObjectStorage

1. Go to ~/.aws/ (macOS and Linux) or C:\Users\<username>\.aws\ (Windows).

2. Create a **credentials** file with authentication data for Object Storage and copy the following information into it:
```
[default]
aws_access_key_id = <идентификатор_статического_ключа>
aws_secret_access_key = <секретный_ключ>
```

3. Create a **config** file with the default region settings and copy the following information into it:
```

[default]
region=ru-central1
```