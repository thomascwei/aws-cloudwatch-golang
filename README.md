# 利用aws sdk讀cloud watch data

## 配置
在`~/.aws`目錄下配置兩個檔案
* config
```yaml
[default]
region = <your-region-name>
```
* credentials
```yaml
[default]
aws_access_key_id = <your-aws-access-key-id>
aws_secret_access_key = <your-aws-secret-key>
```
在docker裏要放在`/root/.aws`
