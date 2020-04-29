#bin/bash

read -p "请输入存放证书的目录:" dir
if [ ! -d $dir ];then
  echo "该目录不存在"
  read -p "请输入徐需要创建的目录" dir
  mkdir $dir
fi
read -p "请输入密钥名称：" name
openssl genrsa -out ${dir}/${name}.key &>/dev/null    #-----使用openssl生成私匙证书
if [ $? -eq 0 ];then
  echo "sucessful"
else
  echo "fail"
fi
openssl req -new -x509 -key ${dir}/${name}.key -subj "/CN=commmon" -out ${dir}/${name}.crt &>/dev/null #------subj选项可以生成证书时，非交互自动填写Common Name信息
if [ $? -eq 0 ];then
  echo "sucessful"
else
  echo "fail"
fi