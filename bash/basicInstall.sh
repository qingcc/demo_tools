#bin/bash

#是否按照yum
hasYum=`yum --version | grep "install yum"`
if [ "${hasYum}" != "" ]; then
  echo "install yum"
  apt install yum
fi

#是否按照git
hasGit=`git --version | grep "command not found"`
if [ "${hasGit}" != "" ]; then
  echo "install git"
  yum install git
  name="qing"
  email="qingcc0503@163.com"
  read -p "please enter git name:" inputName
  if [ "${inputName}" != "" ];then
    name=$inputName
  fi
  read -p "please enter git email:" inputEmail
  if [ "${inputEmail}" != "" ];then
    email=$inputEmail
  fi
  git config --global user.name name
  git config --global user.email email
  ssh-keygen -o
fi

